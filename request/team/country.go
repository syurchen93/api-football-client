package team

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/leagues"
)

type Country struct {
}

func (c Country) GetEndpoint() string {
	return "teams/countries"
}

func (c Country) GetResponseStruct() response.ResponseInterface {
	return leagues.Country{}
}
