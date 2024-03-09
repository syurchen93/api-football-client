package league

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/league"
)

type League struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	CountryName string `json:"country,omitempty"`
	CountryCode    string `json:"code,omitempty" validate:"omitempty,len=2"`
	Season  int    `json:"season,omitempty" validate:"omitempty,gte=1000,lte=9999"`
	Team    int    `json:"team,omitempty"`
	Type    string `json:"type,omitempty" validate:"omitempty,oneof=league cup"`
	Current string `json:"current,omitempty" validate:"omitempty,oneof=true false"`
	Search  string `json:"search,omitempty" validate:"omitempty,min=3"`
	Last    int    `json:"last,omitempty" validate:"omitempty,max=99"`
}

func (l League) GetEndpoint() string {
	return "leagues"
}

func (l League) GetResponseStruct() response.ResponseInterface {
	return league.LeagueData{}
}