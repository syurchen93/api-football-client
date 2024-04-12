package player

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/players"
)

type TopRedCards struct {
	LeagueID int `mapstructure:"league" validate:"required"`
	Season   int `mapstructure:"season" validate:"required,gte=1000,lte=9999"`
}

func (trc TopRedCards) GetEndpoint() string {
	return "players/topredcards"
}

func (trc TopRedCards) GetResponseStruct() response.ResponseInterface {
	return players.PlayerInfo{}
}
