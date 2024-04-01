package misc

import (
	"time"
)

type Coach struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Age         int    `json:"age"`
	Birth       Birth  `json:"birth"`
	Nationality string `json:"nationality"`
	Height      string `json:"height"`
	Weight      string `json:"weight"`
	Team        TeamBasic `json:"team"`
	Career      []Career `json:"career"`
}

type Birth struct {
	Date    time.Time `json:"date"`
	Place   string `json:"place"`
	Country string `json:"country"`
}

type TeamBasic struct {
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	Logo    string     `json:"logo"`
	Updated *time.Time `mapstructure:"update,omitempty"`
}

type Career struct {
    Team  TeamBasic `json:"team"`
	Start time.Time `json:"start"`
	End   *time.Time `json:"end,omitempty"`
}