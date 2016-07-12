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

func (v *ConferenceSeries) Lookup(tx *db.Tx, m *model.ConferenceSeries, payload model.LookupConferenceSeriesRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ConferenceSeries.Lookup").BindError(&err)
		defer g.End()
	}

	r := model.ConferenceSeries{}
	if err = r.Load(tx, payload.ID); err != nil {
		return errors.Wrap(err, "failed to load model.ConferenceSeries from database")
	}
	*m = r
	return nil
}

// Create takes in the transaction, the incoming payload, and a reference to
// a database row. The database row is initialized/populated so that the
// caller can use it afterwards.
func (v *ConferenceSeries) Create(tx *db.Tx, vdb *db.ConferenceSeries, payload model.CreateConferenceSeriesRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ConferenceSeries.Create").BindError(&err)
		defer g.End()
	}

	if err := v.populateRowForCreate(vdb, payload); err != nil {
		return err
	}

	if err := vdb.Create(tx); err != nil {
		return err
	}

	return nil
}

func (v *ConferenceSeries) Update(tx *db.Tx, vdb *db.ConferenceSeries, payload model.UpdateConferenceSeriesRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ConferenceSeries.Update (%s)", vdb.EID).BindError(&err)
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
	return nil
}

func (v *ConferenceSeries) Delete(tx *db.Tx, id string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("ConferenceSeries.Delete (%s)", id)
		defer g.End()
	}

	vdb := db.ConferenceSeries{EID: id}
	if err := vdb.Delete(tx); err != nil {
		return err
	}
	return nil
}

func (v *ConferenceSeries) LoadList(tx *db.Tx, vdbl *db.ConferenceSeriesList, since string, limit int) error {
	return vdbl.LoadSinceEID(tx, since, limit)
}
