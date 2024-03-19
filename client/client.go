package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	//"os"
	"reflect"
	"strings"
	"time"

	"github.com/syurchen93/api-football-client/request"
	"github.com/syurchen93/api-football-client/request/team"
	"github.com/syurchen93/api-football-client/request/fixture"
	"github.com/syurchen93/api-football-client/request/league"
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/leagues"
	"github.com/syurchen93/api-football-client/response/misc"
	"github.com/syurchen93/api-football-client/response/fixtures"

	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
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
	//os.WriteFile("test/response/rounds-2021-cl.json", responseBody, 0644)
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

	var queryToAdd map[string]interface{}

	err = Decode(requestStruct, &queryToAdd)
	if err != nil {
		return "", err
	}

	for key, value := range queryToAdd {
		var valueToAdd string
		switch value := value.(type) {
			case int:
				valueToAdd = fmt.Sprintf("%d", value)
			case string:
				valueToAdd = value
		}
		curQuery.Add(key, valueToAdd)
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

	switch responseStruct.Errors.(type) {
		case []interface{}:
		case map[string]interface{}:
			if len(responseStruct.Errors.(map[string]interface{})) > 0 {
				return nil, fmt.Errorf("API returned errors: %v", responseStruct.Errors)
			}
	}

	var responseMap []interface{}
	switch responseStruct.ResponseMap.(type) {
		case []interface{}:
			responseMap = responseStruct.ResponseMap.([]interface{})
		case interface{}:
			responseMap = append(responseMap, responseStruct.ResponseMap)
	}

	endResponses := make([]response.ResponseInterface, 0)
	responseChan := make(chan response.ResponseInterface)
	errorChan := make(chan error)

	for _, responseMap := range responseMap {
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

	for range responseMap {
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
			ToTimeHookFunc(),
		),
		Result: result,
		WeaklyTypedInput: true,
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

		if f == reflect.TypeOf(team.Statistics{}) {
			response := map[string]interface{} {
				"season": data.(team.Statistics).Season,
				"team": data.(team.Statistics).Team,
				"league": data.(team.Statistics).League,
				"date": data.(team.Statistics).LimitDate.Format("2006-01-02"),
			}

			if (data.(team.Statistics).LimitDate.IsZero()) {
				delete(response, "date")
			}
			return response, nil
		}
		if f == reflect.TypeOf(fixture.Round{}) {
			response := map[string]interface{} {
				"league": data.(fixture.Round).League,
				"season": data.(fixture.Round).Season,
				"current": boolToString(data.(fixture.Round).Current),
			}

			return response, nil
		}
		if f == reflect.TypeOf(league.League{}) {
			response := map[string]interface{} {
				"id": data.(league.League).ID,
				"name": data.(league.League).Name,
				"country": data.(league.League).CountryName,
				"code": data.(league.League).CountryCode,
				"season": data.(league.League).Season,
				"team": data.(league.League).Team,
				"type": data.(league.League).Type,
				"current": boolToString(data.(league.League).Current),
				"search": data.(league.League).Search,
				"last": data.(league.League).Last,
			}
			if (data.(league.League).ID == 0) {
				delete(response, "id")
			}
			if (data.(league.League).Name == "") {
				delete(response, "name")
			}
			if (data.(league.League).CountryName == "") {
				delete(response, "country")
			}
			if (data.(league.League).CountryCode == "") {
				delete(response, "code")
			}
			if (data.(league.League).Season == 0) {
				delete(response, "season")
			}
			if (data.(league.League).Team == 0) {
				delete(response, "team")
			}
			if (data.(league.League).Type == "") {
				delete(response, "type")
			}
			if (data.(league.League).Search == "") {
				delete(response, "search")
			}
			if (data.(league.League).Last == 0) {
				delete(response, "last")
			}
			return response, nil
		}


		if t == reflect.TypeOf(leagues.SeasonYear{}) {
			return leagues.SeasonYear{Year: int(data.(float64))}, nil
		}
		if t == reflect.TypeOf(misc.Timezone{}) {
			return misc.Timezone{Value: data.(string)}, nil
		}
		if t == reflect.TypeOf(fixtures.Round{}) {
			return fixtures.Round{Name: data.(string)}, nil
		}

		if t == reflect.TypeOf(time.Time{}) {
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

		return data, nil
	}
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}