package misc

type Trophy struct {
	LeagueName  string `mapstructure:"league"`
	CountryName string `mapstructure:"country"`
	Season      string `json:"season"`
	Place       string `json:"place"`
}
