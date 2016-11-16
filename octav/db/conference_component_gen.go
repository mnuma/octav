package db

// Automatically generated by gendb utility. DO NOT EDIT!

import (
	"bytes"
	"database/sql"
	"strconv"
	"time"

	"github.com/builderscon/octav/octav/tools"
	"github.com/lestrrat/go-pdebug"
	"github.com/pkg/errors"
)

const ConferenceComponentStdSelectColumns = "conference_components.oid, conference_components.eid, conference_components.conference_id, conference_components.name, conference_components.value, conference_components.created_on, conference_components.modified_on"
const ConferenceComponentTable = "conference_components"

type ConferenceComponentList []ConferenceComponent

func (c *ConferenceComponent) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&c.OID, &c.EID, &c.ConferenceID, &c.Name, &c.Value, &c.CreatedOn, &c.ModifiedOn)
}

func init() {
	hooks = append(hooks, func() {
		stmt := tools.GetBuffer()
		defer tools.ReleaseBuffer(stmt)

		stmt.Reset()
		stmt.WriteString(`DELETE FROM `)
		stmt.WriteString(ConferenceComponentTable)
		stmt.WriteString(` WHERE oid = ?`)
		library.Register("sqlConferenceComponentDeleteByOIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`UPDATE `)
		stmt.WriteString(ConferenceComponentTable)
		stmt.WriteString(` SET eid = ?, conference_id = ?, name = ?, value = ? WHERE oid = ?`)
		library.Register("sqlConferenceComponentUpdateByOIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`SELECT `)
		stmt.WriteString(ConferenceComponentStdSelectColumns)
		stmt.WriteString(` FROM `)
		stmt.WriteString(ConferenceComponentTable)
		stmt.WriteString(` WHERE `)
		stmt.WriteString(ConferenceComponentTable)
		stmt.WriteString(`.eid = ?`)
		library.Register("sqlConferenceComponentLoadByEIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`DELETE FROM `)
		stmt.WriteString(ConferenceComponentTable)
		stmt.WriteString(` WHERE eid = ?`)
		library.Register("sqlConferenceComponentDeleteByEIDKey", stmt.String())

		stmt.Reset()
		stmt.WriteString(`UPDATE `)
		stmt.WriteString(ConferenceComponentTable)
		stmt.WriteString(` SET eid = ?, conference_id = ?, name = ?, value = ? WHERE eid = ?`)
		library.Register("sqlConferenceComponentUpdateByEIDKey", stmt.String())
	})
}

func (c *ConferenceComponent) LoadByEID(tx *Tx, eid string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker(`ConferenceComponent.LoadByEID %s`, eid).BindError(&err)
		defer g.End()
	}
	stmt, err := library.GetStmt("sqlConferenceComponentLoadByEIDKey")
	if err != nil {
		return errors.Wrap(err, `failed to get statement`)
	}
	row := tx.Stmt(stmt).QueryRow(eid)
	if err := c.Scan(row); err != nil {
		return err
	}
	return nil
}

func (c *ConferenceComponent) Create(tx *Tx, opts ...InsertOption) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("db.ConferenceComponent.Create").BindError(&err)
		defer g.End()
		pdebug.Printf("%#v", c)
	}
	if c.EID == "" {
		return errors.New("create: non-empty EID required")
	}

	c.CreatedOn = time.Now()
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
	stmt.WriteString(ConferenceComponentTable)
	stmt.WriteString(` (eid, conference_id, name, value, created_on, modified_on) VALUES (?, ?, ?, ?, ?, ?)`)
	result, err := tx.Exec(stmt.String(), c.EID, c.ConferenceID, c.Name, c.Value, c.CreatedOn, c.ModifiedOn)
	if err != nil {
		return err
	}

	lii, err := result.LastInsertId()
	if err != nil {
		return err
	}

	c.OID = lii
	return nil
}

func (c ConferenceComponent) Update(tx *Tx) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker(`ConferenceComponent.Update`).BindError(&err)
		defer g.End()
	}
	if c.OID != 0 {
		if pdebug.Enabled {
			pdebug.Printf(`Using OID (%d) as key`, c.OID)
		}
		stmt, err := library.GetStmt("sqlConferenceComponentUpdateByOIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.EID, c.ConferenceID, c.Name, c.Value, c.OID)
		return err
	}
	if c.EID != "" {
		if pdebug.Enabled {
			pdebug.Printf(`Using EID (%s) as key`, c.EID)
		}
		stmt, err := library.GetStmt("sqlConferenceComponentUpdateByEIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.EID, c.ConferenceID, c.Name, c.Value, c.EID)
		return err
	}
	return errors.New("either OID/EID must be filled")
}

func (c ConferenceComponent) Delete(tx *Tx) error {
	if c.OID != 0 {
		stmt, err := library.GetStmt("sqlConferenceComponentDeleteByOIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.OID)
		return err
	}

	if c.EID != "" {
		stmt, err := library.GetStmt("sqlConferenceComponentDeleteByEIDKey")
		if err != nil {
			return errors.Wrap(err, `failed to get statement`)
		}
		_, err = tx.Stmt(stmt).Exec(c.EID)
		return err
	}

	return errors.New("either OID/EID must be filled")
}

func (v *ConferenceComponentList) FromRows(rows *sql.Rows, capacity int) error {
	var res []ConferenceComponent
	if capacity > 0 {
		res = make([]ConferenceComponent, 0, capacity)
	} else {
		res = []ConferenceComponent{}
	}

	for rows.Next() {
		vdb := ConferenceComponent{}
		if err := vdb.Scan(rows); err != nil {
			return err
		}
		res = append(res, vdb)
	}
	*v = res
	return nil
}

func (v *ConferenceComponentList) LoadSinceEID(tx *Tx, since string, limit int) error {
	var s int64
	if id := since; id != "" {
		vdb := ConferenceComponent{}
		if err := vdb.LoadByEID(tx, id); err != nil {
			return err
		}

		s = vdb.OID
	}
	return v.LoadSince(tx, s, limit)
}

func (v *ConferenceComponentList) LoadSince(tx *Tx, since int64, limit int) error {
	rows, err := tx.Query(`SELECT `+ConferenceComponentStdSelectColumns+` FROM `+ConferenceComponentTable+` WHERE conference_components.oid > ? ORDER BY oid ASC LIMIT `+strconv.Itoa(limit), since)
	if err != nil {
		return err
	}

	if err := v.FromRows(rows, limit); err != nil {
		return err
	}
	return nil
}
