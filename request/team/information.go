package team

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/team"
)

type Team struct {
	ID      int    `json:"id,omitempty" validate:"omitempty"`
	Name    string `json:"name,omitempty" validate:"omitempty"`
	League  int    `json:"league,omitempty" validate:"omitempty"`
	Season  int    `json:"season,omitempty" validate:"omitempty,gte=1000,lte=9999"`
	Country string `json:"country,omitempty" validate:"omitempty"`
	CountryCode    string `json:"code,omitempty" validate:"omitempty,len=3"`
	Venue   int    `json:"venue,omitempty" validate:"omitempty"`
	Search  string `json:"search,omitempty" validate:"omitempty,min=3"`
}

func (t Team) GetEndpoint() string {
	return "teams"
}

func (t Team) GetResponseStruct() response.ResponseInterface{
	return team.Information{}
}