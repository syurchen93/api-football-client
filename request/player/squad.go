package player

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/players"
)

type Squad struct {
	TeamID   int `mapstructure:"team,omitempty"`
	PlayerID int `mapstructure:"player,omitempty"`
}

func (s Squad) GetEndpoint() string {
	return "players/squads"
}

func (s Squad) GetResponseStruct() response.ResponseInterface {
	return players.Squad{}
}
