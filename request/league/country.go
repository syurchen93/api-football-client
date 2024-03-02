package league

import (
	"github.com/syurchen93/api-football-client/response/league"
	"github.com/syurchen93/api-football-client/response"
)

type Country struct {
	Name string `json:"name,omitempty"`
	Code string `validate:"omitempty,len=2" json:"code,omitempty"`
	Search string `validate:"omitempty,len=3" json:"search,omitempty"`
}

func (c Country) GetEndpoint() string {
	return "countries"
}

func (c Country) GetResponseStruct() response.ResponseInterface {
	return league.Country{}
}