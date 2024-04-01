package league

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/leagues"
)

type League struct {
	ID          int    `mapstructure:"id,omitempty"`
	Name        string `mapstructure:"name,omitempty"`
	CountryName string `mapstructure:"country,omitempty"`
	CountryCode string `mapstructure:"code,omitempty" validate:"omitempty,len=2"`
	Season      int    `mapstructure:"season,omitempty" validate:"omitempty,gte=1000,lte=9999"`
	Team        int    `mapstructure:"team,omitempty"`
	Type        string `mapstructure:"type,omitempty" validate:"omitempty,oneof=league cup"`
	Current     bool   `mapstructure:"current,omitempty" validate:"omitempty"`
	Search      string `mapstructure:"search,omitempty" validate:"omitempty,min=3"`
	Last        int    `mapstructure:"last,omitempty" validate:"omitempty,max=99"`
}

func (l League) GetEndpoint() string {
	return "leagues"
}

func (l League) GetResponseStruct() response.ResponseInterface {
	return leagues.LeagueData{}
}
