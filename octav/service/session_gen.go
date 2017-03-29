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
var _ = model.Session{}
var _ = db.Session{}
var _ = sql.ErrNoRows
var _ = pdebug.Enabled

var sessionSvc SessionSvc
var sessionOnce sync.Once

func Session() *SessionSvc {
	sessionOnce.Do(sessionSvc.Init)
	return &sessionSvc
}

func (v *SessionSvc) LookupFromPayload(ctx context.Context, tx *sql.Tx, m *model.Session, payload *model.LookupSessionRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Session.LookupFromPayload %s", payload.ID).BindError(&err)
		defer g.End()
	}
	if err = v.Lookup(ctx, tx, m, payload.ID); err != nil {
		return errors.Wrap(err, "failed to load model.Session from database")
	}
	if err := v.Decorate(ctx, tx, m, payload.TrustedCall, payload.Lang.String); err != nil {
		return errors.Wrap(err, "failed to load associated data for model.Session from database")
	}
	return nil
}

func (v *SessionSvc) Lookup(ctx context.Context, tx *sql.Tx, m *model.Session, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Session.Lookup %s", id).BindError(&err)
		defer g.End()
	}

	var r model.Session
	c := Cache()
	key := c.Key("Session", id)
	var cacheMiss bool
	_, err = c.GetOrSet(key, &r, func() (interface{}, error) {
		if pdebug.Enabled {
			cacheMiss = true
		}
		if err := r.Load(tx, id); err != nil {
			return nil, errors.Wrap(err, "failed to load model.Session from database")
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
func (v *SessionSvc) Create(ctx context.Context, tx *sql.Tx, vdb *db.Session, payload *model.CreateSessionRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Session.Create").BindError(&err)
		defer g.End()
	}

	if err := v.populateRowForCreate(ctx, vdb, payload); err != nil {
		return errors.Wrap(err, `failed to populate row`)
	}

	if err := vdb.Create(tx, payload.DatabaseOptions...); err != nil {
		return errors.Wrap(err, `failed to insert into database`)
	}

	if err := payload.LocalizedFields.CreateLocalizedStrings(tx, "Session", vdb.EID); err != nil {
		return errors.Wrap(err, `failed to populate localized strings`)
	}
	return nil
}

func (v *SessionSvc) Update(tx *sql.Tx, vdb *db.Session) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Session.Update (%s)", vdb.EID).BindError(&err)
		defer g.End()
	}

	if vdb.EID == `` {
		return errors.New("vdb.EID is required (did you forget to call vdb.Load(tx) before hand?)")
	}

	if err := vdb.Update(tx); err != nil {
		return errors.Wrap(err, `failed to update database`)
	}
	c := Cache()
	key := c.Key("Session", vdb.EID)
	if pdebug.Enabled {
		pdebug.Printf(`CACHE DEL %s`, key)
	}
	cerr := c.Delete(key)
	if pdebug.Enabled {
		if cerr != nil {
			pdebug.Printf(`CACHE ERR: %s`, cerr)
		}
	}
	return nil
}

func (v *SessionSvc) UpdateFromPayload(ctx context.Context, tx *sql.Tx, payload *model.UpdateSessionRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Session.UpdateFromPayload (%s)", payload.ID).BindError(&err)
		defer g.End()
	}
	var vdb db.Session
	if err := vdb.LoadByEID(tx, payload.ID); err != nil {
		return errors.Wrap(err, `failed to load from database`)
	}

	if err := v.PreUpdateFromPayloadHook(ctx, tx, &vdb, payload); err != nil {
		return errors.Wrap(err, `failed to execute PreUpdateFromPayloadHook`)
	}

	if err := v.populateRowForUpdate(ctx, &vdb, payload); err != nil {
		return errors.Wrap(err, `failed to populate row data`)
	}

	if err := v.Update(tx, &vdb); err != nil {
		return errors.Wrap(err, `failed to update row in database`)
	}

	ls := LocalizedString()
	if err := ls.UpdateFields(tx, "Session", vdb.EID, payload.LocalizedFields); err != nil {
		return errors.Wrap(err, `failed to update localized fields`)
	}
	return nil
}

func (v *SessionSvc) ReplaceL10NStrings(tx *sql.Tx, m *model.Session, lang string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Session.ReplaceL10NStrings lang = %s", lang)
		defer g.End()
	}
	ls := LocalizedString()
	list := make([]db.LocalizedString, 0, 2)
	switch lang {
	case "", "en":
		if len(m.Title) > 0 && len(m.Abstract) > 0 {
			return nil
		}
		for _, extralang := range []string{`ja`} {
			list = list[:0]
			if err := ls.LookupFields(tx, "Session", m.ID, extralang, &list); err != nil {
				return errors.Wrap(err, `failed to lookup localized fields`)
			}

			for _, l := range list {
				switch l.Name {
				case "title":
					if len(m.Title) == 0 {
						if pdebug.Enabled {
							pdebug.Printf("Replacing for key 'title' (fallback en -> %s", l.Language)
						}
						m.Title = l.Localized
					}
				case "abstract":
					if len(m.Abstract) == 0 {
						if pdebug.Enabled {
							pdebug.Printf("Replacing for key 'abstract' (fallback en -> %s", l.Language)
						}
						m.Abstract = l.Localized
					}
				}
			}
		}
		return nil
	case "all":
		for _, extralang := range []string{`ja`} {
			list = list[:0]
			if err := ls.LookupFields(tx, "Session", m.ID, extralang, &list); err != nil {
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
			if err := ls.LookupFields(tx, "Session", m.ID, extralang, &list); err != nil {
				return errors.Wrap(err, `failed to lookup localized fields`)
			}

			for _, l := range list {
				switch l.Name {
				case "title":
					if pdebug.Enabled {
						pdebug.Printf("Replacing for key 'title'")
					}
					m.Title = l.Localized
				case "abstract":
					if pdebug.Enabled {
						pdebug.Printf("Replacing for key 'abstract'")
					}
					m.Abstract = l.Localized
				}
			}
		}
	}
	return nil
}

func (v *SessionSvc) Delete(tx *sql.Tx, id string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("Session.Delete (%s)", id)
		defer g.End()
	}

	vdb := db.Session{EID: id}
	if err := vdb.Delete(tx); err != nil {
		return errors.Wrap(err, `failed to delete from database`)
	}
	c := Cache()
	key := c.Key("Session", id)
	c.Delete(key)
	if pdebug.Enabled {
		pdebug.Printf(`CACHE DEL %s`, key)
	}
	if err := db.DeleteLocalizedStringsForParent(tx, id, "Session"); err != nil {
		return errors.Wrap(err, `failed to delete localized strings`)
	}
	return nil
}
