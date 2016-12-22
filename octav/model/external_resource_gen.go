package model

// Automatically generated by genmodel utility. DO NOT EDIT!

import (
	"encoding/json"
	"time"

	"github.com/builderscon/octav/octav/db"
	"github.com/lestrrat/go-pdebug"
)

var _ = pdebug.Enabled
var _ = time.Time{}

type rawExternalResource struct {
	ID          string `json:"id"`
	Description string `json:"description"  l10n:"true"`
	Name        string `json:"name"  l10n:"true"`
	URL         string `json:"url"`
}

func (v ExternalResource) MarshalJSON() ([]byte, error) {
	var raw rawExternalResource
	raw.ID = v.ID
	raw.Description = v.Description
	raw.Name = v.Name
	raw.URL = v.URL
	buf, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}
	return MarshalJSONWithL10N(buf, v.LocalizedFields)
}

func (v *ExternalResource) Load(tx *db.Tx, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("model.ExternalResource.Load %s", id).BindError(&err)
		defer g.End()
	}
	vdb := db.ExternalResource{}
	if err := vdb.LoadByEID(tx, id); err != nil {
		return err
	}

	if err := v.FromRow(&vdb); err != nil {
		return err
	}
	return nil
}

func (v *ExternalResource) FromRow(vdb *db.ExternalResource) error {
	v.ID = vdb.EID
	v.Description = vdb.Description
	v.Name = vdb.Name
	v.URL = vdb.URL
	return nil
}

func (v *ExternalResource) ToRow(vdb *db.ExternalResource) error {
	vdb.EID = v.ID
	vdb.Description = v.Description
	vdb.Name = v.Name
	vdb.URL = v.URL
	return nil
}
