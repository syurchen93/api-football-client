package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"io"

	"github.com/syurchen93/api-football-client/request"
	"github.com/syurchen93/api-football-client/response"
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

func (c *Client) doRequest(requestObject request.RequestInterface) response.ResponseInterface {
	requestBody, err := json.Marshal(requestObject)
	if err != nil {
		panic(fmt.Sprintf("Error serializing the request struct: %v", err))
	}

	httpRequest, err := http.NewRequest(
		"GET", 
		c.baseURL + requestObject.GetEndpoint(), 
		bytes.NewReader(requestBody),
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize http client: %v", err))
	}
	httpRequest.Header.Add("x-rapidapi-host", c.apiHost)
	httpRequest.Header.Add("x-rapidapi-key", c.apiKey)

	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		panic(fmt.Sprintf("Error fetching API response: %v", err))
	}
	if httpResponse.StatusCode != 200 {
		panic(fmt.Sprintf("API response code is not 200: %v", err))
	}

	defer httpResponse.Body.Close()
	responseBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		panic(fmt.Sprintf("Error reading API response: %v", err))
	}

	fmt.Println(string(responseBody))

	return responseBody
}