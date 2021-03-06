package permission_check_handler

import (
	"testing"

	"fmt"
	"net/http"

	. "github.com/bborbe/assert"

	booking_authentication "github.com/bborbe/booking/authentication"
	booking_authorization "github.com/bborbe/booking/authorization"
	booking_handler "github.com/bborbe/booking/handler"
	booking_handler_mock "github.com/bborbe/booking/handler/mock"
	server_mock "github.com/bborbe/server/mock"
)

func createHttpRequestToAuthentication(authentication *booking_authentication.Authentication, err error) HttpRequestToAuthentication {
	return func(request *http.Request) (*booking_authentication.Authentication, error) {
		return authentication, err
	}

}
func createHasRole(valid bool, err error) HasRole {
	return func(authentication *booking_authentication.Authentication, role booking_authorization.Role) (bool, error) {
		return valid, err
	}
}

func createVerifyLogin(valid bool, err error) func(authentication *booking_authentication.Authentication) (bool, error) {
	return func(authentication *booking_authentication.Authentication) (bool, error) {
		return valid, err
	}
}

func TestImplementsHandler(t *testing.T) {
	r := New(nil, nil, nil, booking_authorization.Administrator, nil)
	var i *booking_handler.Handler
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestAllwaysAllowNone(t *testing.T) {
	e := fmt.Errorf("myError")
	handler := booking_handler_mock.New()
	r := New(createVerifyLogin(false, nil), createHasRole(false, e), createHttpRequestToAuthentication(nil, e), booking_authorization.None, handler)
	if err := AssertThat(r.checkPermission(&http.Request{}), NilValue()); err != nil {
		t.Fatal(err)
	}
}

func TestHttpRequestToAuthenticationFailed(t *testing.T) {
	e := fmt.Errorf("myError")
	handler := booking_handler_mock.New()
	r := New(createVerifyLogin(false, nil), createHasRole(true, nil), createHttpRequestToAuthentication(nil, e), booking_authorization.Administrator, handler)
	if err := AssertThat(r.ServeHTTP(server_mock.NewHttpResponseWriterMock(), &http.Request{}), Is(e)); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(handler.Request == nil, Is(true)); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(handler.Response == nil, Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestHasRoleFailed(t *testing.T) {
	e := fmt.Errorf("myError")
	handler := booking_handler_mock.New()
	r := New(createVerifyLogin(true, nil), createHasRole(true, e), createHttpRequestToAuthentication(nil, nil), booking_authorization.Administrator, handler)
	if err := AssertThat(r.ServeHTTP(server_mock.NewHttpResponseWriterMock(), &http.Request{}), Is(e)); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(handler.Request == nil, Is(true)); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(handler.Response == nil, Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestHasRoleReturnFalse(t *testing.T) {
	handler := booking_handler_mock.New()
	r := New(createVerifyLogin(false, nil), createHasRole(false, nil), createHttpRequestToAuthentication(nil, nil), booking_authorization.Administrator, handler)
	if err := AssertThat(r.ServeHTTP(server_mock.NewHttpResponseWriterMock(), &http.Request{}), NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(handler.Request == nil, Is(true)); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(handler.Response == nil, Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestLoginError(t *testing.T) {
	e := fmt.Errorf("myError")
	handler := booking_handler_mock.New()
	handler.Error = e
	r := New(createVerifyLogin(false, e), createHasRole(true, nil), createHttpRequestToAuthentication(nil, nil), booking_authorization.Administrator, handler)
	if err := AssertThat(r.ServeHTTP(server_mock.NewHttpResponseWriterMock(), &http.Request{}), Is(e)); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(handler.Request == nil, Is(true)); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(handler.Response == nil, Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestLoginFailed(t *testing.T) {
	e := fmt.Errorf("myError")
	handler := booking_handler_mock.New()
	handler.Error = e
	r := New(createVerifyLogin(false, nil), createHasRole(true, nil), createHttpRequestToAuthentication(nil, nil), booking_authorization.Administrator, handler)
	if err := AssertThat(r.ServeHTTP(server_mock.NewHttpResponseWriterMock(), &http.Request{}), NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(handler.Request == nil, Is(true)); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(handler.Response == nil, Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestSubHandlerCalled(t *testing.T) {
	e := fmt.Errorf("myError")
	handler := booking_handler_mock.New()
	handler.Error = e
	r := New(createVerifyLogin(true, nil), createHasRole(true, nil), createHttpRequestToAuthentication(nil, nil), booking_authorization.Administrator, handler)
	if err := AssertThat(r.ServeHTTP(server_mock.NewHttpResponseWriterMock(), &http.Request{}), Is(e)); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(handler.Request, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(handler.Response, NotNilValue()); err != nil {
		t.Fatal(err)
	}
}
