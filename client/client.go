package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	//"os"
	"time"
	"reflect"
	"strings"

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

	if err != nil {
		return nil, err
	}

	httpRequest, err := http.NewRequest(
		"GET", 
		requestUrlWithParams, 
		nil,
	)
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Add("x-rapidapi-host", c.apiHost)
	httpRequest.Header.Add("x-rapidapi-key", c.apiKey)

	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	if httpResponse.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status code: %d", httpResponse.StatusCode)
	}

	defer httpResponse.Body.Close()
	responseBody, err := io.ReadAll(httpResponse.Body)
	//os.WriteFile("test/response/leagues-current-de.json", responseBody, 0644)
	if err != nil {
		return nil, err
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
	responseChan := make(chan response.ResponseInterface)
	errorChan := make(chan error)

	for _, responseMap := range responseStruct.ResponseMap {
		go func(rm interface{}) {
			emptyResponseStruct := requestStruct.GetResponseStruct()
			err := Decode(rm, &emptyResponseStruct)
			if err != nil {
				errorChan <- err
			} else {
				responseChan <- emptyResponseStruct
			}
		}(responseMap)
	}

	for range responseStruct.ResponseMap {
		select {
		case response := <-responseChan:
			endResponses = append(endResponses, response)
		case err := <-errorChan:
			return nil, err
		}
	}

	return endResponses, nil
}


func Decode(input interface{}, result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			ToTimeHookFunc()),
		Result: result,
	})
	if err != nil {
		return err
	}

	if err := decoder.Decode(input); err != nil {
		return err
	}
	return err
}

func ToTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			if (strings.Contains(data.(string), "T")) {
				return time.Parse(time.RFC3339, data.(string))
			} else {
				return time.Parse("2006-01-02", data.(string))
			}
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
	}
}
