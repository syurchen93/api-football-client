package fixture

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/fixtures"
)

type Predictions struct {
	FixtureID int `mapstructure:"fixture" validate:"required"`
}

func (r Predictions) GetEndpoint() string {
	return "predictions"
}

func (r Predictions) GetResponseStruct() response.ResponseInterface {
	return fixtures.Predictions{}
}
