package service

import (
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/booking/authentication"
	"github.com/bborbe/booking/database/sqlite"
	"github.com/bborbe/booking/user"

	user_service "github.com/bborbe/booking/user/service"
	user_storage "github.com/bborbe/booking/user/storage"

	model_service "github.com/bborbe/booking/model/service"
	model_storage "github.com/bborbe/booking/model/storage"

	"github.com/bborbe/booking/tokengenerator"
)

func TestImplementsAuthenticationService(t *testing.T) {
	r := New(nil, nil)
	var i *Service
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyLoginUser(t *testing.T) {
	var err error
	var valid bool
	name := "testuser"
	password := "testpassword"
	database := sqlite.New("/tmp/booking_test.db", false)

	userStorage := user_storage.New(database)
	userService := user_service.New(userStorage)

	modelStorage := model_storage.New(database)
	modelService := model_service.New(modelStorage, tokengenerator.New())

	_, err = userService.Create(&user.User{Login: name, Password: password})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	authenticationService := New(userService, modelService)

	valid, err = authenticationService.VerifyLogin(&authentication.Authentication{Login: name, Password: password})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(valid, Is(true)); err != nil {
		t.Fatal(err)
	}

	valid, err = authenticationService.VerifyLogin(&authentication.Authentication{Login: name, Password: "wrong"})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(valid, Is(false)); err != nil {
		t.Fatal(err)
	}

	valid, err = authenticationService.VerifyLogin(&authentication.Authentication{Login: "wrong", Password: password})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(valid, Is(false)); err != nil {
		t.Fatal(err)
	}
}
