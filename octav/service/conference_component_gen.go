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
var _ = model.ConferenceComponent{}
var _ = db.ConferenceComponent{}
var _ = pdebug.Enabled

var conferenceComponentSvc ConferenceComponentSvc
var conferenceComponentOnce sync.Once

func ConferenceComponent() *ConferenceComponentSvc {
	conferenceComponentOnce.Do(conferenceComponentSvc.Init)
	return &conferenceComponentSvc
}

func (v *ConferenceComponentSvc) LookupFromPayload(tx *db.Tx, m *model.ConferenceComponent, payload *model.LookupConferenceComponentRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ConferenceComponent.LookupFromPayload").BindError(&err)
		defer g.End()
	}
	if err = v.Lookup(tx, m, payload.ID); err != nil {
		return errors.Wrap(err, "failed to load model.ConferenceComponent from database")
	}
	return nil
}

func (v *ConferenceComponentSvc) Lookup(tx *db.Tx, m *model.ConferenceComponent, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ConferenceComponent.Lookup").BindError(&err)
		defer g.End()
	}

	var r model.ConferenceComponent
	key := `api.ConferenceComponent.` + id
	c := Cache()
	var cacheMiss bool
	_, err = c.GetOrSet(key, &r, func() (interface{}, error) {
		if pdebug.Enabled {
			cacheMiss = true
		}
		if err := r.Load(tx, id); err != nil {
			return nil, errors.Wrap(err, "failed to load model.ConferenceComponent from database")
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
func (v *ConferenceComponentSvc) Create(tx *db.Tx, vdb *db.ConferenceComponent, payload *model.CreateConferenceComponentRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ConferenceComponent.Create").BindError(&err)
		defer g.End()
	}

	if err := v.populateRowForCreate(vdb, payload); err != nil {
		return errors.Wrap(err, `failed to populate row`)
	}

	if err := vdb.Create(tx); err != nil {
		return errors.Wrap(err, `failed to insert into database`)
	}

	return nil
}

func (v *ConferenceComponentSvc) Update(tx *db.Tx, vdb *db.ConferenceComponent) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ConferenceComponent.Update (%s)", vdb.EID).BindError(&err)
		defer g.End()
	}

	if vdb.EID == `` {
		return errors.New("vdb.EID is required (did you forget to call vdb.Load(tx) before hand?)")
	}

	if err := vdb.Update(tx); err != nil {
		return errors.Wrap(err, `failed to update database`)
	}
	key := `api.ConferenceComponent.` + vdb.EID
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

func (v *ConferenceComponentSvc) UpdateFromPayload(ctx context.Context, tx *db.Tx, payload *model.UpdateConferenceComponentRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ConferenceComponent.UpdateFromPayload (%s)", payload.ID).BindError(&err)
		defer g.End()
	}
	var vdb db.ConferenceComponent
	if err := vdb.LoadByEID(tx, payload.ID); err != nil {
		return errors.Wrap(err, `failed to load from database`)
	}

	if err := v.populateRowForUpdate(&vdb, payload); err != nil {
		return errors.Wrap(err, `failed to populate row data`)
	}

	if err := v.Update(tx, &vdb); err != nil {
		return errors.Wrap(err, `failed to update row in database`)
	}
	return nil
}

func (v *ConferenceComponentSvc) Delete(tx *db.Tx, id string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("ConferenceComponent.Delete (%s)", id)
		defer g.End()
	}

	vdb := db.ConferenceComponent{EID: id}
	if err := vdb.Delete(tx); err != nil {
		return err
	}
	key := `api.ConferenceComponent.` + id
	c := Cache()
	c.Delete(key)
	if pdebug.Enabled {
		pdebug.Printf(`CACHE DEL %s`, key)
	}
	return nil
}
