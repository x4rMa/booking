package shooting

type Shooting interface {
	GetName() string
	SetName(name string)
}

type shooting struct {
	Name string `json:"name"`
}

func New() *shooting {
	return new(shooting)
}

func (d *shooting) SetName(name string) {
	d.Name = name
}

func (d *shooting) GetName() string {
	return d.Name
}
