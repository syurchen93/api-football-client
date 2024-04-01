package misc

import (
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/misc"
)

type Timezone struct{}

func (t Timezone) GetEndpoint() string {
	return "timezone"
}

func (t Timezone) GetResponseStruct() response.ResponseInterface {
	return misc.Timezone{}
}
