package date

type Date struct {
	Id    int    `json:"id" sql:"AUTO_INCREMENT"`
	Start string `json:"start"`
	End   string `json:"end"`
}
