package tokengenerator

import (
	uuid "github.com/nu7hatch/gouuid"
)

type tokengenerator struct{}

func New() *tokengenerator {
	return new(tokengenerator)
}

func (t *tokengenerator) GenerateToken() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
