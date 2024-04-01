package team

import (
	"github.com/syurchen93/api-football-client/response/leagues"
)

type Statistics struct {
	League        leagues.League `json:"league"`
	Team          Team           `json:"team"`
	FormString    string         `json:"form" mapstructure:"form"`
	Fixtures      Fixtures       `json:"fixtures"`
	Goals         Goals          `json:"goals"`
	Biggest       Biggest        `json:"biggest"`
	CleanSheet    CleanSheet     `json:"clean_sheet"`
	FailedToScore FailedToScore  `json:"failed_to_score"`
	Penalty       Penalty        `json:"penalty"`
	Lineups       []Lineup       `json:"lineups"`
	Cards         Cards          `json:"cards"`
}

type Fixtures struct {
	Played FixtureDetails `json:"played"`
	Wins   FixtureDetails `json:"wins"`
	Draws  FixtureDetails `json:"draws"`
	Loses  FixtureDetails `json:"loses"`
}

type FixtureDetails struct {
	Home  int `json:"home"`
	Away  int `json:"away"`
	Total int `json:"total"`
}

type Goals struct {
	For     GoalDetails `json:"for"`
	Against GoalDetails `json:"against"`
}

type GoalDetails struct {
	Total   GoalTotalDetails   `json:"total"`
	Average GoalAverage        `json:"average"`
	Minute  MinuteDistribution `json:"minute"`
}

type GoalTotalDetails struct {
	Home  int `json:"home"`
	Away  int `json:"away"`
	Total int `json:"total"`
}

type GoalAverage struct {
	Home  float32 `json:"home"`
	Away  float32 `json:"away"`
	Total float32 `json:"total"`
}

type MinuteDistribution struct {
	ZeroToFifteen                   MinuteDetails `json:"0-15" mapstructure:"0-15"`
	SixteenToThirty                 MinuteDetails `json:"16-30" mapstructure:"16-30"`
	ThirtyOneToFortyFive            MinuteDetails `json:"31-45" mapstructure:"31-45"`
	FortySixToSixty                 MinuteDetails `json:"46-60" mapstructure:"46-60"`
	SixtyOneToSeventyFive           MinuteDetails `json:"61-75" mapstructure:"61-75"`
	SeventySixToNinety              MinuteDetails `json:"76-90" mapstructure:"76-90"`
	NinetyOneToOneHundredFive       MinuteDetails `json:"91-105" mapstructure:"91-105"`
	OneHundredSixToOneHundredTwenty MinuteDetails `json:"106-120" mapstructure:"106-120"`
}

type MinuteDetails struct {
	Total      *int     `json:"total"`
	Percentage *float32 `json:"percentage"`
}

type Biggest struct {
	Streak StreakDetails `json:"streak"`
	Wins   ScoreDetails  `json:"wins"`
	Loses  ScoreDetails  `json:"loses"`
	Goals  GoalScore     `json:"goals"`
}

type StreakDetails struct {
	Wins  int `json:"wins"`
	Draws int `json:"draws"`
	Loses int `json:"loses"`
}

type ScoreDetails struct {
	Home string `json:"home"`
	Away string `json:"away"`
}

type GoalScore struct {
	For     Score `json:"for"`
	Against Score `json:"against"`
}

type Score struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type CleanSheet struct {
	Home  int `json:"home"`
	Away  int `json:"away"`
	Total int `json:"total"`
}

type FailedToScore struct {
	Home  int `json:"home"`
	Away  int `json:"away"`
	Total int `json:"total"`
}

type Penalty struct {
	Scored PenaltyDetails `json:"scored"`
	Missed PenaltyDetails `json:"missed"`
	Total  int            `json:"total"`
}

type PenaltyDetails struct {
	Total      int    `json:"total"`
	Percentage string `json:"percentage"`
}

type Lineup struct {
	Formation string `json:"formation"`
	Played    int    `json:"played"`
}

type Cards struct {
	Yellow MinuteDistribution `json:"yellow"`
	Red    MinuteDistribution `json:"red"`
}
