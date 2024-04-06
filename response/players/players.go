package players

import (
	"github.com/martinlindhe/unit"
	"github.com/syurchen93/api-football-client/response/fixtures"
	"github.com/syurchen93/api-football-client/response/leagues"
	"time"
)

type PlayerInfo struct {
	Player     Player       `json:"player"`
	Statistics []Statistics `json:"statistics"`
}

type Player struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Firstname   string      `json:"firstname"`
	Lastname    string      `json:"lastname"`
	Age         int         `json:"age"`
	Birth       Birth       `json:"birth"`
	Nationality string      `json:"nationality"`
	Height      unit.Length `json:"height"`
	Weight      unit.Mass   `json:"weight"`
	Injured     bool        `json:"injured"`
	Photo       string      `json:"photo"`
}

type Birth struct {
	Date    time.Time `json:"date"`
	Place   string    `json:"place"`
	Country string    `json:"country"`
}

type Statistics struct {
	Team        fixtures.TeamBasic     `json:"team"`
	League      leagues.League         `json:"league"`
	Games       Games                  `json:"games"`
	Substitutes Substitutes            `json:"substitutes"`
	Shots       fixtures.ShotsStats    `json:"shots"`
	Goals       fixtures.GoalStats     `json:"goals"`
	Passes      fixtures.PassesStats   `json:"passes"`
	Tackles     fixtures.TacklesStats  `json:"tackles"`
	Duels       fixtures.DuelsStats    `json:"duels"`
	Dribbles    fixtures.DribblesStats `json:"dribbles"`
	Fouls       fixtures.FoulsStats    `json:"fouls"`
	Cards       Cards                  `json:"cards"`
	Penalty     fixtures.PenaltyStats  `json:"penalty"`
}

type Games struct {
	Appearances int     `json:"appearences"`
	Lineups     int     `json:"lineups"`
	Minutes     int     `json:"minutes"`
	Number      int     `json:"number"`
	Position    string  `json:"position"`
	Rating      float32 `json:"rating"`
	Captain     bool    `json:"captain"`
}

type Substitutes struct {
	In    int `json:"in"`
	Out   int `json:"out"`
	Bench int `json:"bench"`
}

type Cards struct {
	Yellow    int `json:"yellow"`
	YellowRed int `json:"yellowred"`
	Red       int `json:"red"`
}
