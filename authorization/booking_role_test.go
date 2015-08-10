package authorization

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestRoleAdministrator(t *testing.T) {
	if err := AssertThat(string(Administrator), Is("administrator")); err != nil {
		t.Fatal(err)
	}
}

func TestRoleOrganizer(t *testing.T) {
	if err := AssertThat(string(Organizer), Is("organizer")); err != nil {
		t.Fatal(err)
	}
}

func TestRoleParticipant(t *testing.T) {
	if err := AssertThat(string(Participant), Is("participant")); err != nil {
		t.Fatal(err)
	}
}

func TestRoleByNameIllegal(t *testing.T) {
	role := RoleByName("asdf")
	if err := AssertThat(role == nil, Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestRoleByNameAdministrator(t *testing.T) {
	var role *Role
	role = RoleByName("administrator")
	if err := AssertThat(*role, Is(Administrator)); err != nil {
		t.Fatal(err)
	}
}

func TestRoleByNameOrganizer(t *testing.T) {
	var role *Role
	role = RoleByName("organizer")
	if err := AssertThat(*role, Is(Organizer)); err != nil {
		t.Fatal(err)
	}
}

func TestRoleByNameParticipant(t *testing.T) {
	var role *Role
	role = RoleByName("participant")
	if err := AssertThat(*role, Is(Participant)); err != nil {
		t.Fatal(err)
	}
}
