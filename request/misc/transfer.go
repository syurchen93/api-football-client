package misc

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/misc"
)

type Transfer struct {
	PlayerID int `mapstructure:"player,omitempty"`
	TeamID   int `mapstructure:"team,omitempty"`
}

func (t Transfer) GetEndpoint() string {
	return "transfers"
}

func (t Transfer) GetResponseStruct() response.ResponseInterface {
	return misc.Transfer{}
}
