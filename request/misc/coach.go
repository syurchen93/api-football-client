package misc

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/misc"
)

type Coach struct {
	ID      int    `mapstructure:"id,omitempty"`
	TeamID  int    `mapstructure:"team,omitempty"`
	Search  string `mapstructure:"search,omitempty" validate:"omitempty,min=3"`
}

func (c Coach) GetEndpoint() string {
	return "coachs"
}

func (c Coach) GetResponseStruct() response.ResponseInterface {
	return misc.Coach{}
}

