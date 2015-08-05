package main

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestCreateServer(t *testing.T) {
	server := createServer(":45678", "/tmp", "dbname", "dbuser", "dbpass", true)
	err := AssertThat(server, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
}
