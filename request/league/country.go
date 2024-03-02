package league

type Country struct {
	Name string `json:"name"`
	Code string `validate:"len=2" json:"code"`
	Search string `validate:"len=3" json:"search"`
}

func (c Country) GetEndpoint() string {
	return "countries"
}
