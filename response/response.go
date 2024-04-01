package response

type Response struct {
	Errors      interface{} `json:"errors,omitempty"`
	Results     int         `json:"results"`
	Paging      Paging      `json:"paging"`
	ResponseMap interface{} `json:"response"`
}

type Error struct {
	Time   string `json:"time"`
	Bug    string `json:"bug"`
	Report string `json:"report"`
}

type Paging struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}
