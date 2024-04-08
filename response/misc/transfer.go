package misc

import "time"

type Transfer struct {
	Player    PlayerIdent      `json:"player"`
	Update    time.Time        `json:"update"`
	Transfers []TransferDetail `json:"transfers"`
}

type PlayerIdent struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TransferDetail struct {
	Date  time.Time     `json:"date"`
	Type  string        `json:"type"`
	Teams TransferTeams `json:"teams"`
}

type TransferTeams struct {
	In  Team `json:"in"`
	Out Team `json:"out"`
}
