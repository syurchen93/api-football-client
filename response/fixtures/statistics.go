package fixtures

import (
	"github.com/syurchen93/api-football-client/common"
)

type TeamStatistics struct {
	Team       TeamBasic `json:"team"`
	Statistics []Statistic `json:"statistics"`
}

type Statistic struct {
	Type  common.StatsType `json:"type"`
	Value int `json:"value"`
}

type TeamBasic struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}