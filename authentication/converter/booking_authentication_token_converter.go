package converter

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	booking_authentication "github.com/bborbe/booking/authentication"
)

const AUTH_HEADER_FIELD = "X-Auth-Token"

type Converter interface {
	HttpRequestToAuthentication(request *http.Request) (*booking_authentication.Authentication, error)
	TokenToAuthentication(token string) (*booking_authentication.Authentication, error)
	AuthenticationToToken(authentication *booking_authentication.Authentication) (string, error)
}

type converter struct{}

func New() *converter {
	return new(converter)
}

func (c *converter) HttpRequestToAuthentication(request *http.Request) (*booking_authentication.Authentication, error) {
	values := request.Header[AUTH_HEADER_FIELD]
	if values == nil {
		return nil, fmt.Errorf("header field %s missing", AUTH_HEADER_FIELD)
	}
	if len(values) != 1 || len(values[0]) == 0 {
		return nil, fmt.Errorf("header field %s empty", AUTH_HEADER_FIELD)
	}
	return c.TokenToAuthentication(values[0])
}

func (c *converter) TokenToAuthentication(token string) (*booking_authentication.Authentication, error) {
	content, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	var f booking_authentication.Authentication
	err = json.Unmarshal(content, &f)
	if err != nil {
		return nil, err
	}
	return &f, err
}

func (c *converter) AuthenticationToToken(authentication *booking_authentication.Authentication) (string, error) {
	b, err := json.Marshal(authentication)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
