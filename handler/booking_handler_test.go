package handler

import (
	"net/http"
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/booking/database"

	server_mock "github.com/bborbe/server/mock"
	io_mock "github.com/bborbe/io/mock"


	booking_date_service "github.com/bborbe/booking/date/service"
	booking_date_storage "github.com/bborbe/booking/date/storage"

	booking_model_service "github.com/bborbe/booking/model/service"
	booking_model_storage "github.com/bborbe/booking/model/storage"


	booking_shooting_service "github.com/bborbe/booking/shooting/service"
	booking_shooting_storage "github.com/bborbe/booking/shooting/storage"

	booking_tokengenerator "github.com/bborbe/booking/tokengenerator"
)

func TestNewHandlerImplementsHttpHandler(t *testing.T) {
	r := createHandler()
	var i (*http.Handler) = nil
	err := AssertThat(r, Implements(i).Message("check implements server.Server"))
	if err != nil {
		t.Fatal(err)
	}
}
func createHandler() http.Handler {
	db := database.New("/tmp/booking_test.db", true)
	tokengenerator := booking_tokengenerator.New()
	dateService := booking_date_service.New(booking_date_storage.New(db))
	modelService := booking_model_service.New(booking_model_storage.New(db), tokengenerator)
	shootingService := booking_shooting_service.New(booking_shooting_storage.New(db))
	return NewHandler("/tmp", dateService, modelService, shootingService)
}

func TestDate(t *testing.T) {
	handler := createHandler()
	resp := server_mock.NewHttpResponseWriterMock()
	req, err := server_mock.NewHttpRequestMock("/date")
	if err = AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(resp, req)
	if err = AssertThat(resp.Status(), Is(200)); err != nil {
		t.Fatal(err)
	}
}

func TestGetDate(t *testing.T) {
	handler := createHandler()
	resp := server_mock.NewHttpResponseWriterMock()
	req, err := server_mock.NewHttpRequestMock("/date")
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
	handler := createHandler()
	resp := server_mock.NewHttpResponseWriterMock()
	req, err := server_mock.NewHttpRequestMock("/date")
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
