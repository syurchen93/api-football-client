package players

import (
	"github.com/syurchen93/api-football-client/response/fixtures"
)

type Squad struct {
	Team    fixtures.TeamBasic `json:"team"`
	Players []PlayerBasic      `json:"players"`
}

type PlayerBasic struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Number   int    `json:"number"`
	Position string `json:"position"`
	Photo    string `json:"photo"`
}
