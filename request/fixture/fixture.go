package fixture

import (
	"github.com/syurchen93/api-football-client/common"
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/fixtures"
	"time"
)

type Fixture struct {
	ID       int                    `mapstructure:"id,omitempty"`
	IDs      []int                  `mapstructure:"ids,omitempty" validate:"omitempty,max=20"`
	Live     string                 `mapstructure:"live,omitempty" validate:"omitempty,oneof=all id-id"`
	Date     time.Time              `mapstructure:"date,omitempty"`
	League   int                    `mapstructure:"league,omitempty"`
	Season   int                    `mapstructure:"season,omitempty" validate:"omitempty,gte=1000,lte=9999"`
	Team     int                    `mapstructure:"team,omitempty"`
	Last     int                    `mapstructure:"last,omitempty" validate:"omitempty,max=99"`
	Next     int                    `mapstructure:"next,omitempty" validate:"omitempty,max=99"`
	From     time.Time              `mapstructure:"from,omitempty"`
	To       time.Time              `mapstructure:"to,omitempty"`
	Round    string                 `mapstructure:"round,omitempty"`
	Statuses []common.FixtureStatus `mapstructure:"status,omitempty"`
	Venue    int                    `mapstructure:"venue,omitempty"`
	Timezone string                 `mapstructure:"timezone,omitempty"`
}

func (f Fixture) GetEndpoint() string {
	return "fixtures"
}

func (f Fixture) GetResponseStruct() response.ResponseInterface {
	return fixtures.Fixture{}
}
