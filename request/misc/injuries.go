package misc

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/misc"
	"time"
)

type Injuries struct {
	LeagueID  int       `mapstructure:"league,omitempty" validate:"omitempty"`
	Season    int       `mapstructure:"season,omitempty" validate:"omitempty,gte=1000,lte=9999"`
	FixtureID int       `mapstructure:"fixture,omitempty" validate:"omitempty"`
	TeamID    int       `mapstructure:"team,omitempty" validate:"omitempty"`
	PlayerID  int       `mapstructure:"player,omitempty" validate:"omitempty"`
	Date      time.Time `mapstructure:"date,omitempty" validate:"omitempty"`
	Timezone  string    `mapstructure:"timezone,omitempty" validate:"omitempty"`
}

func (i Injuries) GetEndpoint() string {
	return "injuries"
}

func (i Injuries) GetResponseStruct() response.ResponseInterface {
	return misc.InjuryInfo{}
}
