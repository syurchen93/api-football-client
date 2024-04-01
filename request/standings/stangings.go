package standings

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/standings"
)

type Standings struct {
	League int `mapstructure:"league,omitempty"`
	Season int `mapstructure:"season" validate:"required,gte=1000,lte=9999"`
	Team   int `mapstructure:"team,omitempty"`
}

func (s Standings) GetEndpoint() string {
	return "standings"
}

func (s Standings) GetResponseStruct() response.ResponseInterface {
	return standings.Standings{}
}
