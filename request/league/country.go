package league

import (
	"github.com/syurchen93/api-football-client/response/leagues"
	"github.com/syurchen93/api-football-client/response"
)

type Country struct {
	Name string `mapstructure:"name,omitempty"`
	Code string `validate:"omitempty,len=2" mapstructure:"code,omitempty"`
	Search string `validate:"omitempty,len=3" mapstructure:"search,omitempty"`
}

func (c Country) GetEndpoint() string {
	return "countries"
}

func (c Country) GetResponseStruct() response.ResponseInterface {
	return leagues.Country{}
}