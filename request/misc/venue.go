package misc

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/misc"
)

type Venue struct {
	ID      int    `mapstructure:"id,omitempty"`
	Name    string `mapstructure:"name,omitempty"`
	City    string `mapstructure:"city,omitempty"`
	Country string `mapstructure:"country,omitempty"`
	Search  string `mapstructure:"search,omitempty" validate:"min=3,omitempty"`
}

func (v Venue) GetEndpoint() string {
	return "venues"
}

func (v Venue) GetResponseStruct() response.ResponseInterface {
	return misc.Venue{}
}
