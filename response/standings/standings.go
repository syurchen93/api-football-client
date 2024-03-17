package standings

import (
	"github.com/syurchen93/api-football-client/response/team"
	"time"
)

type Standings struct {
	League League `json:"league"`
}

type League struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Country  string    `json:"country"`
	Logo     string    `json:"logo"`
	Flag     string    `json:"flag"`
	Season   int       `json:"season"`
	Standings [][]Ranking `json:"standings"`
}

type Ranking struct {
	Rank        int    `json:"rank"`
	Team        team.Team   `json:"team"`
	Points      int    `json:"points"`
	GoalsDiff   int    `json:"goalsDiff"`
	Group       string `json:"group"`
	Form        string `json:"form"`
	Status      string `json:"status"`
	Description string `json:"description"`
	All         Match  `json:"all"`
	Home        Match  `json:"home"`
	Away        Match  `json:"away"`
	Updated     time.Time `json:"update" mapstructure:"update"`
}

type Match struct {
	Played int   `json:"played"`
	Win    int   `json:"win"`
	Draw   int   `json:"draw"`
	Lose   int   `json:"lose"`
	Goals  Goals `json:"goals"`
}

type Goals struct {
	For     int `json:"for"`
	Against int `json:"against"`
}