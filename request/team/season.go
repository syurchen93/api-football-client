package team

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/leagues"
)

type TeamSeason struct {
	Team int `mapstructure:"team" validate:"required"`
}

func (s TeamSeason) GetEndpoint() string {
	return "teams/seasons"
}

func (s TeamSeason) GetResponseStruct() response.ResponseInterface {
	return leagues.SeasonYear{}
}
