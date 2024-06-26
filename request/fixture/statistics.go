package fixture

import (
	"github.com/syurchen93/api-football-client/common"
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/fixtures"
)

type TeamStatistics struct {
	FixtureID int `mapstructure:"fixture" validate:"required"`
	// Team id is required even though documentation says it's optional
	TeamID int              `mapstructure:"team" validate:"required"`
	Type   common.StatsType `mapstructure:"type"`
}

type PlayerStatistics struct {
	FixtureID int `mapstructure:"fixture" validate:"required"`
	// Team id is required even though documentation says it's optional
	TeamID int `mapstructure:"team,omitempty"`
}

func (ts TeamStatistics) GetEndpoint() string {
	return "fixtures/statistics"
}

func (ts TeamStatistics) GetResponseStruct() response.ResponseInterface {
	return fixtures.TeamStatistics{}
}

func (ps PlayerStatistics) GetEndpoint() string {
	return "fixtures/players"
}

func (ps PlayerStatistics) GetResponseStruct() response.ResponseInterface {
	return fixtures.PlayerStatistics{}
}
