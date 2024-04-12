package misc

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/misc"
)

type Sideilined struct {
	PlayerID int `mapstructure:"player,omitempty" validate:"required_without=CoachID,excluded_with=CoachID"`
	CoachID  int `mapstructure:"coach,omitempty" validate:"required_without=PlayerID,excluded_with=PlayerID"`
}

func (s Sideilined) GetEndpoint() string {
	return "sidelined"
}

func (s Sideilined) GetResponseStruct() response.ResponseInterface {
	return misc.SidelinedEntry{}
}
