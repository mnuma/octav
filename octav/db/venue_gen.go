// Automatically generated by gendb utility. DO NOT EDIT!
package db

import (
	"errors"
)

var VenueTable = "venues"

func (v *Venue) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&v.OID, &v.EID, &v.Name, &v.Address, &v.CreatedOn, &v.ModifiedOn)
}

func (v *Venue) LoadByEID(tx *Tx, eid string) error {
	row := tx.QueryRow(`SELECT oid, eid, name, address, created_on, modified_on FROM `+VenueTable+` WHERE eid = ?`, eid)
	if err := v.Scan(row); err != nil {
		return err
	}
	return nil
}

func (v *Venue) Create(tx *Tx) error {
	if v.EID == "" {
		return errors.New("create: non-empty EID required")
	}

	result, err := tx.Exec(`INSERT INTO `+VenueTable+` (eid, name, address, created_on, modified_on) VALUES (?, ?, ?, ?, ?)`, v.EID, v.Name, v.Address, v.CreatedOn, v.ModifiedOn)
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

func (v Venue) Delete(tx *Tx) error {
	if v.OID != 0 {
		_, err := tx.Exec(`DELETE FROM `+VenueTable+` WHERE oid = ?`, v.OID)
		return err
	}

	if v.EID != "" {
		_, err := tx.Exec(`DELETE FROM `+VenueTable+` WHERE eid = ?`, v.EID)
		return err
	}

	return errors.New("either OID/EID musti be filled")
}
