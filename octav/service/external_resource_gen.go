package service

// Automatically generated by genmodel utility. DO NOT EDIT!

import (
	"context"
	"database/sql"
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
var _ = model.ExternalResource{}
var _ = db.ExternalResource{}
var _ = sql.ErrNoRows
var _ = pdebug.Enabled

var externalResourceSvc ExternalResourceSvc
var externalResourceOnce sync.Once

func ExternalResource() *ExternalResourceSvc {
	externalResourceOnce.Do(externalResourceSvc.Init)
	return &externalResourceSvc
}

func (v *ExternalResourceSvc) LookupFromPayload(ctx context.Context, tx *sql.Tx, m *model.ExternalResource, payload *model.LookupExternalResourceRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ExternalResource.LookupFromPayload %s", payload.ID).BindError(&err)
		defer g.End()
	}
	if err = v.Lookup(ctx, tx, m, payload.ID); err != nil {
		return errors.Wrap(err, "failed to load model.ExternalResource from database")
	}
	if err := v.Decorate(ctx, tx, m, payload.TrustedCall, payload.Lang.String); err != nil {
		return errors.Wrap(err, "failed to load associated data for model.ExternalResource from database")
	}
	return nil
}

func (v *ExternalResourceSvc) Lookup(ctx context.Context, tx *sql.Tx, m *model.ExternalResource, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ExternalResource.Lookup %s", id).BindError(&err)
		defer g.End()
	}

	var r model.ExternalResource
	c := Cache()
	key := c.Key("ExternalResource", id)
	var cacheMiss bool
	_, err = c.GetOrSet(key, &r, func() (interface{}, error) {
		if pdebug.Enabled {
			cacheMiss = true
		}
		if err := r.Load(tx, id); err != nil {
			return nil, errors.Wrap(err, "failed to load model.ExternalResource from database")
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
func (v *ExternalResourceSvc) Create(ctx context.Context, tx *sql.Tx, vdb *db.ExternalResource, payload *model.CreateExternalResourceRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ExternalResource.Create").BindError(&err)
		defer g.End()
	}

	if err := v.populateRowForCreate(ctx, vdb, payload); err != nil {
		return errors.Wrap(err, `failed to populate row`)
	}

	if err := vdb.Create(tx, payload.DatabaseOptions...); err != nil {
		return errors.Wrap(err, `failed to insert into database`)
	}

	if err := payload.LocalizedFields.CreateLocalizedStrings(tx, "ExternalResource", vdb.EID); err != nil {
		return errors.Wrap(err, `failed to populate localized strings`)
	}
	if err := v.PostCreateHook(ctx, tx, vdb); err != nil {
		return errors.Wrap(err, `post create hook failed`)
	}
	return nil
}

func (v *ExternalResourceSvc) Update(tx *sql.Tx, vdb *db.ExternalResource) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ExternalResource.Update (%s)", vdb.EID).BindError(&err)
		defer g.End()
	}

	if vdb.EID == `` {
		return errors.New("vdb.EID is required (did you forget to call vdb.Load(tx) before hand?)")
	}

	if err := vdb.Update(tx); err != nil {
		return errors.Wrap(err, `failed to update database`)
	}
	c := Cache()
	key := c.Key("ExternalResource", vdb.EID)
	if pdebug.Enabled {
		pdebug.Printf(`CACHE DEL %s`, key)
	}
	cerr := c.Delete(key)
	if pdebug.Enabled {
		if cerr != nil {
			pdebug.Printf(`CACHE ERR: %s`, cerr)
		}
	}
	if err := v.PostUpdateHook(tx, vdb); err != nil {
		return errors.Wrap(err, `post update hook failed`)
	}
	return nil
}

func (v *ExternalResourceSvc) UpdateFromPayload(ctx context.Context, tx *sql.Tx, payload *model.UpdateExternalResourceRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ExternalResource.UpdateFromPayload (%s)", payload.ID).BindError(&err)
		defer g.End()
	}
	var vdb db.ExternalResource
	if err := vdb.LoadByEID(tx, payload.ID); err != nil {
		return errors.Wrap(err, `failed to load from database`)
	}

	if err := v.populateRowForUpdate(ctx, &vdb, payload); err != nil {
		return errors.Wrap(err, `failed to populate row data`)
	}

	if err := v.Update(tx, &vdb); err != nil {
		return errors.Wrap(err, `failed to update row in database`)
	}

	ls := LocalizedString()
	if err := ls.UpdateFields(tx, "ExternalResource", vdb.EID, payload.LocalizedFields); err != nil {
		return errors.Wrap(err, `failed to update localized fields`)
	}
	return nil
}

func (v *ExternalResourceSvc) ReplaceL10NStrings(tx *sql.Tx, m *model.ExternalResource, lang string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("service.ExternalResource.ReplaceL10NStrings lang = %s", lang)
		defer g.End()
	}
	ls := LocalizedString()
	list := make([]db.LocalizedString, 0, 2)
	switch lang {
	case "", "en":
		if len(m.Description) > 0 && len(m.Title) > 0 {
			return nil
		}
		for _, extralang := range []string{`ja`} {
			list = list[:0]
			if err := ls.LookupFields(tx, "ExternalResource", m.ID, extralang, &list); err != nil {
				return errors.Wrap(err, `failed to lookup localized fields`)
			}

			for _, l := range list {
				switch l.Name {
				case "description":
					if len(m.Description) == 0 {
						if pdebug.Enabled {
							pdebug.Printf("Replacing for key 'description' (fallback en -> %s", l.Language)
						}
						m.Description = l.Localized
					}
				case "title":
					if len(m.Title) == 0 {
						if pdebug.Enabled {
							pdebug.Printf("Replacing for key 'title' (fallback en -> %s", l.Language)
						}
						m.Title = l.Localized
					}
				}
			}
		}
		return nil
	case "all":
		for _, extralang := range []string{`ja`} {
			list = list[:0]
			if err := ls.LookupFields(tx, "ExternalResource", m.ID, extralang, &list); err != nil {
				return errors.Wrap(err, `failed to lookup localized fields`)
			}

			for _, l := range list {
				if pdebug.Enabled {
					pdebug.Printf("Adding key '%s#%s'", l.Name, l.Language)
				}
				m.LocalizedFields.Set(l.Language, l.Name, l.Localized)
			}
		}
	default:
		for _, extralang := range []string{`ja`} {
			list = list[:0]
			if err := ls.LookupFields(tx, "ExternalResource", m.ID, extralang, &list); err != nil {
				return errors.Wrap(err, `failed to lookup localized fields`)
			}

			for _, l := range list {
				switch l.Name {
				case "description":
					if pdebug.Enabled {
						pdebug.Printf("Replacing for key 'description'")
					}
					m.Description = l.Localized
				case "title":
					if pdebug.Enabled {
						pdebug.Printf("Replacing for key 'title'")
					}
					m.Title = l.Localized
				}
			}
		}
	}
	return nil
}

func (v *ExternalResourceSvc) Delete(tx *sql.Tx, id string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("ExternalResource.Delete (%s)", id)
		defer g.End()
	}
	original := db.ExternalResource{EID: id}
	if err := original.LoadByEID(tx, id); err != nil {
		return errors.Wrap(err, `failed load before delete`)
	}

	vdb := db.ExternalResource{EID: id}
	if err := vdb.Delete(tx); err != nil {
		return errors.Wrap(err, `failed to delete from database`)
	}
	c := Cache()
	key := c.Key("ExternalResource", id)
	c.Delete(key)
	if pdebug.Enabled {
		pdebug.Printf(`CACHE DEL %s`, key)
	}
	if err := db.DeleteLocalizedStringsForParent(tx, id, "ExternalResource"); err != nil {
		return errors.Wrap(err, `failed to delete localized strings`)
	}
	if err := v.PostDeleteHook(tx, &original); err != nil {
		return errors.Wrap(err, `post delete hook failed`)
	}
	return nil
}
