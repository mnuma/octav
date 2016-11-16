package service

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
	"time"

	"context"

	"cloud.google.com/go/storage"

	"github.com/builderscon/octav/octav/db"
	"github.com/builderscon/octav/octav/internal/errors"
	"github.com/builderscon/octav/octav/model"
	"github.com/builderscon/octav/octav/tools"
	"github.com/lestrrat/go-pdebug"
)

func (v *ConferenceSvc) Init() {
	v.mediaStorage = MediaStorage
	v.credentialStorage = CredentialStorage
}

func (v *ConferenceSvc) populateRowForCreate(vdb *db.Conference, payload *model.CreateConferenceRequest) error {
	vdb.EID = tools.UUID()
	vdb.Slug = payload.Slug
	vdb.Title = payload.Title
	vdb.SeriesID = payload.SeriesID
	vdb.Status = "private"

	if payload.SubTitle.Valid() {
		vdb.SubTitle.Valid = true
		vdb.SubTitle.String = payload.SubTitle.String
	}

	if payload.Timezone.Valid() {
		if _, err := time.LoadLocation(payload.Timezone.String); err == nil {
			vdb.Timezone = payload.Timezone.String
		}
	} else {
		vdb.Timezone = "UTC"
	}
	return nil
}

func (v *ConferenceSvc) populateRowForUpdate(vdb *db.Conference, payload *model.UpdateConferenceRequest) error {
	if payload.SeriesID.Valid() {
		vdb.SeriesID = payload.SeriesID.String
	}

	if payload.CoverURL.Valid() {
		vdb.CoverURL.Valid = true
		vdb.CoverURL.String = payload.CoverURL.String
	}

	if payload.Slug.Valid() {
		vdb.Slug = payload.Slug.String
	}

	if payload.Title.Valid() {
		vdb.Title = payload.Title.String
	}

	if payload.Status.Valid() {
		vdb.Status = payload.Status.String
	}

	if payload.SubTitle.Valid() {
		vdb.SubTitle.Valid = true
		vdb.SubTitle.String = payload.SubTitle.String
	}

	if payload.Timezone.Valid() {
		if _, err := time.LoadLocation(payload.Timezone.String); err == nil {
			vdb.Timezone = payload.Timezone.String
		}
	}
	return nil
}

func (v *ConferenceSvc) CreateDefaultSessionTypes(tx *db.Tx, c *model.Conference) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Conference.CreateDefaultSessionTypes").BindError(&err)
		defer g.End()
	}
	sst := SessionType()

	var stocktypes []model.AddSessionTypeRequest

	for _, dur := range []int{3600, 1800} {
		r := model.AddSessionTypeRequest{
			Name:     fmt.Sprintf("%d min", dur/60),
			Abstract: fmt.Sprintf("%d minute session", dur/60),
			Duration: dur,
		}
		r.LocalizedFields.Set("ja", "abstract", fmt.Sprintf("%d分枠", dur/60))
		stocktypes = append(stocktypes, r)
	}
	r := model.AddSessionTypeRequest{
		Name:     "Lightning Talk",
		Abstract: "5 minute session about anything you want",
		Duration: 300,
	}
	r.LocalizedFields.Set("ja", "abstract", "5分間で強制終了、初心者も安心枠")
	stocktypes = append(stocktypes, r)

	for _, r := range stocktypes {
		var vdb db.SessionType
		r.ConferenceID = c.ID
		if err := sst.Create(tx, &vdb, &model.CreateSessionTypeRequest{&r}); err != nil {
			return errors.Wrap(err, "failed to create default session type")
		}
	}
	return nil
}

func (v *ConferenceSvc) CreateFromPayload(tx *db.Tx, payload *model.CreateConferenceRequest, result *model.Conference) error {
	su := User()
	if err := su.IsConferenceSeriesAdministrator(tx, payload.SeriesID, payload.UserID); err != nil {
		return errors.Wrap(err, "creating a conference requires conference administrator privilege")
	}

	var vdb db.Conference
	if err := v.Create(tx, &vdb, payload); err != nil {
		return errors.Wrap(err, "failed to store in database")
	}

	// Description, CFPLead, CFPPresubmitInstructions, CFPPostsubmitInstruction
	// must be created
	cc := db.ConferenceComponent{
		ConferenceID: vdb.EID,
		CreatedOn:    time.Now(),
	}
	if payload.Description.Valid() && payload.Description.String != "" {
		cc.EID = tools.UUID()
		cc.Name = "description"
		cc.Value = payload.Description.String
		if err := cc.Create(tx); err != nil {
			return errors.Wrap(err, "failed to insert description")
		}
	}

	if payload.CFPLeadText.Valid() && payload.CFPLeadText.String != "" {
		cc.EID = tools.UUID()
		cc.Name = "cfp_lead_text"
		cc.Value = payload.CFPLeadText.String
		if err := cc.Create(tx); err != nil {
			return errors.Wrap(err, "failed to cfp lead text")
		}
	}

	if payload.CFPPostSubmitInstructions.Valid() && payload.CFPPostSubmitInstructions.String != "" {
		cc.EID = tools.UUID()
		cc.Name = "cfp_post_submit_instructions"
		cc.Value = payload.CFPPostSubmitInstructions.String
		if err := cc.Create(tx); err != nil {
			return errors.Wrap(err, "failed to insert cfp post-sumibt instructions")
		}
	}

	if payload.CFPPreSubmitInstructions.Valid() && payload.CFPPreSubmitInstructions.String != "" {
		cc.EID = tools.UUID()
		cc.Name = "cfp_pre_submit_instructions"
		cc.Value = payload.CFPPreSubmitInstructions.String
		if err := cc.Create(tx); err != nil {
			return errors.Wrap(err, "failed to insert cfp pre-submit instructions")
		}
	}

	if payload.ContactInformation.Valid() && payload.ContactInformation.String != "" {
		cc.EID = tools.UUID()
		cc.Name = "contact_information"
		cc.Value = payload.ContactInformation.String
		if err := cc.Create(tx); err != nil {
			return errors.Wrap(err, "failed to insert contact information")
		}
	}

	if err := v.AddAdministrator(tx, vdb.EID, payload.UserID); err != nil {
		return errors.Wrap(err, "failed to associate administrators to conference")
	}

	var c model.Conference
	if err := c.FromRow(vdb); err != nil {
		return errors.Wrap(err, "failed to populate model from database")
	}

	if err := v.CreateDefaultSessionTypes(tx, &c); err != nil {
		return errors.Wrap(err, "failed to create default session types")
	}

	*result = c
	return nil
}

var slugSplitRx = regexp.MustCompile(`^/([^/]+)/(.+)$`)

func (v *ConferenceSvc) LookupBySlug(tx *db.Tx, c *model.Conference, payload *model.LookupConferenceBySlugRequest) error {
	matches := slugSplitRx.FindStringSubmatch(payload.Slug)
	if matches == nil {
		return errors.New("invalid slug pattern")
	}
	seriesSlug := matches[1]
	confSlug := matches[2]

	// XXX cache this later!!!
	// This is in two steps so we can leverage existing vdb.LoadByEID()
	row := tx.QueryRow(`SELECT `+db.ConferenceTable+`.eid FROM `+db.ConferenceTable+` JOIN `+db.ConferenceSeriesTable+` ON `+db.ConferenceSeriesTable+`.eid = `+db.ConferenceTable+`.series_id WHERE `+db.ConferenceSeriesTable+`.slug = ? AND `+db.ConferenceTable+`.slug = ?`, seriesSlug, confSlug)

	var eid string
	if err := row.Scan(&eid); err != nil {
		return errors.Wrap(err, "failed to select conference id from slug")
	}

	return v.LookupFromPayload(tx, c, &model.LookupConferenceRequest{ID: eid, Lang: payload.Lang})
}

func (v *ConferenceSvc) AddAdministrator(tx *db.Tx, cid, uid string) error {
	c := db.ConferenceAdministrator{
		ConferenceID: cid,
		UserID:       uid,
	}
	return c.Create(tx, db.WithInsertIgnore(true))
}

func (v *ConferenceSvc) AddAdministratorFromPayload(tx *db.Tx, payload *model.AddConferenceAdminRequest) error {
	su := User()
	if err := su.IsConferenceAdministrator(tx, payload.ConferenceID, payload.UserID); err != nil {
		return errors.Wrap(err, "adding a conference administrator requires conference administrator privilege")
	}

	return errors.Wrap(v.AddAdministrator(tx, payload.ConferenceID, payload.AdminID), "failed to add administrator")
}

const datefmt = `2006-01-02`

func (v *ConferenceSvc) LoadByRange(tx *db.Tx, vdbl *db.ConferenceList, since, rangeStart, rangeEnd string, limit int) error {
	var rs time.Time
	var re time.Time
	var err error

	if rangeStart != "" {
		rs, err = time.Parse(datefmt, rangeStart)
		if err != nil {
			return err
		}
	}

	if rangeEnd != "" {
		re, err = time.Parse(datefmt, rangeEnd)
		if err != nil {
			return err
		}
	}

	if err := vdbl.LoadByRange(tx, since, rs, re, limit); err != nil {
		return err
	}

	return nil
}

func (v *ConferenceSvc) AddDatesFromPayload(tx *db.Tx, payload *model.CreateConferenceDateRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Conference.AddDatesFromPayload").BindError(&err)
		defer g.End()
	}

	su := User()
	if err := su.IsConferenceAdministrator(tx, payload.ConferenceID, payload.UserID); err != nil {
		return errors.Wrap(err, "adding conference dates requires conference administrator privilege")
	}

	var vdb db.ConferenceDate
	s := ConferenceDate()
	if err := s.Create(tx, &vdb, payload); err != nil {
		return errors.Wrap(err, "failed to insert into database")
	}

	return nil
}

func (v *ConferenceSvc) DeleteDateFromPayload(tx *db.Tx, payload *model.DeleteConferenceDateRequest) error {
	su := User()
	if err := su.IsConferenceAdministrator(tx, payload.ConferenceID, payload.UserID); err != nil {
		return errors.Wrap(err, "deleting conference dates requires conference administrator privilege")
	}

	var vdb db.ConferenceDate
	return vdb.DeleteDate(tx, payload.ConferenceID, payload.Date)
}

func (v *ConferenceSvc) LoadDates(tx *db.Tx, cdl *model.ConferenceDateList, cid string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Conference.LoadDates").BindError(&err)
		defer g.End()
	}

	var vdbl db.ConferenceDateList
	if err := vdbl.LoadByConferenceID(tx, cid); err != nil {
		return err
	}

	if pdebug.Enabled {
		pdebug.Printf("Loaded %d dates", len(vdbl))
	}

	res := make(model.ConferenceDateList, len(vdbl))
	for i, vdb := range vdbl {
		if vdb.Open.Valid {
			res[i].Open = vdb.Open.Time
		}

		if vdb.Close.Valid {
			res[i].Close = vdb.Close.Time
		}
	}
	*cdl = res
	return nil
}

func (v *ConferenceSvc) DeleteAdministratorFromPayload(tx *db.Tx, payload *model.DeleteConferenceAdminRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Conference.DeleteAdministratorFromPayload").BindError(&err)
		defer g.End()
	}

	su := User()
	if err := su.IsConferenceAdministrator(tx, payload.ConferenceID, payload.UserID); err != nil {
		return errors.Wrap(err, "deleting a conference administrator requires conference administrator privilege")
	}

	return db.DeleteConferenceAdministrator(tx, payload.ConferenceID, payload.AdminID)
}

func (v *ConferenceSvc) LoadAdmins(tx *db.Tx, cdl *model.UserList, trustedCall bool, cid, lang string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Conference.LoadAdmins").BindError(&err)
		defer g.End()
	}

	var vdbl db.UserList
	if err := db.LoadConferenceAdministrators(tx, &vdbl, cid); err != nil {
		return err
	}

	if pdebug.Enabled {
		pdebug.Printf("Loaded %d admins", len(vdbl))
	}

	res := make(model.UserList, len(vdbl))
	su := User()
	for i, vdb := range vdbl {
		if err := res[i].FromRow(vdb); err != nil {
			return errors.Wrap(err, "failed to map database to model")
		}
		if err := su.Decorate(tx, &res[i], trustedCall, lang); err != nil {
			return errors.Wrap(err, "failed to decorate administrator")
		}
	}
	*cdl = res
	return nil
}

func (v *ConferenceSvc) AddVenueFromPayload(tx *db.Tx, payload *model.AddConferenceVenueRequest) error {
	su := User()
	if err := su.IsConferenceAdministrator(tx, payload.ConferenceID, payload.UserID); err != nil {
		return errors.Wrap(err, "adding a conference venue requires conference administrator privilege")
	}
	cd := db.ConferenceVenue{
		ConferenceID: payload.ConferenceID,
		VenueID:      payload.VenueID,
	}
	if err := cd.Create(tx, db.WithInsertIgnore(true)); err != nil {
		return errors.Wrap(err, "failed to insert new conference/venue relation")
	}

	return nil
}

func (v *ConferenceSvc) DeleteVenueFromPayload(tx *db.Tx, payload *model.DeleteConferenceVenueRequest) error {
	su := User()
	if err := su.IsConferenceAdministrator(tx, payload.ConferenceID, payload.UserID); err != nil {
		return errors.Wrap(err, "deleting a conference venue requires conference administrator privilege")
	}
	return errors.Wrap(db.DeleteConferenceVenue(tx, payload.ConferenceID, payload.VenueID), "failed to delete conference venue")
}

func (v *ConferenceSvc) LoadVenues(tx *db.Tx, cdl *model.VenueList, cid string) error {
	var vdbl db.VenueList
	if err := db.LoadConferenceVenues(tx, &vdbl, cid); err != nil {
		return err
	}

	res := make(model.VenueList, len(vdbl))
	for i, vdb := range vdbl {
		var u model.Venue
		if err := u.FromRow(vdb); err != nil {
			return err
		}
		res[i] = u
	}
	*cdl = res
	return nil
}

func (v *ConferenceSvc) LoadTextComponents(tx *db.Tx, c *model.Conference) error {
	var ccl db.ConferenceComponentList

	if err := ccl.LoadByConferenceID(tx, c.ID); err != nil {
		return errors.Wrap(err, "failed to load text components for conference")
	}

	for _, cc := range ccl {
		switch cc.Name {
		case "description":
			c.Description = cc.Value
		case "cfp_lead_text":
			c.CFPLeadText = cc.Value
		case "cfp_pre_submit_instructions":
			c.CFPPreSubmitInstructions = cc.Value
		case "cfp_post_submit_instructions":
			c.CFPPostSubmitInstructions = cc.Value
		case "contact_information":
			c.ContactInformation = cc.Value
		}
	}
	return nil
}

func (v *ConferenceSvc) Decorate(tx *db.Tx, c *model.Conference, trustedCall bool, lang string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Conference.Decorate").BindError(&err)
		defer g.End()
	}

	if seriesID := c.SeriesID; seriesID != "" {
		var sdb db.ConferenceSeries
		if err := sdb.LoadByEID(tx, seriesID); err != nil {
			return errors.Wrapf(err, "failed to load conferences series '%s'", seriesID)
		}

		var s model.ConferenceSeries
		if err := s.FromRow(sdb); err != nil {
			return errors.Wrapf(err, "failed to load conferences series '%s'", seriesID)
		}
		c.Series = &s
		c.FullSlug = s.Slug + "/" + c.Slug
	}

	if c.CoverURL == "" {
		// TODO: fix later
		c.CoverURL = "https://builderscon.io/assets/images/heroimage.png"
	}

	if err := v.LoadTextComponents(tx, c); err != nil {
		return errors.Wrapf(err, "failed to load conference text components for '%s'", c.ID)
	}

	if err := v.LoadDates(tx, &c.Dates, c.ID); err != nil {
		return errors.Wrapf(err, "failed to load conference date for '%s'", c.ID)
	}

	if err := v.LoadAdmins(tx, &c.Administrators, trustedCall, c.ID, lang); err != nil {
		return errors.Wrapf(err, "failed to load administrators for '%s'", c.ID)
	}

	if err := v.LoadVenues(tx, &c.Venues, c.ID); err != nil {
		return errors.Wrapf(err, "failed to load venues for '%s'", c.ID)
	}

	if err := v.LoadFeaturedSpeakers(tx, &c.FeaturedSpeakers, c.ID); err != nil {
		return errors.Wrapf(err, "failed to load featured speakers for '%s'", c.ID)
	}

	if err := v.LoadSponsors(tx, &c.Sponsors, c.ID); err != nil {
		return errors.Wrapf(err, "failed to load sponsors for '%s'", c.ID)
	}

	if err := v.LoadSessionTypes(tx, &c.SessionTypes, c.ID); err != nil {
		return errors.Wrapf(err, "failed to load session types for '%s'", c.ID)
	}

	sv := Venue()
	for i := range c.Venues {
		if err := sv.Decorate(tx, &c.Venues[i], trustedCall, lang); err != nil {
			return errors.Wrap(err, "failed to decorate venue with associated data")
		}
	}

	sfs := FeaturedSpeaker()
	for i := range c.FeaturedSpeakers {
		if err := sfs.Decorate(tx, &c.FeaturedSpeakers[i], trustedCall, lang); err != nil {
			return errors.Wrap(err, "failed to decorate featured speakers with associated data")
		}
	}

	sps := Sponsor()
	for i := range c.Sponsors {
		if err := sps.Decorate(tx, &c.Sponsors[i], trustedCall, lang); err != nil {
			return errors.Wrap(err, "failed to decorate sponsors with associated data")
		}
	}

	sts := SessionType()
	for i := range c.SessionTypes {
		if err := sts.Decorate(tx, &c.SessionTypes[i], trustedCall, lang); err != nil {
			return errors.Wrap(err, "failed to decorate session types with associated data")
		}
	}

	switch lang {
	case "", "en":
	default:
		if err := v.ReplaceL10NStrings(tx, c, lang); err != nil {
			return errors.Wrap(err, "failed to replace L10N strings")
		}
	}

	return nil
}

func (v *ConferenceSvc) UploadImagesFromPayload(ctx context.Context, tx *db.Tx, payload *model.UpdateConferenceRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Conference.UploadImagesFromPayload").BindError(&err)
		defer g.End()
	}

	// There's nothing to do
	if payload.MultipartForm == nil || payload.MultipartForm.File == nil {
		return nil
	}

	field := "cover"
	fhs := payload.MultipartForm.File[field]
	if len(fhs) == 0 {
		return nil
	}

	var imgf multipart.File
	imgf, err = fhs[0].Open()
	if err != nil {
		return errors.Wrap(err, "failed to open cover file from multipart form")
	}

	var imgbuf bytes.Buffer
	if _, err := io.Copy(&imgbuf, imgf); err != nil {
		return errors.Wrap(err, "failed to copy cover image data to memory")
	}
	ct := http.DetectContentType(imgbuf.Bytes())

	// Only work with image/png or image/jpeg
	var suffix string
	switch ct {
	case "image/png":
		suffix = "png"
	case "image/jpeg":
		suffix = "jpeg"
	default:
		return errors.Errorf("Unsupported image type %s", ct)
	}

	// TODO: Validate the image
	// TODO: Avoid Google Storage hardcoding?
	// Upload this to a temporary location, then upon successful write to DB
	// rename it to $conference_id/$sponsor_id
	tmpname := time.Now().UTC().Format("2006-01-02") + "/" + tools.RandomString(64) + "." + suffix
	cl := MediaStorage
	err = cl.Upload(ctx, tmpname, &imgbuf, WithObjectAttrs(storage.ObjectAttrs{
		ContentType: ct,
		ACL: []storage.ACLRule{
			{storage.AllUsers, storage.RoleReader},
		},
	}))
	if err != nil {
		return errors.Wrap(err, "failed to upload file")
	}

	if pdebug.Enabled {
		pdebug.Printf("Writing '%s' to %s", field, tmpname)
	}

	dstname := "conferences/" + payload.ID + "/cover." + suffix
	payload.CoverURL.Set(cl.URLFor(dstname))

	return finalizeFunc(func() (err error) {
		if pdebug.Enabled {
			g := pdebug.Marker("Finalizer for service.Conference.UploadImagesFromPayload").BindError(&err)
			defer g.End()
		}
		return cl.Move(ctx, tmpname, dstname)
	})
}

func (v *ConferenceSvc) UpdateFromPayload(ctx context.Context, tx *db.Tx, payload *model.UpdateConferenceRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Conference.UpdateFromPayload (%s)", payload.ID).BindError(&err)
		defer g.End()
	}

	su := User()
	if err := su.IsConferenceAdministrator(tx, payload.ID, payload.UserID); err != nil {
		return errors.Wrap(err, "updating a conference requires conference administrator privilege")
	}

	var vdb db.Conference
	if err := vdb.LoadByEID(tx, payload.ID); err != nil {
		return errors.Wrap(err, `failed to load from database`)
	}

	if err := v.populateRowForUpdate(&vdb, payload); err != nil {
		return err
	}

	uploadErr := v.UploadImagesFromPayload(ctx, tx, payload)
	if !errors.IsIgnorable(uploadErr) {
		return errors.Wrap(uploadErr, "failed to process image uploads")
	}

	if err := v.Update(tx, &vdb); err != nil {
		return errors.Wrap(err, "failed to update databased")
	}

	err = payload.LocalizedFields.Foreach(func(l, k, x string) error {
		if pdebug.Enabled {
			pdebug.Printf("Updating l10n string for '%s' (%s)", k, l)
		}
		ls := db.LocalizedString{
			ParentType: "Conference",
			ParentID:   vdb.EID,
			Language:   l,
			Name:       k,
			Localized:  x,
		}
		return ls.Upsert(tx)
	})
	if err != nil {
		return errors.Wrap(err, "failed to update localized fields")
	}
	var ccs ConferenceComponentSvc
	deletedTextComponents := []string{}
	addedTextComponents := map[string]string{}
	if payload.Description.Valid() {
		s := payload.Description.String
		n := "description"
		if len(s) == 0 {
			deletedTextComponents = append(deletedTextComponents, n)
		} else {
			addedTextComponents[n] = s
		}
	}

	if payload.CFPLeadText.Valid() {
		s := payload.CFPLeadText.String
		n := "cfp_lead_text"
		if len(s) == 0 {
			deletedTextComponents = append(deletedTextComponents, n)
		} else {
			addedTextComponents[n] = s
		}
	}

	if payload.CFPPreSubmitInstructions.Valid() {
		s := payload.CFPPreSubmitInstructions.String
		n := "cfp_pre_submit_instructions"
		if len(s) == 0 {
			deletedTextComponents = append(deletedTextComponents, n)
		} else {
			addedTextComponents[n] = s
		}
	}

	if payload.CFPPostSubmitInstructions.Valid() {
		s := payload.CFPPostSubmitInstructions.String
		n := "cfp_post_submit_instructions"
		if len(s) == 0 {
			deletedTextComponents = append(deletedTextComponents, n)
		} else {
			addedTextComponents[n] = s
		}
	}

	if payload.ContactInformation.Valid() {
		s := payload.ContactInformation.String
		n := "contact_information"
		if len(s) == 0 {
			deletedTextComponents = append(deletedTextComponents, n)
		} else {
			addedTextComponents[n] = s
		}
	}

	if len(deletedTextComponents) > 0 {
		if err := ccs.DeleteByConferenceIDAndName(tx, payload.ID, deletedTextComponents...); err != nil {
			return errors.Wrap(err, "failed to delete components")
		}
	}

	if len(addedTextComponents) > 0 {
		if err := ccs.UpsertByConferenceIDAndName(tx, payload.ID, addedTextComponents); err != nil {
			return errors.Wrap(err, "failed to register components")
		}
	}

	if _, ok := errors.IsFinalizationRequired(uploadErr); ok {
		return uploadErr
	}

	return nil
}

func (v *ConferenceSvc) ListFromPayload(tx *db.Tx, l *model.ConferenceList, payload *model.ListConferenceRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Conference.ListFromPayload").BindError(&err)
		defer g.End()
	}

	var rs time.Time
	var re time.Time

	if payload.RangeStart.Valid() {
		if rs, err = time.Parse(datefmt, payload.RangeStart.String); err != nil {
			return errors.Wrap(err, "failed to parse range start")
		}
	}

	if payload.RangeEnd.Valid() {
		if re, err = time.Parse(datefmt, payload.RangeEnd.String); err != nil {
			return errors.Wrap(err, "failed to parse range end")
		}
	}

	status := payload.Status
	switch len(status) {
	case 0:
		status = []string{"public"}
	case 1:
		if status[0] == "any" {
			status = []string{"public", "private"}
		}
	}

	var vdbl db.ConferenceList
	if err := vdbl.LoadFromQuery(tx, status, payload.Organizers, rs, re, payload.Since.String, int(payload.Limit.Int)); err != nil {
		return errors.Wrap(err, "failed to load list from database")
	}

	r := make(model.ConferenceList, len(vdbl))
	for i, vdb := range vdbl {
		if err := (r[i]).FromRow(vdb); err != nil {
			return errors.Wrap(err, "failed populate model from database")
		}
		if err := v.Decorate(tx, &r[i], false, payload.Lang.String); err != nil {
			return errors.Wrap(err, "failed to decorate venue with associated data")
		}
	}

	*l = r
	return nil
}

func (v *ConferenceSvc) LoadFeaturedSpeakers(tx *db.Tx, cdl *model.FeaturedSpeakerList, cid string) error {
	var vdbl db.FeaturedSpeakerList
	if err := db.LoadFeaturedSpeakers(tx, &vdbl, cid); err != nil {
		return err
	}

	res := make(model.FeaturedSpeakerList, len(vdbl))
	for i, vdb := range vdbl {
		var u model.FeaturedSpeaker
		if err := u.FromRow(vdb); err != nil {
			return err
		}
		res[i] = u
	}
	*cdl = res
	return nil
}

func (v *ConferenceSvc) LoadSponsors(tx *db.Tx, cdl *model.SponsorList, cid string) error {
	var vdbl db.SponsorList
	if err := db.LoadSponsors(tx, &vdbl, cid); err != nil {
		return err
	}

	res := make(model.SponsorList, len(vdbl))
	for i, vdb := range vdbl {
		var u model.Sponsor
		if err := u.FromRow(vdb); err != nil {
			return err
		}
		res[i] = u
	}
	*cdl = res
	return nil
}

func (v *ConferenceSvc) LoadSessionTypes(tx *db.Tx, cdl *model.SessionTypeList, cid string) error {
	var vdbl db.SessionTypeList
	if err := db.LoadSessionTypes(tx, &vdbl, cid); err != nil {
		return err
	}

	res := make(model.SessionTypeList, len(vdbl))
	for i, vdb := range vdbl {
		var u model.SessionType
		if err := u.FromRow(vdb); err != nil {
			return err
		}
		res[i] = u
	}
	*cdl = res
	return nil
}

func (v *ConferenceSvc) ListByOrganizerFromPayload(tx *db.Tx, l *model.ConferenceList, payload *model.ListConferencesByOrganizerRequest) (err error) {
	var vdbl db.ConferenceList
	if err := vdbl.LoadFromQuery(tx, payload.Status, payload.OrganizerID, time.Time{}, time.Time{}, payload.Since.String, int(payload.Limit.Int)); err != nil {
		return err
	}

	res := make(model.ConferenceList, len(vdbl))
	for i, vdb := range vdbl {
		if err := (res[i]).FromRow(vdb); err != nil {
			return errors.Wrap(err, "failed populate model from database")
		}
		if err := v.Decorate(tx, &res[i], false, payload.Lang.String); err != nil {
			return errors.Wrap(err, "failed to decorate conference with associated data")
		}
	}
	*l = res
	return nil

}

func (v *ConferenceSvc) TweetFromPayload(ctx context.Context, tx *db.Tx, payload *model.TweetAsConferenceRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Conference.Tweet").BindError(&err)
		defer g.End()
	}

	// You have to be conference admin to do this
	su := User()
	if err := su.IsConferenceAdministrator(tx, payload.ConferenceID, payload.UserID); err != nil {
		return errors.Wrap(err, "adding a conference credentials requires conference administrator privilege")
	}

	return Twitter().TweetAsConference(payload.ConferenceID, payload.Tweet)
}

func (v *ConferenceSvc) AddCredentialFromPayload(ctx context.Context, tx *db.Tx, payload *model.AddConferenceCredentialRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Conference.AddCredentialFromPayload").BindError(&err)
		defer g.End()
	}

	// You have to be conference admin to do this
	su := User()
	if err := su.IsConferenceAdministrator(tx, payload.ConferenceID, payload.UserID); err != nil {
		return errors.Wrap(err, "adding a conference credentials requires conference administrator privilege")
	}

	// Data is in text, and it must be base64
	decoded, err := base64.StdEncoding.DecodeString(payload.Data)
	if err != nil {
		return errors.Wrap(err, "failed to decode payload.Data with base64")
	}

	cl := v.credentialStorage
	name := "conferences/" + payload.ConferenceID + "/credentials/" + payload.Type
	err = cl.Upload(ctx, name, bytes.NewBuffer(decoded), WithObjectAttrs(storage.ObjectAttrs{
		ContentType: "text/plain",
		// Note: DO NOT ADD PERMISSIVE ACLS HERE!
	}))
	if err != nil {
		return errors.Wrap(err, "failed to upload file")
	}

	return nil
}
