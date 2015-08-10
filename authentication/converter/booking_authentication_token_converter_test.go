package converter

import (
	"testing"

	. "github.com/bborbe/assert"
	booking_authentication "github.com/bborbe/booking/authentication"
)

func TestTokenConvertionEmpty(t *testing.T) {
	c := New()
	token, err := c.AuthenticationToToken(&booking_authentication.Authentication{})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(len(token), Gt(0)); err != nil {
		t.Fatal(err)
	}
	authentication, err := c.TokenToAuthentication(token)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(authentication, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(authentication.Token, Is("")); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(authentication.Login, Is("")); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(authentication.Password, Is("")); err != nil {
		t.Fatal(err)
	}
}

func TestTokenConvertionFilled(t *testing.T) {
	c := New()
	login := "testlogin"
	password := "testpass"
	token := "testtoken"
	s, err := c.AuthenticationToToken(&booking_authentication.Authentication{Login: login, Password: password, Token: token})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(len(s), Gt(0)); err != nil {
		t.Fatal(err)
	}
	authentication, err := c.TokenToAuthentication(s)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(authentication, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(authentication.Token, Is(token)); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(authentication.Login, Is(login)); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(authentication.Password, Is(password)); err != nil {
		t.Fatal(err)
	}
}
