package fixtures

import (
	"github.com/syurchen93/api-football-client/common"
)

type Event struct {
	Time     Time                    `json:"time"`
	Team     Team                    `json:"team"`
	Player   Player                  `json:"player"`
	Assist   Player                  `json:"assist"`
	Type     common.EventType        `json:"type"`
	Detail   common.EventTypeDetails `json:"detail"`
	Comments string                  `json:"comments"`
}

type Time struct {
	Elapsed int `json:"elapsed"`
	Extra   int `json:"extra"`
}

type Player struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Photo string `mapstructure:"omitempty" json:"photo,omitempty"`
}

func (e Event) IsPenaltyShootout() bool {
	return e.Comments == common.CommentPenaltyShootout
}

func (e Event) IsFoul() bool {
	return e.Comments == common.CommentFoul
}
