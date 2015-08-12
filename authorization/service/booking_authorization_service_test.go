package service

import (
	"testing"

	. "github.com/bborbe/assert"
	booking_authentication "github.com/bborbe/booking/authentication"
	booking_authorization "github.com/bborbe/booking/authorization"
)

func TestImplementsAuthenticationService(t *testing.T) {
	r := New(nil)
	var i *Service
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func createVerifyLogin(valid bool, err error) VerifyLogin {
	return func(*booking_authentication.Authentication) (bool, error) {
		return valid, err
	}
}

func TestHasRoleNone(t *testing.T) {
	r := New(createVerifyLogin(false, nil))
	hasRole, err := r.HasRole(&booking_authentication.Authentication{}, booking_authorization.None)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(hasRole, Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestHasRoleParticipantVerifyLoginFailed(t *testing.T) {
	r := New(createVerifyLogin(false, nil))
	authentication := &booking_authentication.Authentication{Token: "token"}
	hasRole, err := r.HasRole(authentication, booking_authorization.Participant)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(hasRole, Is(false)); err != nil {
		t.Fatal(err)
	}
}

func TestHasRoleParticipantVerifyLoginSuccess(t *testing.T) {
	r := New(createVerifyLogin(true, nil))
	authentication := &booking_authentication.Authentication{Token: "token"}
	hasRole, err := r.HasRole(authentication, booking_authorization.Participant)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(hasRole, Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestHasRoleParticipantIllegalParameter(t *testing.T) {
	r := New(createVerifyLogin(true, nil))
	authentication := &booking_authentication.Authentication{Token: "token", Login: "bad"}
	hasRole, err := r.HasRole(authentication, booking_authorization.Participant)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(hasRole, Is(false)); err != nil {
		t.Fatal(err)
	}
}

func TestHasRoleAdminVerifyLoginFailed(t *testing.T) {
	r := New(createVerifyLogin(false, nil))
	authentication := &booking_authentication.Authentication{Login: "admin"}
	hasRole, err := r.HasRole(authentication, booking_authorization.Administrator)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(hasRole, Is(false)); err != nil {
		t.Fatal(err)
	}
}

func TestHasRoleAdminVerifyLoginSuccess(t *testing.T) {
	r := New(createVerifyLogin(true, nil))
	authentication := &booking_authentication.Authentication{Login: "admin"}
	hasRole, err := r.HasRole(authentication, booking_authorization.Administrator)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(hasRole, Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestHasRoleAdminVerifyLoginSuccessNoAdmin(t *testing.T) {
	r := New(createVerifyLogin(true, nil))
	authentication := &booking_authentication.Authentication{Login: "user"}
	hasRole, err := r.HasRole(authentication, booking_authorization.Administrator)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(hasRole, Is(false)); err != nil {
		t.Fatal(err)
	}
}

func TestHasRoleAdminIllegalParameter(t *testing.T) {
	r := New(createVerifyLogin(true, nil))
	authentication := &booking_authentication.Authentication{Token: "token", Login: "admin"}
	hasRole, err := r.HasRole(authentication, booking_authorization.Administrator)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(hasRole, Is(false)); err != nil {
		t.Fatal(err)
	}
}

func TestHasRoleOrganizerVerifyLoginFailed(t *testing.T) {
	r := New(createVerifyLogin(false, nil))
	authentication := &booking_authentication.Authentication{Login: "orga"}
	hasRole, err := r.HasRole(authentication, booking_authorization.Organizer)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(hasRole, Is(false)); err != nil {
		t.Fatal(err)
	}
}

func TestHasRoleOrganizerVerifyLoginSuccess(t *testing.T) {
	r := New(createVerifyLogin(true, nil))
	authentication := &booking_authentication.Authentication{Login: "orga"}
	hasRole, err := r.HasRole(authentication, booking_authorization.Organizer)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(hasRole, Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestHasRoleOrganizerIllegalParameter(t *testing.T) {
	r := New(createVerifyLogin(true, nil))
	authentication := &booking_authentication.Authentication{Token: "token", Login: "orga"}
	hasRole, err := r.HasRole(authentication, booking_authorization.Organizer)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(hasRole, Is(false)); err != nil {
		t.Fatal(err)
	}
}

func TestAdminHasOrganizerRole(t *testing.T) {
	r := New(createVerifyLogin(true, nil))
	authentication := &booking_authentication.Authentication{Login: "admin"}
	hasRole, err := r.HasRole(authentication, booking_authorization.Organizer)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(hasRole, Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestAdminHasParticipantRole(t *testing.T) {
	r := New(createVerifyLogin(true, nil))
	authentication := &booking_authentication.Authentication{Login: "admin"}
	hasRole, err := r.HasRole(authentication, booking_authorization.Participant)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(hasRole, Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestOrganizerHasOrganizerRole(t *testing.T) {
	r := New(createVerifyLogin(true, nil))
	authentication := &booking_authentication.Authentication{Login: "orga"}
	hasRole, err := r.HasRole(authentication, booking_authorization.Participant)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(hasRole, Is(true)); err != nil {
		t.Fatal(err)
	}
}
