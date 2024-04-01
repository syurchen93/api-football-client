package fixtures

import (
	"github.com/syurchen93/api-football-client/response/team"
)

type Predictions struct {
	Prediction FixturePrediction `mapstructure:"predictions"`
	League      League      `mapstructure:"league"`
	Teams       TeamStats       `mapstructure:"teams"`
	Comparison   Comparison  `mapstructure:"comparison"`
	HeadToHead   []HeadToHead  `mapstructure:"h2h"`
}

type FixturePrediction struct {
	Winner     Team   `mapstructure:"winner"`
	WinOrDraw  bool   `mapstructure:"win_or_draw"`
	UnderOver  *float32 `mapstructure:"under_over"`
	Goals      AverageGoals  `mapstructure:"goals"`
	Advice     string `mapstructure:"advice"`
	Percent    Percent `mapstructure:"percent"`
}

type AverageGoals struct {
	Home float32 `mapstructure:"home"`
	Away float32 `mapstructure:"away"`
}

type Percent struct {
	Home int `mapstructure:"home"`
	Draw int `mapstructure:"draw"`
	Away int `mapstructure:"away"`
}

type TeamStats struct {
	Home TeamDetails `mapstructure:"home"`
	Away TeamDetails `mapstructure:"away"`
}

type TeamDetails struct {
	ID     int    `mapstructure:"id"`
	Name   string `mapstructure:"name"`
	Logo   string `mapstructure:"logo"`
	Last5    Last5   `mapstructure:"last_5"`
	League   LeaguePerformance `mapstructure:"league"`
}

type Last5 struct {
	FormPercentage int `mapstructure:"form"`
	AttPercentage int `mapstructure:"att"`
	DefPercentage int `mapstructure:"def"`
	Goals GoalAmountData `mapstructure:"goals"`
}

type GoalAmountData struct {
	For  GoalAmount `mapstructure:"for"`
	Against GoalAmount `mapstructure:"against"`
}

type GoalAmount struct {
	Total int `mapstructure:"total"`
	Average float32 `mapstructure:"average"`
}

type LeaguePerformance struct {
	Form     string   `mapstructure:"form"`
	FormData FormData `mapstructure:"fixtures"`
	Goals    team.Goals `mapstructure:"goals"`
	Biggest  Biggest `mapstructure:"biggest"`
	CleanSheet Details `mapstructure:"clean_sheet"`
	FailedToScore Details `mapstructure:"failed_to_score"`
	PenaltyStats TeamPenaltyStats `mapstructure:"penalty"`
	LineupUsage []LineupUsage `mapstructure:"lineups"`
	Cards     team.Cards `mapstructure:"cards"`
}

type TeamPenaltyStats struct {
	Scored AverageCount `mapstructure:"scored"`
	Missed AverageCount `mapstructure:"missed"`
	Total int `mapstructure:"total"`
}

type AverageCount struct {
	Total int `mapstructure:"total"`
	Average float32 `mapstructure:"percentage"`
}

type LineupUsage struct {
	Formation string `mapstructure:"formation"`
	PlayedCount int `mapstructure:"played"`
}

type FormData struct {
    Played Details `mapstructure:"played"`
    Wins   Details `mapstructure:"wins"`
    Draws  Details `mapstructure:"draws"`
    Loses  Details `mapstructure:"loses"`
}

type Details struct {
    Home  int `mapstructure:"home"`
    Away  int `mapstructure:"away"`
    Total int `mapstructure:"total"`
}

type AverageDetails struct {
    Home  float32 `mapstructure:"home"`
    Away  float32 `mapstructure:"away"`
    Total float32 `mapstructure:"total"`
}

type Biggest struct {
    Streak Streak `mapstructure:"streak"`
    Wins   BiggestScore `mapstructure:"wins"`
    Loses  BiggestScore `mapstructure:"loses"`
    Goals  GoalsForAgainst  `mapstructure:"goals"`
}

type GoalsForAgainst struct {
	For     Goals `mapstructure:"for"`
	Against Goals `mapstructure:"against"`
}

type Streak struct {
    Wins  int `mapstructure:"wins"`
    Draws int `mapstructure:"draws"`
    Loses int `mapstructure:"loses"`
}

type BiggestScore struct {
    Home string `mapstructure:"home"`
    Away string `mapstructure:"away"`
}

type Comparison struct {
    Form                Percent `mapstructure:"form"`
    Att                 Percent `mapstructure:"att"`
    Def                 Percent `mapstructure:"def"`
    PoissonDistribution Percent `mapstructure:"poisson_distribution"`
    H2H                 Percent `mapstructure:"h2h"`
    Goals               Percent `mapstructure:"goals"`
    Total               Percent `mapstructure:"total"`
}