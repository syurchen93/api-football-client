package response

type Response struct {
	Errors []Error `json:"errors"`
	Results int `json:"results"`
	Paging Paging `json:"paging"`
	ResponseMap []interface{} `json:"response"`
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