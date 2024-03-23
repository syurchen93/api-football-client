package fixtures

import (
	"time"
	"github.com/syurchen93/api-football-client/common"
	"github.com/syurchen93/api-football-client/response/misc"
)

type Fixture struct {
	Fixture FixtureDetails `json:"fixture"`
	League  League         `json:"league"`
	Teams   Teams          `json:"teams"`
	Goals   Goals          `json:"goals"`
	Score   Score          `json:"score"`
}

type FixtureDetails struct {
	ID        int       `json:"id"`
	Referee   string    `json:"referee"`
	Timezone  string    `json:"timezone"`
	Date      time.Time `json:"date"`
	Timestamp int       `json:"timestamp"`
	Periods   Periods   `json:"periods"`
	Venue     misc.Venue `json:"venue"`
	Status    Status    `json:"status"`
}

type Periods struct {
	First  int `json:"first"`
	Second int `json:"second"`
}

type Venue struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type Status struct {
	Long    string `json:"long"`
	Value common.FixtureStatus `json:"short"`
	Elapsed int    `json:"elapsed"`
}

type League struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Logo    string `json:"logo"`
	Flag    string `json:"flag"`
	Season  int    `json:"season"`
	Round   string `json:"round"`
}

type Teams struct {
	Home Team `json:"home"`
	Away Team `json:"away"`
}

type Team struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Winner bool   `json:"winner"`
}

type Goals struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type Score struct {
	Halftime  Goals `json:"halftime"`
	Fulltime  Goals `json:"fulltime"`
	Extratime Goals `json:"extratime"`
	Penalty   Goals `json:"penalty"`
}