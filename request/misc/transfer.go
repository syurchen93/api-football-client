package misc

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/misc"
)

type Transfer struct {
	ID     int `mapstucture:"id,omitempty"`
	TeamID int `mapstructure:"team,omitempty"`
}

func (t Transfer) GetEndpoint() string {
	return "transfers"
}

func (t Transfer) GetResponseStruc() response.ResponseInterface {
	return misc.Transfer{}
}
