package team

import (
	"github.com/syurchen93/api-football-client/response/league"
	"github.com/syurchen93/api-football-client/response"
)

type Season struct {
	Team int `mapstructure:"team" validate:"required"`
}

func (s Season) GetEndpoint() string {
	return "teams/seasons"
}

func (s Season) GetResponseStruct() response.ResponseInterface {
	return league.SeasonYear{}
}