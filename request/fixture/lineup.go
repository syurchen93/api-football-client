package fixture

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/fixtures"
)

type Lineup struct {
	FixtureID int        `mapstructure:"fixture" validate:"required"`
	TeamID    int        `mapstructure:"team,omitempty"`
	PlayerID  int        `mapstructure:"player,omitempty"`
	Type      LineupType `mapstructure:"type,omitempty"`
}

type LineupType string

const (
	StartingXI LineupType = "startxi"
	Substitutes LineupType = "substitutes"
	Formation LineupType = "formation"
	Coach LineupType = "coach"
)

func (l Lineup) GetEndpoint() string {
	return "fixtures/lineups"
}

func (l Lineup) GetResponseStruct() response.ResponseInterface {
	return fixtures.Lineup{}
}