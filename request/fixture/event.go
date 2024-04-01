package fixture

import (
	"github.com/syurchen93/api-football-client/common"
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/fixtures"
)

type Event struct {
	FixtureID int              `mapstructure:"fixture" validate:"required"`
	TeamID    int              `mapstructure:"team,omitempty"`
	PlayerID  int              `mapstructure:"player,omitempty"`
	Type      common.EventType `mapstructure:"type,omitempty"`
}

func (e Event) GetEndpoint() string {
	return "fixtures/events"
}

func (e Event) GetResponseStruct() response.ResponseInterface {
	return fixtures.Event{}
}
