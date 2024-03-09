package league

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/league"
)

type Season struct {
}

func (s Season) GetEndpoint() string {
	return "leagues/seasons"
}

func (s Season) GetResponseStruct() response.ResponseInterface {
	return league.SeasonYear{}
}