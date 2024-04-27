Module Api Football Client
=================

A go lang SDK for [API Football](https://www.api-football.com/documentation-v3)

[![Go Coverage](https://github.com/syurchen93/api-football-client/wiki/coverage.svg)](https://raw.githack.com/wiki/syurchen93/api-football-client/coverage.html) [![Go Report Card](https://goreportcard.com/badge/github.com/syurchen93/api-football-client)](https://goreportcard.com/report/github.com/syurchen93/api-football-client)

Usage example:
```go
package main

import (
	"github.com/syurchen93/api-football-client/request/league"
	"github.com/syurchen93/api-football-client/client"
	"fmt"
)

func main() {
	apiClient := client.NewClient("your secret", client.RateLimiterSettings{})

	getCountriesRequest := league.Country{}
	getCountriesRequest.Code = "US"

	resp, err := apiClient.DoRequest(getCountriesRequest)
	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
```
The client is validating the request object for you so you don't waste precious API call on requests that will never give a result. Some undocumented edge cases are not inclided, feel free to contribute 😊

Since version 2.0 the rate limiter was added. Pass an empty settings object if you are using a free version of the API or customize it according to your plan.

The betting endpoints will not be implemented, because I don't support gambling. Fell free to fork and do those yourself