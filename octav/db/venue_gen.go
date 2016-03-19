package db

// Automatically generated by gendb utility. DO NOT EDIT!

import (
	"bytes"
	"database/sql"
	"errors"
	"strconv"
	"time"
)

const VenueStdSelectColumns = "venues.oid, venues.eid, venues.name, venues.address, venues.latitude, venues.longitude, venues.created_on, venues.modified_on"
const VenueTable = "venues"

type VenueList []Venue

func (v *Venue) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&v.OID, &v.EID, &v.Name, &v.Address, &v.Latitude, &v.Longitude, &v.CreatedOn, &v.ModifiedOn)
}

func (v *Venue) LoadByEID(tx *Tx, eid string) error {
	row := tx.QueryRow(`SELECT `+VenueStdSelectColumns+` FROM `+VenueTable+` WHERE venues.eid = ?`, eid)
	if err := v.Scan(row); err != nil {
		return err
	}
	return nil
}

func (v *Venue) Create(tx *Tx, opts ...InsertOption) error {
	if v.EID == "" {
		return errors.New("create: non-empty EID required")
	}

	v.CreatedOn = time.Now()
	doIgnore := false
	for _, opt := range opts {
		switch opt.(type) {
		case insertIgnoreOption:
			doIgnore = true
		}
	}

	stmt := bytes.Buffer{}
	stmt.WriteString("INSERT ")
	if doIgnore {
		stmt.WriteString("IGNORE ")
	}
	stmt.WriteString("INTO ")
	stmt.WriteString(VenueTable)
	stmt.WriteString(` (eid, name, address, latitude, longitude, created_on, modified_on) VALUES (?, ?, ?, ?, ?, ?, ?)`)
	result, err := tx.Exec(stmt.String(), v.EID, v.Name, v.Address, v.Latitude, v.Longitude, v.CreatedOn, v.ModifiedOn)
	if err != nil {
		return err
	}

	lii, err := result.LastInsertId()
	if err != nil {
		return err
	}

	v.OID = lii
	return nil
}

func (v Venue) Update(tx *Tx) error {
	if v.OID != 0 {
		_, err := tx.Exec(`UPDATE `+VenueTable+` SET eid = ?, name = ?, address = ?, latitude = ?, longitude = ? WHERE oid = ?`, v.EID, v.Name, v.Address, v.Latitude, v.Longitude, v.OID)
		return err
	}
	if v.EID != "" {
		_, err := tx.Exec(`UPDATE `+VenueTable+` SET name = ?, address = ?, latitude = ?, longitude = ? WHERE eid = ?`, v.Name, v.Address, v.Latitude, v.Longitude, v.EID)
		return err
	}
	return errors.New("either OID/EID must be filled")
}

func (v Venue) Delete(tx *Tx) error {
	if v.OID != 0 {
		_, err := tx.Exec(`DELETE FROM `+VenueTable+` WHERE oid = ?`, v.OID)
		return err
	}

	if v.EID != "" {
		_, err := tx.Exec(`DELETE FROM `+VenueTable+` WHERE eid = ?`, v.EID)
		return err
	}

	return errors.New("either OID/EID must be filled")
}

func (v *VenueList) FromRows(rows *sql.Rows, capacity int) error {
	var res []Venue
	if capacity > 0 {
		res = make([]Venue, 0, capacity)
	} else {
		res = []Venue{}
	}

	for rows.Next() {
		vdb := Venue{}
		if err := vdb.Scan(rows); err != nil {
			return err
		}
		res = append(res, vdb)
	}
	*v = res
	return nil
}

func (v *VenueList) LoadSinceEID(tx *Tx, since string, limit int) error {
	var s int64
	if id := since; id != "" {
		vdb := Venue{}
		if err := vdb.LoadByEID(tx, id); err != nil {
			return err
		}

		s = vdb.OID
	}
	return v.LoadSince(tx, s, limit)
}

func (v *VenueList) LoadSince(tx *Tx, since int64, limit int) error {
	rows, err := tx.Query(`SELECT `+VenueStdSelectColumns+` FROM `+VenueTable+` WHERE venues.oid > ? ORDER BY oid ASC LIMIT `+strconv.Itoa(limit), since)
	if err != nil {
		return err
	}

	if err := v.FromRows(rows, limit); err != nil {
		return err
	}
	return nil
}
