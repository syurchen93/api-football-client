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
	apiClient := client.NewClient("your secret")

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
The client is validating the request object for you so you don't vaste precious API call on requests that will never give a result
### The project is open for pull requests 😊

### Contribution guide
- Implement request and response structures in proper folders according to API docs
- Add annotations for json (de)serialization
- Add validation annotation to the request according to API docs
- Do a request using a test file provided above
- Make sure that the response is mapped correctly
- Re-run the request storing request snapshot in a proper dir (just uncomment lines writing response body to a snapshot file in client.go)
- Add a test struct for your request in client_test.go
- Run tests ```go test -v ./client``` to generate a serialization result snapshot
