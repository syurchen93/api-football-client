package player

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/players"
)

type PlayerInfo struct {
	ID       int `mapstructure:"id,omitempty"`
	TeamID   int `mapstructure:"team,omitempty"`
	LeagueID int `mapstructure:"league,omitempty"`
	// Requires the fields Id, League or Team
	Season int `mapstructure:"season,omitempty" validate:"omitempty,gte=1000,lte=9999"`
	// Requires the fields League or Team
	Search string `mapstructure:"search,omitempty" validate:"omitempty,min=4"`
	Page   int    `mapstructure:"page,omitempty"`
}

func (p PlayerInfo) GetEndpoint() string {
	return "players"
}

func (p PlayerInfo) GetResponseStruct() response.ResponseInterface {
	return players.PlayerInfo{}
}
