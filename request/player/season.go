package player

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/leagues"
)

type PlayerSeason struct {
	PlayerID int `mapstructure:"player,omitempty"`
}

func (s PlayerSeason) GetEndpoint() string {
	return "players/seasons"
}

func (s PlayerSeason) GetResponseStruct() response.ResponseInterface {
	return leagues.SeasonYear{}
}
