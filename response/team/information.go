package team

type Information struct {
	Team  Team  `json:"team"`
	Venue Venue `json:"venue"`
}

type Team struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Code     string `json:"code,omitempty"`
	Country  string `json:"country,omitempty"`
	Founded  int    `json:"founded,omitempty"`
	National bool   `json:"national,omitempty"`
	Logo     string `json:"logo,omitempty"`
}

type Venue struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Address  string `json:"address,omitempty"`
	City     string `json:"city,omitempty"`
	Capacity int    `json:"capacity,omitempty"`
	Surface  string `json:"surface,omitempty"`
	Image    string `json:"image,omitempty"`
}
