package fixture

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/fixtures"
)

type Round struct {
	League int `mapstructure:"league" validate:"required"`
	Season int `mapstructure:"season" validate:"required,gte=1000,lte=9999"`
	Current bool `mapstructure:"current"`
}

func (r Round) GetEndpoint() string {
	return "fixtures/rounds"
}

func (r Round) GetResponseStruct() response.ResponseInterface {
	return fixtures.Round{}
}