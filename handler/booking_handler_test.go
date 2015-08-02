package handler

import (
	"net/http"
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/booking/database"
	booking_date_service "github.com/bborbe/booking/date/service"
	booking_date_storage "github.com/bborbe/booking/date/storage"
	"github.com/bborbe/server/mock"
	io_mock "github.com/bborbe/io/mock"
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
	return NewHandler("/tmp", booking_date_service.New(booking_date_storage.New(database.New("/tmp/booking_test.db", true))))
}

func TestDate(t *testing.T) {
	handler := createHandler()
	resp := mock.NewHttpResponseWriterMock()
	req, err := mock.NewHttpRequestMock("/date")
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
	resp := mock.NewHttpResponseWriterMock()
	req, err := mock.NewHttpRequestMock("/date")
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
	resp := mock.NewHttpResponseWriterMock()
	req, err := mock.NewHttpRequestMock("/date")
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
