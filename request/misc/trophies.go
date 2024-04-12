package misc

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/misc"
)

type Trophies struct {
	PlayerID int `mapstructure:"player,omitempty" validate:"required_without=CoachID,excluded_with=CoachID"`
	CoachID  int `mapstructure:"coach,omitempty" validate:"required_without=PlayerID,excluded_with=PlayerID"`
}

func (p Trophies) GetEndpoint() string {
	return "trophies"
}

func (p Trophies) GetResponseStruct() response.ResponseInterface {
	return misc.Trophy{}
}
