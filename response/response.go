package response

import "github.com/syurchen93/api-football-client/response/league"

type Response struct {
	Errors []Error `json:"errors"`
	Results int `json:"results"`
	Paging Paging `json:"paging"`
	Response []league.Country `json:"response"`
}

type Error struct {
	Time string `json:"time"`
	Bug string `json:"bug"`
	Report string `json:"report"`
}

type Paging struct {
	Current int `json:"current"`
	Total int `json:"total"`
}