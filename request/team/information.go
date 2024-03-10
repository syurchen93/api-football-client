package team

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/team"
)

type Team struct {
	ID      int    `mapstructure:"id,omitempty" validate:"omitempty"`
	Name    string `mapstructure:"name,omitempty" validate:"omitempty"`
	League  int    `mapstructure:"league,omitempty" validate:"omitempty"`
	Season  int    `mapstructure:"season,omitempty" validate:"omitempty,gte=1000,lte=9999"`
	Country string `mapstructure:"country,omitempty" validate:"omitempty"`
	CountryCode    string `mapstructure:"code,omitempty" validate:"omitempty,len=3"`
	Venue   int    `mapstructure:"venue,omitempty" validate:"omitempty"`
	Search  string `mapstructure:"search,omitempty" validate:"omitempty,min=3"`
}

func (t Team) GetEndpoint() string {
	return "teams"
}

func (t Team) GetResponseStruct() response.ResponseInterface{
	return team.Information{}
}