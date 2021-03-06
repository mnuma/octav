package service

import (
	"database/sql"
	"time"

	"github.com/builderscon/octav/octav/cache"
	"github.com/builderscon/octav/octav/db"
	"github.com/builderscon/octav/octav/internal/context"
	"github.com/builderscon/octav/octav/model"
	"github.com/builderscon/octav/octav/tools"
	pdebug "github.com/lestrrat/go-pdebug"
	"github.com/pkg/errors"
)

func (v *RoomSvc) Init() {}

func (v *RoomSvc) populateRowForCreate(ctx context.Context, vdb *db.Room, payload *model.CreateRoomRequest) error {
	vdb.EID = tools.UUID()

	if payload.VenueID.Valid() {
		vdb.VenueID = payload.VenueID.String
	}

	if payload.Name.Valid() {
		vdb.Name = payload.Name.String
	}

	if payload.Capacity.Valid() {
		vdb.Capacity = uint(payload.Capacity.Uint)
	}

	return nil
}

func (v *RoomSvc) populateRowForUpdate(ctx context.Context, vdb *db.Room, payload *model.UpdateRoomRequest) error {
	if payload.VenueID.Valid() {
		vdb.VenueID = payload.VenueID.String
	}

	if payload.Name.Valid() {
		vdb.Name = payload.Name.String
	}

	if payload.Capacity.Valid() {
		vdb.Capacity = uint(payload.Capacity.Uint)
	}

	return nil
}

func (v *RoomSvc) CreateFromPayload(ctx context.Context, tx *sql.Tx, result *model.Room, payload *model.CreateRoomRequest) error {
	su := User()
	if err := su.IsAdministrator(ctx, tx, context.GetUserID(ctx)); err != nil {
		return errors.Wrap(err, "creating a room requires conference administrator privilege")
	}

	vdb := db.Room{}
	if err := v.Create(ctx, tx, &vdb, payload); err != nil {
		return errors.Wrap(err, "failed to store in database")
	}

	var r model.Room
	if err := r.FromRow(&vdb); err != nil {
		return errors.Wrap(err, "failed to populate model from database")
	}

	*result = r
	return nil
}

func (v *RoomSvc) ListFromPayload(ctx context.Context, tx *sql.Tx, result *model.RoomList, payload *model.ListRoomRequest) error {
	var m model.RoomList
	if err := m.LoadForVenue(tx, payload.VenueID, payload.Since.String, int(payload.Limit.Int)); err != nil {
		return errors.Wrap(err, "failed to load from database")
	}

	for i := range m {
		if err := v.Decorate(ctx, tx, &m[i], payload.VerifiedCall, payload.Lang.String); err != nil {
			return errors.Wrap(err, "failed to associate data to model")
		}
	}

	*result = m
	return nil
}

func (v *RoomSvc) PreUpdateFromPayloadHook(ctx context.Context, tx *sql.Tx, vdb *db.Room, payload *model.UpdateRoomRequest) error {
	su := User()
	if err := su.IsAdministrator(ctx, tx, context.GetUserID(ctx)); err != nil {
		return errors.Wrap(err, "deleting rooms require administrator privileges")
	}
	return nil
}

func (v *RoomSvc) PostCreateHook(ctx context.Context, tx *sql.Tx, vdb *db.Room) error {
	return invalidateRoomLoadByVenueID(vdb.VenueID)
}

func (v *RoomSvc) PostUpdateHook(tx *sql.Tx, vdb *db.Room) error {
	return invalidateRoomLoadByVenueID(vdb.VenueID)
}

func invalidateRoomLoadByVenueID(venueID string) error {
	c := Cache()
	key := c.Key("Room", "LoadByVenueID", venueID)
	c.Delete(key)
	if pdebug.Enabled {
		pdebug.Printf("CACHE DEL: %s", key)
	}
	return nil
}

func (v *RoomSvc) DeleteFromPayload(ctx context.Context, tx *sql.Tx, payload *model.DeleteRoomRequest) error {
	su := User()
	if err := su.IsAdministrator(ctx, tx, context.GetUserID(ctx)); err != nil {
		return errors.Wrap(err, "deleting rooms require administrator privileges")
	}

	return errors.Wrap(v.Delete(tx, payload.ID), "failed to delete from ddatabase")
}

func (v *RoomSvc) Decorate(ctx context.Context, tx *sql.Tx, room *model.Room, verifiedCall bool, lang string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Room.Decorate (%s, %t, %s)", room.ID, verifiedCall, lang).BindError(&err)
		defer g.End()
	}

	if lang != "" {
		if err := v.ReplaceL10NStrings(tx, room, lang); err != nil {
			return errors.Wrap(err, "failed to replace L10N strings")
		}
	}
	return nil
}

func (v *RoomSvc) LoadByVenueID(ctx context.Context, tx *sql.Tx, cdl *model.RoomList, venueID string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Room.LoadByVenueID %s", venueID).BindError(&err)
		defer g.End()
	}

	var ids []string
	c := Cache()
	key := c.Key("Room", "LoadByVenueID", venueID)
	if err := c.Get(key, &ids); err == nil {
		if pdebug.Enabled {
			pdebug.Printf("CACHE HIT: %s", key)
		}
		m := make(model.RoomList, len(ids))
		for i, id := range ids {
			if err := v.Lookup(ctx, tx, &m[i], id); err != nil {
				return errors.Wrap(err, "failed to load from database")
			}
		}

		*cdl = m
		return nil
	}

	if pdebug.Enabled {
		pdebug.Printf("CACHE MISS: %s", key)
	}
	var vdbl db.RoomList
	if err := db.LoadVenueRooms(tx, &vdbl, venueID); err != nil {
		return err
	}

	ids = make([]string, len(vdbl))
	res := make(model.RoomList, len(vdbl))
	for i, vdb := range vdbl {
		var u model.Room
		if err := u.FromRow(&vdb); err != nil {
			return err
		}
		ids[i] = vdb.EID
		res[i] = u
	}
	*cdl = res

	c.Set(key, ids, cache.WithExpires(15*time.Minute))
	return nil
}
