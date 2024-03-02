Module Api Football Client
=================

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