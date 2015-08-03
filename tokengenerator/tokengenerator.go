package tokengenerator

type tokengenerator struct{}

func New() *tokengenerator {
	return new(tokengenerator)
}

func (t *tokengenerator) GenerateToken() (string, error) {
	return "abc", nil
}
