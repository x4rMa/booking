package main

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestCreateServer(t *testing.T) {
	server := createServer(":45678", "/tmp")
	err := AssertThat(server, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
}
