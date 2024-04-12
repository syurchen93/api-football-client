package player

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/players"
)

type TopYellowCards struct {
	LeagueID int `mapstructure:"league" validate:"required"`
	Season   int `mapstructure:"season" validate:"required,gte=1000,lte=9999"`
}

func (tyc TopYellowCards) GetEndpoint() string {
	return "players/topyellowcards"
}

func (tyc TopYellowCards) GetResponseStruct() response.ResponseInterface {
	return players.PlayerInfo{}
}
