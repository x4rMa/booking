package handler_configuration

import (
	"net/http"
	"testing"

	"github.com/bborbe/eventbus"

	. "github.com/bborbe/assert"
	booking_database "github.com/bborbe/booking/database"
	booking_database_sqlite "github.com/bborbe/booking/database/sqlite"

	io_mock "github.com/bborbe/io/mock"
	server_mock "github.com/bborbe/server/mock"

	booking_date_service "github.com/bborbe/booking/date/service"
	booking_date_storage "github.com/bborbe/booking/date/storage"

	booking_model_service "github.com/bborbe/booking/model/service"
	booking_model_storage "github.com/bborbe/booking/model/storage"

	booking_shooting_service "github.com/bborbe/booking/shooting/service"
	booking_shooting_storage "github.com/bborbe/booking/shooting/storage"

	booking_user "github.com/bborbe/booking/user"
	booking_user_service "github.com/bborbe/booking/user/service"
	booking_user_storage "github.com/bborbe/booking/user/storage"

	booking_authentication "github.com/bborbe/booking/authentication"
	booking_authentication_service "github.com/bborbe/booking/authentication/service"

	booking_authentication_converter "github.com/bborbe/booking/authentication/converter"

	"encoding/base64"
	"encoding/json"

	booking_authorization "github.com/bborbe/booking/authorization"
	booking_authorization_service "github.com/bborbe/booking/authorization/service"
	booking_permission_check_handler "github.com/bborbe/booking/permission_check_handler"
	booking_tokengenerator "github.com/bborbe/booking/tokengenerator"
)

func createHasRole(valid bool, err error) booking_permission_check_handler.HasRole {
	return func(authentication *booking_authentication.Authentication, role booking_authorization.Role) (bool, error) {
		return valid, err
	}
}

func createHttpRequestToAuthentication(authentication *booking_authentication.Authentication, err error) func(request *http.Request) (*booking_authentication.Authentication, error) {
	return func(request *http.Request) (*booking_authentication.Authentication, error) {
		return authentication, err
	}
}

func createRequest(path string, userService booking_user_service.Service) (*http.Request, error) {
	login := "admin"
	pass := "test123"
	token, err := createToken(&booking_authentication.Authentication{Login: login, Password: pass})
	if err != nil {
		return nil, err
	}
	_, err = userService.Create(&booking_user.User{Login: login, Password: pass})
	if err != nil {
		return nil, err
	}
	req, err := server_mock.NewHttpRequestMock(path)
	if err != nil {
		return nil, err
	}
	req.Header["X-Auth-Token"] = []string{token}
	return req, nil
}

func createToken(authentication *booking_authentication.Authentication) (string, error) {
	b, err := json.Marshal(authentication)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func createHandler(db booking_database.Database, userService booking_user_service.Service) http.Handler {
	tokengenerator := booking_tokengenerator.New()
	modelService := booking_model_service.New(booking_model_storage.New(db), tokengenerator)
	dateService := booking_date_service.New(booking_date_storage.New(db))
	authenticationService := booking_authentication_service.New(userService, modelService)
	shootingService := booking_shooting_service.New(booking_shooting_storage.New(db), eventbus.New())
	authorizationService := booking_authorization_service.New(authenticationService.VerifyLogin)
	authenticationConverter := booking_authentication_converter.New()
	handlerConfiguration := New("/tmp", dateService, modelService, shootingService, userService, authenticationService, authorizationService, authenticationConverter)
	return handlerConfiguration.GetHandler()
}

func TestNewHandlerImplementsHttpHandler(t *testing.T) {
	db := booking_database_sqlite.New("/tmp/booking_test.db", false)
	userService := booking_user_service.New(booking_user_storage.New(db))
	handler := createHandler(db, userService)
	var i (*http.Handler) = nil
	err := AssertThat(handler, Implements(i).Message("check implements server.Server"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestDate(t *testing.T) {
	resp := server_mock.NewHttpResponseWriterMock()
	db := booking_database_sqlite.New("/tmp/booking_test.db", false)
	userService := booking_user_service.New(booking_user_storage.New(db))
	handler := createHandler(db, userService)
	req, err := createRequest("/date", userService)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(resp, req)
	if err = AssertThat(resp.Status(), Is(200)); err != nil {
		t.Fatal(err)
	}
}

func TestGetDate(t *testing.T) {
	resp := server_mock.NewHttpResponseWriterMock()
	db := booking_database_sqlite.New("/tmp/booking_test.db", false)
	userService := booking_user_service.New(booking_user_storage.New(db))
	handler := createHandler(db, userService)
	req, err := createRequest("/date", userService)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	req.Method = "GET"
	handler.ServeHTTP(resp, req)
	if err = AssertThat(resp.Status(), Is(200)); err != nil {
		t.Fatal(err)
	}
}

func TestPutDate(t *testing.T) {
	resp := server_mock.NewHttpResponseWriterMock()
	db := booking_database_sqlite.New("/tmp/booking_test.db", false)
	userService := booking_user_service.New(booking_user_storage.New(db))
	handler := createHandler(db, userService)
	req, err := createRequest("/date", userService)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	req.Body = io_mock.NewReadCloserString("{}")
	req.Method = "PUT"
	handler.ServeHTTP(resp, req)
	if err = AssertThat(resp.Status(), Is(200)); err != nil {
		t.Fatal(err)
	}
}

func TestVerifyLoginHandlerFound(t *testing.T) {
	resp := server_mock.NewHttpResponseWriterMock()
	db := booking_database_sqlite.New("/tmp/booking_test.db", false)
	userService := booking_user_service.New(booking_user_storage.New(db))
	_, err := userService.Create(&booking_user.User{Login: "testuser", Password: "testpassword"})
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	handler := createHandler(db, userService)
	req, err := createRequest("/authentication/verifyLogin", userService)
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	req.Body = io_mock.NewReadCloserString(`{"login":"testuser","password":"testpassword"}`)
	req.Method = "POST"
	handler.ServeHTTP(resp, req)
	if err = AssertThat(resp.Status(), Is(200)); err != nil {
		t.Fatal(err)
	}
}
