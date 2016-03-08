package octav

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/builderscon/octav/octav/db"
	"github.com/lestrrat/go-pdebug"
)

var ErrInvalidFieldType = errors.New("placeholder error")

func (v Venue) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["id"] = v.ID
	m["name"] = v.Name
	m["address"] = v.Address
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	if v.L10N.Len() == 0 {
		return buf, nil
	}

	l10buf, err := json.Marshal(v.L10N)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(buf[:len(buf)-1])
	b.WriteRune(',') // Replace closing '}'
	b.Write(l10buf[1:])

	return b.Bytes(), nil
}

func (v *Venue) UnmarshalJSON(data []byte) error {
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
			return ErrInvalidFieldType
		}
	}
	if jv, ok := m["name"]; ok {
		switch jv.(type) {
		case string:
			v.Name = jv.(string)
			delete(m, "name")
		default:
			return ErrInvalidFieldType
		}
	}
	if jv, ok := m["address"]; ok {
		switch jv.(type) {
		case string:
			v.Address = jv.(string)
			delete(m, "address")
		default:
			return ErrInvalidFieldType
		}
	}

	if err := ExtractL10NFields(m, &v.L10N, []string{"address", "name"}); err != nil {
		return err
	}

	return nil
}

func (v *Venue) Load(tx *db.Tx, id string) error {
	vdb := db.Venue{}
	if err := vdb.LoadByEID(tx, id); err != nil {
		return err
	}

	v.ID = vdb.EID
	v.Name = vdb.Name
	v.Address = vdb.Address

	ls, err := db.LoadLocalizedStringsForParent(tx, v.ID, "Venue")
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

func (v *Venue) Create(tx *db.Tx) error {
	if v.ID == "" {
		v.ID = UUID()
	}

	vdb := db.Venue{
		EID:     v.ID,
		Name:    v.Name,
		Address: v.Address,
	}
	if err := vdb.Create(tx); err != nil {
		return err
	}

	// vdb.EID, vdb.
	if v.L10N.Len() > 0 {
		err := v.L10N.Foreach(func(lang, key, val string) error {
			ldb := db.LocalizedString{
				ParentType: "Venue",
				ParentID:   vdb.EID,
				Language:   lang,
				Localized:  val,
				Name:       key,
			}
			return ldb.Create(tx)
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *Venue) Delete(tx *db.Tx) error {
	if pdebug.Enabled {
		g := pdebug.Marker("Venue.Delete (%s)", v.ID)
		defer g.End()
	}

	vdb := db.Venue{EID: v.ID}
	if err := vdb.Delete(tx); err != nil {
		return err
	}

	if err := db.DeleteLocalizedStringsForParent(tx, v.ID, "Venue"); err != nil {
		return err
	}
	return nil
}

func (v *VenueList) Load(tx *db.Tx, since string) error {
	var s int64
	if id := since; id != "" {
		vdb := db.Venue{}
		if err := vdb.LoadByEID(tx, id); err != nil {
			return err
		}

		s = vdb.OID
	}

	rows, err := tx.Query(`SELECT eid, name FROM venues WHERE oid > ? ORDER BY oid LIMIT 10`, s)
	if err != nil {
		return err
	}

	// Not using db.Venue here
	res := make([]Venue, 0, 10)
	for rows.Next() {
		v := Venue{}
		if err := rows.Scan(&v.ID, &v.Name); err != nil {
			return err
		}
		res = append(res, v)
	}
	*v = res
	return nil
}
