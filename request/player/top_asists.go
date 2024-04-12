package player

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/players"
)

type TopAssists struct {
	LeagueID int `mapstructure:"league" validate:"required"`
	Season   int `mapstructure:"season" validate:"required,gte=1000,lte=9999"`
}

func (ta TopAssists) GetEndpoint() string {
	return "players/topassists"
}

func (ta TopAssists) GetResponseStruct() response.ResponseInterface {
	return players.PlayerInfo{}
}
