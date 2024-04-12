package player

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/players"
)

type TopScorer struct {
	LeagueID int `mapstructure:"league" validate:"required"`
	Season   int `mapstructure:"season" validate:"required,gte=1000,lte=9999"`
}

func (p TopScorer) GetEndpoint() string {
	return "players/topscorers"
}

func (p TopScorer) GetResponseStruct() response.ResponseInterface {
	return players.PlayerInfo{}
}
