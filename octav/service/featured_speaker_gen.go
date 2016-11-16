package service

// Automatically generated by genmodel utility. DO NOT EDIT!

import (
	"context"
	"sync"
	"time"

	"github.com/builderscon/octav/octav/cache"

	"github.com/builderscon/octav/octav/db"
	"github.com/builderscon/octav/octav/internal/errors"
	"github.com/builderscon/octav/octav/model"
	"github.com/lestrrat/go-pdebug"
)

var _ = time.Time{}
var _ = cache.WithExpires(time.Minute)
var _ = context.Background
var _ = errors.Wrap
var _ = model.FeaturedSpeaker{}
var _ = db.FeaturedSpeaker{}
var _ = pdebug.Enabled

var featuredSpeakerSvc FeaturedSpeakerSvc
var featuredSpeakerOnce sync.Once

func FeaturedSpeaker() *FeaturedSpeakerSvc {
	featuredSpeakerOnce.Do(featuredSpeakerSvc.Init)
	return &featuredSpeakerSvc
}

func (v *FeaturedSpeakerSvc) LookupFromPayload(tx *db.Tx, m *model.FeaturedSpeaker, payload *model.LookupFeaturedSpeakerRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.FeaturedSpeaker.LookupFromPayload").BindError(&err)
		defer g.End()
	}
	if err = v.Lookup(tx, m, payload.ID); err != nil {
		return errors.Wrap(err, "failed to load model.FeaturedSpeaker from database")
	}
	if err := v.Decorate(tx, m, payload.TrustedCall, payload.Lang.String); err != nil {
		return errors.Wrap(err, "failed to load associated data for model.FeaturedSpeaker from database")
	}
	return nil
}

func (v *FeaturedSpeakerSvc) Lookup(tx *db.Tx, m *model.FeaturedSpeaker, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.FeaturedSpeaker.Lookup").BindError(&err)
		defer g.End()
	}

	var r model.FeaturedSpeaker
	key := `api.FeaturedSpeaker.` + id
	c := Cache()
	var cacheMiss bool
	_, err = c.GetOrSet(key, &r, func() (interface{}, error) {
		if pdebug.Enabled {
			cacheMiss = true
		}
		if err := r.Load(tx, id); err != nil {
			return nil, errors.Wrap(err, "failed to load model.FeaturedSpeaker from database")
		}
		return &r, nil
	}, cache.WithExpires(time.Hour))
	if pdebug.Enabled {
		cacheSt := `HIT`
		if cacheMiss {
			cacheSt = `MISS`
		}
		pdebug.Printf(`CACHE %s: %s`, cacheSt, key)
	}
	*m = r
	return nil
}

// Create takes in the transaction, the incoming payload, and a reference to
// a database row. The database row is initialized/populated so that the
// caller can use it afterwards.
func (v *FeaturedSpeakerSvc) Create(tx *db.Tx, vdb *db.FeaturedSpeaker, payload *model.CreateFeaturedSpeakerRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.FeaturedSpeaker.Create").BindError(&err)
		defer g.End()
	}

	if err := v.populateRowForCreate(vdb, payload); err != nil {
		return errors.Wrap(err, `failed to populate row`)
	}

	if err := vdb.Create(tx); err != nil {
		return errors.Wrap(err, `failed to insert into database`)
	}

	if err := payload.LocalizedFields.CreateLocalizedStrings(tx, "FeaturedSpeaker", vdb.EID); err != nil {
		return errors.Wrap(err, `failed to populate localized strings`)
	}
	return nil
}

func (v *FeaturedSpeakerSvc) Update(tx *db.Tx, vdb *db.FeaturedSpeaker) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.FeaturedSpeaker.Update (%s)", vdb.EID).BindError(&err)
		defer g.End()
	}

	if vdb.EID == `` {
		return errors.New("vdb.EID is required (did you forget to call vdb.Load(tx) before hand?)")
	}

	if err := vdb.Update(tx); err != nil {
		return errors.Wrap(err, `failed to update database`)
	}
	key := `api.FeaturedSpeaker.` + vdb.EID
	if pdebug.Enabled {
		pdebug.Printf(`CACHE DEL %s`, key)
	}
	c := Cache()
	cerr := c.Delete(key)
	if pdebug.Enabled {
		if cerr != nil {
			pdebug.Printf(`CACHE ERR: %%s`, cerr)
		}
	}
	return nil
}

func (v *FeaturedSpeakerSvc) UpdateFromPayload(ctx context.Context, tx *db.Tx, payload *model.UpdateFeaturedSpeakerRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.FeaturedSpeaker.UpdateFromPayload (%s)", payload.ID).BindError(&err)
		defer g.End()
	}
	var vdb db.FeaturedSpeaker
	if err := vdb.LoadByEID(tx, payload.ID); err != nil {
		return errors.Wrap(err, `failed to load from database`)
	}

	if err := v.PreUpdateFromPayloadHook(ctx, tx, &vdb, payload); err != nil {
		return errors.Wrap(err, `failed to execute PreUpdateFromPayloadHook`)
	}

	if err := v.populateRowForUpdate(&vdb, payload); err != nil {
		return errors.Wrap(err, `failed to populate row data`)
	}

	if err := v.Update(tx, &vdb); err != nil {
		return errors.Wrap(err, `failed to update row in database`)
	}

	return payload.LocalizedFields.Foreach(func(l, k, x string) error {
		if pdebug.Enabled {
			pdebug.Printf("Updating l10n string for '%s' (%s)", k, l)
		}
		ls := db.LocalizedString{
			ParentType: "FeaturedSpeaker",
			ParentID:   vdb.EID,
			Language:   l,
			Name:       k,
			Localized:  x,
		}
		return ls.Upsert(tx)
	})
}

func (v *FeaturedSpeakerSvc) ReplaceL10NStrings(tx *db.Tx, m *model.FeaturedSpeaker, lang string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("service.FeaturedSpeaker.ReplaceL10NStrings lang = %s", lang)
		defer g.End()
	}
	switch lang {
	case "", "en":
		if len(m.DisplayName) > 0 && len(m.Description) > 0 {
			return nil
		}
		for _, extralang := range []string{`ja`} {
			rows, err := tx.Query(`SELECT oid, parent_id, parent_type, name, language, localized FROM localized_strings WHERE parent_type = ? AND parent_id = ? AND language = ?`, "FeaturedSpeaker", m.ID, extralang)
			if err != nil {
				if errors.IsSQLNoRows(err) {
					break
				}
				return errors.Wrap(err, `failed to excute query`)
			}

			var l db.LocalizedString
			for rows.Next() {
				if err := l.Scan(rows); err != nil {
					return err
				}
				if len(l.Localized) == 0 {
					continue
				}
				switch l.Name {
				case "display_name":
					if len(m.DisplayName) == 0 {
						if pdebug.Enabled {
							pdebug.Printf("Replacing for key 'display_name' (fallback en -> %s", l.Language)
						}
						m.DisplayName = l.Localized
					}
				case "description":
					if len(m.Description) == 0 {
						if pdebug.Enabled {
							pdebug.Printf("Replacing for key 'description' (fallback en -> %s", l.Language)
						}
						m.Description = l.Localized
					}
				}
			}
		}
		return nil
	case "all":
		rows, err := tx.Query(`SELECT oid, parent_id, parent_type, name, language, localized FROM localized_strings WHERE parent_type = ? AND parent_id = ?`, "FeaturedSpeaker", m.ID)
		if err != nil {
			return err
		}

		var l db.LocalizedString
		for rows.Next() {
			if err := l.Scan(rows); err != nil {
				return err
			}
			if len(l.Localized) == 0 {
				continue
			}
			if pdebug.Enabled {
				pdebug.Printf("Adding key '%s#%s'", l.Name, l.Language)
			}
			m.LocalizedFields.Set(l.Language, l.Name, l.Localized)
		}
	default:
		rows, err := tx.Query(`SELECT oid, parent_id, parent_type, name, language, localized FROM localized_strings WHERE parent_type = ? AND parent_id = ? AND language = ?`, "FeaturedSpeaker", m.ID, lang)
		if err != nil {
			return err
		}

		var l db.LocalizedString
		for rows.Next() {
			if err := l.Scan(rows); err != nil {
				return err
			}
			if len(l.Localized) == 0 {
				continue
			}

			switch l.Name {
			case "display_name":
				if pdebug.Enabled {
					pdebug.Printf("Replacing for key 'display_name'")
				}
				m.DisplayName = l.Localized
			case "description":
				if pdebug.Enabled {
					pdebug.Printf("Replacing for key 'description'")
				}
				m.Description = l.Localized
			}
		}
	}
	return nil
}

func (v *FeaturedSpeakerSvc) Delete(tx *db.Tx, id string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("FeaturedSpeaker.Delete (%s)", id)
		defer g.End()
	}

	vdb := db.FeaturedSpeaker{EID: id}
	if err := vdb.Delete(tx); err != nil {
		return err
	}
	key := `api.FeaturedSpeaker.` + id
	c := Cache()
	c.Delete(key)
	if pdebug.Enabled {
		pdebug.Printf(`CACHE DEL %s`, key)
	}
	if err := db.DeleteLocalizedStringsForParent(tx, id, "FeaturedSpeaker"); err != nil {
		return err
	}
	return nil
}
