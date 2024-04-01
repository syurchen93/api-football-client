package league

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/leagues"
)

type Season struct {
}

func (s Season) GetEndpoint() string {
	return "leagues/seasons"
}

func (s Season) GetResponseStruct() response.ResponseInterface {
	return leagues.SeasonYear{}
}
