package service

import (
	"testing"

	. "github.com/bborbe/assert"

	booking_database_sqlite "github.com/bborbe/booking/database/sqlite"
	booking_model "github.com/bborbe/booking/model"
	booking_model_storage "github.com/bborbe/booking/model/storage"
	booking_tokengenerator "github.com/bborbe/booking/tokengenerator"
)

func TestImplementsModelService(t *testing.T) {
	r := New(nil, nil)
	var i *Service
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyLogin(t *testing.T) {
	var valid bool
	var m *booking_model.Model
	var err error
	database := booking_database_sqlite.New("/tmp/booking_test.db", true)
	modelStorage := booking_model_storage.New(database)
	modelService := New(modelStorage, booking_tokengenerator.New())

	m, err = modelService.Create(&booking_model.Model{})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	token := m.Token

	valid, err = modelService.VerifyLogin(&booking_model.Model{Token: token})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(valid, Is(true)); err != nil {
		t.Fatal(err)
	}

	valid, err = modelService.VerifyLogin(&booking_model.Model{Token: "wrong"})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(valid, Is(false)); err != nil {
		t.Fatal(err)
	}
}
