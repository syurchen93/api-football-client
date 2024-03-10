package team

import (
	//"time"
	"github.com/syurchen93/api-football-client/response/team"
	"github.com/syurchen93/api-football-client/response"
)

type Statistics struct {
	League int    `mapstructure:"league" validate:"required"`
	Season int    `mapstructure:"season" validate:"required,gte=1000,lte=9999"`
	Team   int    `mapstructure:"team" validate:"required"`
	//LimitDate time.Time `mapstructure:"date,omitempty" validate:"omitempty"` 
}

func (s Statistics) GetEndpoint() string {
	return "teams/statistics"
}

func (s Statistics) GetResponseStruct() response.ResponseInterface {
	return team.Statistics{}
}