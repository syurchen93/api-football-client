package request

import "github.com/syurchen93/api-football-client/response"

type RequestInterface interface {
	GetEndpoint() string
	GetResponseStruct() response.ResponseInterface
}