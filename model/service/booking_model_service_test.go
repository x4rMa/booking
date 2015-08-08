package service

import (
	"testing"

	. "github.com/bborbe/assert"

	"github.com/bborbe/booking/database/sqlite"
	"github.com/bborbe/booking/model"
	"github.com/bborbe/booking/model/storage"
	"github.com/bborbe/booking/tokengenerator"
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
	var m *model.Model
	var err error
	database := sqlite.New("/tmp/booking_test.db", true)
	modelStorage := storage.New(database)
	modelService := New(modelStorage, tokengenerator.New())

	m, err = modelService.Create(&model.Model{})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	token := m.Token

	valid, err = modelService.VerifyLogin(&model.Model{Token: token})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(valid, Is(true)); err != nil {
		t.Fatal(err)
	}

	valid, err = modelService.VerifyLogin(&model.Model{Token: "wrong"})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(valid, Is(false)); err != nil {
		t.Fatal(err)
	}
}
