package service

// Automatically generated by genmodel utility. DO NOT EDIT!

import (
	"sync"
	"time"

	"github.com/builderscon/octav/octav/db"
	"github.com/builderscon/octav/octav/model"
	"github.com/lestrrat/go-pdebug"
	"github.com/pkg/errors"
)

var _ = time.Time{}

var questionSvc *QuestionSvc
var questionOnce sync.Once

func Question() *QuestionSvc {
	questionOnce.Do(questionSvc.Init)
	return questionSvc
}

func (v *QuestionSvc) LookupFromPayload(tx *db.Tx, m *model.Question, payload model.LookupQuestionRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Question.LookupFromPayload").BindError(&err)
		defer g.End()
	}
	if err = v.Lookup(tx, m, payload.ID); err != nil {
		return errors.Wrap(err, "failed to load model.Question from database")
	}
	return nil
}
func (v *QuestionSvc) Lookup(tx *db.Tx, m *model.Question, id string) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Question.Lookup").BindError(&err)
		defer g.End()
	}

	r := model.Question{}
	if err = r.Load(tx, id); err != nil {
		return errors.Wrap(err, "failed to load model.Question from database")
	}
	*m = r
	return nil
}

// Create takes in the transaction, the incoming payload, and a reference to
// a database row. The database row is initialized/populated so that the
// caller can use it afterwards.
func (v *QuestionSvc) Create(tx *db.Tx, vdb *db.Question, payload model.CreateQuestionRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Question.Create").BindError(&err)
		defer g.End()
	}

	if err := v.populateRowForCreate(vdb, payload); err != nil {
		return err
	}

	if err := vdb.Create(tx); err != nil {
		return err
	}

	return nil
}

func (v *QuestionSvc) Update(tx *db.Tx, vdb *db.Question, payload model.UpdateQuestionRequest) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("service.Question.Update (%s)", vdb.EID).BindError(&err)
		defer g.End()
	}

	if vdb.EID == "" {
		return errors.New("vdb.EID is required (did you forget to call vdb.Load(tx) before hand?)")
	}

	if err := v.populateRowForUpdate(vdb, payload); err != nil {
		return err
	}

	if err := vdb.Update(tx); err != nil {
		return err
	}
	return nil
}

func (v *QuestionSvc) Delete(tx *db.Tx, id string) error {
	if pdebug.Enabled {
		g := pdebug.Marker("Question.Delete (%s)", id)
		defer g.End()
	}

	vdb := db.Question{EID: id}
	if err := vdb.Delete(tx); err != nil {
		return err
	}
	return nil
}

func (v *QuestionSvc) LoadList(tx *db.Tx, vdbl *db.QuestionList, since string, limit int) error {
	return vdbl.LoadSinceEID(tx, since, limit)
}
