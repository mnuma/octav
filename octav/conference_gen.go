// Automatically generated by genmodel utility. DO NOT EDIT!
package octav

import (
	"encoding/json"
	"github.com/builderscon/octav/octav/db"
	"github.com/lestrrat/go-pdebug"
)

func (v Conference) GetPropNames() ([]string, error) {
	l, _ := v.L10N.GetPropNames()
	return append(l, "id", "title", "sub_title", "slug", "dates"), nil
}

func (v Conference) GetPropValue(s string) (interface{}, error) {
	switch s {
	case "id":
		return v.ID, nil
	case "title":
		return v.Title, nil
	case "sub_title":
		return v.SubTitle, nil
	case "slug":
		return v.Slug, nil
	case "dates":
		return v.Dates, nil
	default:
		return v.L10N.GetPropValue(s)
	}
}

func (v Conference) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["id"] = v.ID
	m["title"] = v.Title
	m["sub_title"] = v.SubTitle
	m["slug"] = v.Slug
	m["dates"] = v.Dates
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return marshalJSONWithL10N(buf, v.L10N)
}

func (v *Conference) UnmarshalJSON(data []byte) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	if jv, ok := m["id"]; ok {
		switch jv.(type) {
		case string:
			v.ID = jv.(string)
			delete(m, "id")
		default:
			return ErrInvalidFieldType{Field: "id"}
		}
	}

	if jv, ok := m["title"]; ok {
		switch jv.(type) {
		case string:
			v.Title = jv.(string)
			delete(m, "title")
		default:
			return ErrInvalidFieldType{Field: "title"}
		}
	}

	if jv, ok := m["sub_title"]; ok {
		switch jv.(type) {
		case string:
			v.SubTitle = jv.(string)
			delete(m, "sub_title")
		default:
			return ErrInvalidFieldType{Field: "sub_title"}
		}
	}

	if jv, ok := m["slug"]; ok {
		switch jv.(type) {
		case string:
			v.Slug = jv.(string)
			delete(m, "slug")
		default:
			return ErrInvalidFieldType{Field: "slug"}
		}
	}

	if jv, ok := m["dates"]; ok {
		switch jv.(type) {
		case []ConferenceDate:
			v.Dates = jv.([]ConferenceDate)
			delete(m, "dates")
		default:
			return ErrInvalidFieldType{Field: "dates"}
		}
	}
	return nil
}

func (v *Conference) Load(tx *db.Tx, id string) error {
	vdb := db.Conference{}
	if err := vdb.LoadByEID(tx, id); err != nil {
		return err
	}

	if err := v.FromRow(vdb); err != nil {
		return err
	}
	if err := v.LoadLocalizedFields(tx); err != nil {
		return err
	}
	return nil
}

func (v *Conference) LoadLocalizedFields(tx *db.Tx) error {
	ls, err := db.LoadLocalizedStringsForParent(tx, v.ID, "Conference")
	if err != nil {
		return err
	}

	if len(ls) > 0 {
		v.L10N = LocalizedFields{}
		for _, l := range ls {
			v.L10N.Set(l.Language, l.Name, l.Localized)
		}
	}
	return nil
}

func (v *Conference) FromRow(vdb db.Conference) error {
	v.ID = vdb.EID
	v.Title = vdb.Title
	if vdb.SubTitle.Valid {
		v.SubTitle = vdb.SubTitle.String
	}
	v.Slug = vdb.Slug
	return nil
}

func (v *Conference) Create(tx *db.Tx) error {
	if v.ID == "" {
		v.ID = UUID()
	}

	vdb := db.Conference{}
	vdb.EID = v.ID
	vdb.Title = v.Title
	vdb.SubTitle.Valid = true
	vdb.SubTitle.String = v.SubTitle
	vdb.Slug = v.Slug
	if err := vdb.Create(tx); err != nil {
		return err
	}

	if err := v.L10N.CreateLocalizedStrings(tx, "Conference", v.ID); err != nil {
		return err
	}
	return nil
}

func (v *Conference) Delete(tx *db.Tx) error {
	if pdebug.Enabled {
		g := pdebug.Marker("Conference.Delete (%s)", v.ID)
		defer g.End()
	}

	vdb := db.Conference{EID: v.ID}
	if err := vdb.Delete(tx); err != nil {
		return err
	}
	if err := db.DeleteLocalizedStringsForParent(tx, v.ID, "Conference"); err != nil {
		return err
	}
	return nil
}

func (v *ConferenceList) Load(tx *db.Tx, since string, limit int) error {
	vdbl := db.ConferenceList{}
	if err := vdbl.LoadSinceEID(tx, since, limit); err != nil {
		return err
	}
	res := make([]Conference, len(vdbl))
	for i, vdb := range vdbl {
		v := Conference{}
		if err := v.FromRow(vdb); err != nil {
			return err
		}
		if err := v.LoadLocalizedFields(tx); err != nil {
			return err
		}
		res[i] = v
	}
	*v = res
	return nil
}
