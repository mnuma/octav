package service

// Automatically generated by genmodel utility. DO NOT EDIT!

import (
	"time"

	"github.com/builderscon/octav/octav/db"
	"github.com/builderscon/octav/octav/model"
	"github.com/lestrrat/go-pdebug"
	"github.com/pkg/errors"
)

var _ = time.Time{}

func (v *Room) Lookup(tx *db.Tx, m *model.Room, payload model.LookupRoomRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Room.Lookup").BindError(&err)
		defer g.End()
	}

	r := model.Room{}
	if err = r.Load(tx, payload.ID); err != nil {
		return errors.Wrap(err, "failed to load model.Room from database")
	}
	*m = r
	return nil
}

// Create takes in the transaction, the incoming payload, and a reference to
// a database row. The database row is initialized/populated so that the
// caller can use it afterwards.
func (v *Room) Create(tx *db.Tx, vdb *db.Room, payload model.CreateRoomRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Room.Create").BindError(&err)
		defer g.End()
	}

	if err := v.populateRowForCreate(vdb, payload); err != nil {
		return err
	}

	if err := vdb.Create(tx); err != nil {
		return err
	}

	if err := payload.L10N.CreateLocalizedStrings(tx, "Room", vdb.EID); err != nil {
		return err
	}
	return nil
}

func (v *Room) Update(tx *db.Tx, vdb *db.Room, payload model.UpdateRoomRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Room.Update (%s)", vdb.EID).BindError(&err)
		defer g.End()
	}

	if vdb.EID == "" {
		return errors.New("vdb.EID is required (did you forget to call vdb.Load(tx) before hand?)")
	}

	if err := v.populateRowForUpdate(vdb, payload); err != nil {
		return err
	}

	if err := vdb.Update(tx); err != nil {
		return err
	}

	return payload.L10N.Foreach(func(l, k, x string) error {
		if pdebug.Enabled {
			pdebug.Printf("Updating l10n string for '%s' (%s)", k, l)
		}
		ls := db.LocalizedString{
			ParentType: "Room",
			ParentID:   vdb.EID,
			Language:   l,
			Name:       k,
			Localized:  x,
		}
		return ls.Upsert(tx)
	})
}

func (v *Room) ReplaceL10NStrings(tx *db.Tx, m *model.Room, lang string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Room.ReplaceL10NStrings")
		defer g.End()
	}
	rows, err := tx.Query(`SELECT oid, parent_id, parent_type, name, language, localized FROM localized_strings WHERE parent_type = ? AND parent_id = ? AND language = ?`, "Room", m.ID, lang)
	if err != nil {
		return err
	}

	var l db.LocalizedString
	for rows.Next() {
		if err := l.Scan(rows); err != nil {
			return err
		}

		switch l.Name {
		case "name":
			if pdebug.Enabled {
				pdebug.Printf("Replacing for key 'name'")
			}
			m.Name = l.Localized
		}
	}
	return nil
}

func (v *Room) Delete(tx *db.Tx, id string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("Room.Delete (%s)", id)
		defer g.End()
	}

	vdb := db.Room{EID: id}
	if err := vdb.Delete(tx); err != nil {
		return err
	}
	if err := db.DeleteLocalizedStringsForParent(tx, id, "Room"); err != nil {
		return err
	}
	return nil
}

func (v *Room) LoadList(tx *db.Tx, vdbl *db.RoomList, since string, limit int) error {
	return vdbl.LoadSinceEID(tx, since, limit)
}
