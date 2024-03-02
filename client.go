package client

import (
	"api-football-client/request"
	"api-football-client/response"
	"net/http"
	"encoding/json"
)

var baseURL = "https://v3.football.api-sports.io/"
var apiHost = "v3.football.api-sports.io"

type Client struct {
	apiKey string
	baseURL string
	apiHost string
	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		baseURL: baseURL,
		apiHost: apiHost,
		httpClient: &http.Client{},
	}
}

func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

func (c *Client) SetApiHost(apiHost string) {
	c.apiHost = apiHost
}

func (c *Client) doRequest(request request.RequestInterface) response.ResponseInterface {
	httpRequest, err := c.httpClient.NewRequest(
		"GET", 
		c.baseURL + request.getEndpoint(), 
		json.Marshal(request),
	)
	if err != nil {
		panic("Failed to initialize http client!", err)
	}
	httpRequest.Header.Add("x-rapidapi-host", c.apiHost)
	httpRequest.Header.Add("x-rapidapi-key", c.apiKey)

	httpResponse, err := httpClient.Do(request)
	if err != nil {
		panic("Error fetching API response", err)
	}
	httpResponse.Code != 200 {
		panic("API response code is not 200", err)
	}

	defer httpResponse.Body.Close()
	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		panic("Error reading API response", err)
	}

	fmt.Println(string(responseBody))

	return response.ResponseInterface{}
}