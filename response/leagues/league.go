package leagues

import (
	"time"
)

type LeagueData struct {
	League  League   `json:"league"`
	Country Country  `json:"country"`
	Seasons []Season `json:"seasons"`
}

type League struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country,omitempty"`
	Type    string `json:"type"`
	Logo    string `json:"logo"`
	Flag    string `json:"flag,omitempty"`
	Season  int    `json:"season,omitempty"`
}

type Season struct {
	Year     int       `json:"year"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Current  bool      `json:"current"`
	Coverage Coverage  `json:"coverage"`
}

type Coverage struct {
	Fixtures struct {
		Events             bool `json:"events"`
		Lineups            bool `json:"lineups"`
		StatisticsFixtures bool `json:"statistics_fixtures"`
		StatisticsPlayers  bool `json:"statistics_players"`
	} `json:"fixtures"`
	Standings   bool `json:"standings"`
	Players     bool `json:"players"`
	TopScorers  bool `json:"top_scorers"`
	TopAssists  bool `json:"top_assists"`
	TopCards    bool `json:"top_cards"`
	Injuries    bool `json:"injuries"`
	Predictions bool `json:"predictions"`
	Odds        bool `json:"odds"`
}
