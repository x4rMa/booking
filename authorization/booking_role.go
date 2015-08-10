package authorization

type Role string

const (
	Administrator Role = "administrator"
	Organizer     Role = "organizer"
	Participant   Role = "participant"
	None          Role = "none"
)

func RoleByName(roleName string) *Role {
	roles := []Role{Administrator, Organizer, Participant}
	for _, r := range roles {
		if string(r) == roleName {
			return &r
		}
	}
	return nil
}
