package fixture

import (
	"github.com/syurchen93/api-football-client/common"
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/fixtures"
	"time"
)

type HeadToHead struct {
	H2H      []int                  `mapstructure:"h2h" validate:"required,len=2"`
	Date     time.Time              `mapstructure:"date,omitempty"`
	League   int                    `mapstructure:"league,omitempty"`
	Season   int                    `mapstructure:"season,omitempty" validate:"omitempty,gte=1000,lte=9999"`
	Last     int                    `mapstructure:"last,omitempty"`
	Next     int                    `mapstructure:"next,omitempty"`
	From     time.Time              `mapstructure:"from,omitempty"`
	To       time.Time              `mapstructure:"to,omitempty"`
	Statuses []common.FixtureStatus `mapstructure:"status,omitempty"`
	Venue    int                    `mapstructure:"venue,omitempty"`
	Timezone string                 `mapstructure:"timezone,omitempty"`
}

func (h2h HeadToHead) GetEndpoint() string {
	return "fixtures/headtohead"
}

func (h2h HeadToHead) GetResponseStruct() response.ResponseInterface {
	return fixtures.HeadToHead{}
}
