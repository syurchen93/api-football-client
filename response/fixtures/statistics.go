package fixtures

import (
	"time"

	"github.com/syurchen93/api-football-client/common"
)

type TeamStatistics struct {
	Team       TeamBasic       `json:"team"`
	Statistics []TeamStatistic `json:"statistics"`
}

type PlayerStatistics struct {
	Team    TeamBasic         `json:"team"`
	Players []TeamPlayerStats `mapstructure:"players" json:"players"`
}



type TeamStatistic struct {
	Type  common.StatsType `json:"type"`
	Value int              `json:"value"`
}

type TeamBasic struct {
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	Logo    string     `json:"logo"`
	Updated *time.Time `mapstructue:"update,omitempty" json:"updated,omitempty"`
}

type TeamPlayerStats struct {
	Player     Player               `json:"player"`
	Statistics PlayerGameStatistics `json:"statistics"`
}

type PlayerGameStatistics struct {
	General  GeneralStats  `mapstructure:"games" json:"genaral"`
	Offsides int           `json:"offsides"`
	Shots    ShotsStats    `json:"shots"`
	Goals    GoalStats     `json:"goals"`
	Passes   PassesStats   `json:"passes"`
	Tackles  TacklesStats  `json:"tackles"`
	Duels    DuelsStats    `json:"duels"`
	Dribbles DribblesStats `json:"dribbles"`
	Fouls    FoulsStats    `json:"fouls"`
	Cards    CardsStats    `json:"cards"`
	Penalty  PenaltyStats  `json:"penalty"`
}

type GeneralStats struct {
	Minutes    int            `json:"minutes"`
	Number     int            `json:"number"`
	Position   LineupPosition `json:"position"`
	Rating     float32        `json:"rating"`
	Captain    bool           `json:"captain"`
	Substitute bool           `json:"substitute"`
}

type ShotsStats struct {
	Total int `json:"total"`
	On    int `json:"on"`
}

type GoalStats struct {
	Total    int `json:"total"`
	Conceded int `json:"conceded"`
	Assists  int `json:"assists"`
	Saves    int `json:"saves"`
}

type PassesStats struct {
	Total           int `json:"total"`
	Key             int `json:"key"`
	AccuracyPercent int `mapstructure:"accuracy" json:"accuracy"`
}

type TacklesStats struct {
	Total         int `json:"total"`
	Blocks        int `json:"blocks"`
	Interceptions int `json:"interceptions"`
}

type DuelsStats struct {
	Total int `json:"total"`
	Won   int `json:"won"`
}

type DribblesStats struct {
	Attempts int `json:"attempts"`
	Success  int `json:"success"`
	Past     int `json:"past"`
}

type FoulsStats struct {
	Drawn     int `json:"drawn"`
	Committed int `json:"committed"`
}

type CardsStats struct {
	Yellow int `json:"yellow"`
	Red    int `json:"red"`
}

type PenaltyStats struct {
	Won      int `json:"won"`
	Committed int `mapstructure:"commited" json:"committed"`
	Scored   int `json:"scored"`
	Missed   int `json:"missed"`
	Saved    int `json:"saved"`
}
