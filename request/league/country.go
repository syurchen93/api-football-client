package league

import (
	"github.com/syurchen93/api-football-client/response/league"
	"github.com/syurchen93/api-football-client/response"
)

type Country struct {
	Name string `json:"name"`
	Code string `validate:"len=2" json:"code"`
	Search string `validate:"len=3" json:"search"`
}

func (c Country) GetEndpoint() string {
	return "countries"
}

func (c Country) GetResponseStruct() response.ResponseInterface {
	return league.Country{}
}