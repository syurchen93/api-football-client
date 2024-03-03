package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/syurchen93/api-football-client/request"
	"github.com/syurchen93/api-football-client/response"

	"github.com/mitchellh/mapstructure"
	"github.com/go-playground/validator/v10"
)

var baseURL = "https://v3.football.api-sports.io/"
var apiHost = "v3.football.api-sports.io"
var validate *validator.Validate

type Client struct {
	apiKey string
	baseURL string
	apiHost string
	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	validate = validator.New(validator.WithRequiredStructEnabled())

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

func (c *Client) DoRequest(requestStruct request.RequestInterface) ([]response.ResponseInterface, error) {
	err := validate.Struct(requestStruct)
	if err != nil {
		return nil, err
	}

	requestUrlWithParams, err := c.prepareUrlWithParams(requestStruct)
	fmt.Println(requestUrlWithParams)

	if err != nil {
		return nil, err
	}

	httpRequest, err := http.NewRequest(
		"GET", 
		requestUrlWithParams, 
		nil,
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize http client: %v", err))
	}
	httpRequest.Header.Add("x-rapidapi-host", c.apiHost)
	httpRequest.Header.Add("x-rapidapi-key", c.apiKey)

	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	if httpResponse.StatusCode != 200 {
		return nil, err
	}

	defer httpResponse.Body.Close()
	responseBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		panic(fmt.Sprintf("Error reading API response: %v", err))
	}

	return mapResponseToCorrectStruct(responseBody, requestStruct)
}

func (c Client) prepareUrlWithParams(requestStruct request.RequestInterface) (string, error) {
	urlStruct, err := url.Parse(c.baseURL + requestStruct.GetEndpoint())
	curQuery := urlStruct.Query()

	if err != nil {
		return "", err
	}

	var queryToAdd map[string]string
	jsonTemp, jsonErr := json.Marshal(requestStruct)
	if jsonErr != nil {
		return "", jsonErr
	}

	err = json.Unmarshal(jsonTemp, &queryToAdd)
	if err != nil {
		return "", err
	}

	for key, value := range queryToAdd {
		curQuery.Add(key, value)
	}

	urlStruct.RawQuery = curQuery.Encode()

	return urlStruct.String(), nil
}

func mapResponseToCorrectStruct(
		responseBody []byte, 
		requestStruct request.RequestInterface,
	) ([]response.ResponseInterface, error) {
	responseStruct := response.Response{}

	jsonErr := json.Unmarshal(responseBody, &responseStruct)

	if jsonErr != nil {
		return nil, jsonErr
	}

	endResponses := make([]response.ResponseInterface, 0)
	for _, responseMap := range responseStruct.ResponseMap {
		emptyResponseStruct := requestStruct.GetResponseStruct()
		err := mapstructure.Decode(responseMap, &emptyResponseStruct)
		if err == nil {
			endResponses = append(endResponses, emptyResponseStruct)
		}
	}

	return endResponses, nil
}