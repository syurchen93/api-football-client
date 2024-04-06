package misc

import "time"

type Transfer struct {
	Player    Player           `json:"player"`
	Update    time.Time        `json:"update"`
	Transfers []TransferDetail `json:"transfers"`
}

type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TransferDetail struct {
	Date  string        `json:"date"`
	Type  string        `json:"type"`
	Teams TransferTeams `json:"teams"`
}

type TransferTeams struct {
	In  Team `json:"in"`
	Out Team `json:"out"`
}
