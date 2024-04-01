package misc

import (
	"time"
)

type InjuryInfo struct {
	Player  PlayerInjured `json:"player"`
	Team    Team          `json:"team"`
	Fixture Fixture       `json:"fixture"`
	League  League        `json:"league"`
}

type PlayerInjured struct {
	ID     int        `json:"id"`
	Name   string     `json:"name"`
	Photo  string     `json:"photo"`
	Type   InjuryType `json:"type"`
	Reason string     `json:"reason"`
}

type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type Fixture struct {
	ID        int       `json:"id"`
	Timezone  string    `json:"timezone"`
	Date      time.Time `json:"date"`
	Timestamp int       `json:"timestamp"`
}

type League struct {
	ID      int     `json:"id"`
	Season  int     `json:"season"`
	Name    string  `json:"name"`
	Country string  `json:"country"`
	Logo    string  `json:"logo"`
	Flag    *string `json:"flag"`
}

type InjuryType string

const (
	MissingFixture InjuryType = "Missing Fixture"
	Questionable   InjuryType = "Questionable"
)
