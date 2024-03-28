package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	//"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/syurchen93/api-football-client/common"
	"github.com/syurchen93/api-football-client/request"
	"github.com/syurchen93/api-football-client/request/fixture"
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/response/fixtures"
	"github.com/syurchen93/api-football-client/response/leagues"
	"github.com/syurchen93/api-football-client/response/misc"

	"github.com/syurchen93/api-football-client/parser"

	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
)

var baseURL = "https://v3.football.api-sports.io/"
var apiHost = "v3.football.api-sports.io"
var validate *validator.Validate
var timeFormatShort = "2006-01-02"

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
	//os.WriteFile("test/response/fixture-lineup-full.json", responseBody, 0644)
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

	var queryToAddTemp map[string]interface{}

	err = Decode(requestStruct, &queryToAddTemp)
	if err != nil {
		return "", err
	}
	queryToAdd := stringifyMapContent(queryToAddTemp)

	for key, value := range queryToAdd {
		if value != "" {
			curQuery.Add(key, value)
		}
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

		if t == reflect.TypeOf(fixtures.Lineup{}) {
			dataMap := data.(map[string]interface{})
			formation := dataMap["formation"]
			dataMap["formation"] = parser.ParseFormationStringIntoMap(formation.(string))
			if dataMap["startXI"] != nil {
				dataMap["startXI"] = preparePlayerMap(dataMap["startXI"])
			}
			if dataMap["substitutes"] != nil {
				dataMap["substitutes"] = preparePlayerMap(dataMap["substitutes"])
			}
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
						return time.Parse(timeFormatShort, data.(string))
					}
				case reflect.Float64:
					return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
				case reflect.Int64:
					return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
				default:
					return data, nil
				}
		}
		if f.String() == "string" && t.String() == "int" && strings.Contains(data.(string), "%") {
			return strconv.ParseInt(strings.TrimSuffix(data.(string), "%"), 10, 64)
		}

		if f.String() == "*time.Time" {
			return map[string]time.Time {
				"date": *data.(*time.Time),
			}, nil
		}
		if t == reflect.TypeOf(fixtures.Status{}) {
			elapsed := data.(map[string]interface{})["elapsed"]
			if nil != elapsed {
				elapsed = int(elapsed.(float64))
			} else {
				elapsed = 0
			}
			return fixtures.Status{
				Long: data.(map[string]interface{})["long"].(string),
				Value: common.FixtureStatus(data.(map[string]interface{})["short"].(string)),
				Elapsed: elapsed.(int),
			}, nil
		}

		return data, nil
	}
}

func stringifyMapContent(mapData map[string]interface{}) map[string]string {
	stringifiedMap := make(map[string]string)

	for key, value := range mapData {
		var stringValue string
		switch value := value.(type) {
			case bool:
				stringValue = boolToString(value)
			case string:
				stringValue = value
			case common.StatsType:
				stringValue = string(value)
			case common.EventType:
				stringValue = string(value)
			case fixture.LineupType:
				stringValue = string(value)
			case int:
				stringValue = fmt.Sprintf("%d", value)
			case time.Time:
				stringValue = value.Format(timeFormatShort)
			case []common.FixtureStatus:
				statusStrings := make([]string, len(value))
				for i, status := range value {
					statusStrings[i] = fmt.Sprintf("%v", status)
				}
				stringValue = strings.Join(statusStrings, "-")
			case []int:
				statusStrings := make([]string, len(value))
				for i, status := range value {
					statusStrings[i] = fmt.Sprintf("%v", status)
				}
				stringValue = strings.Join(statusStrings, "-")
			case map[string]interface{}:
				dateValue, ok := value["date"]
				if ok && !dateValue.(time.Time).IsZero() {
					stringValue = dateValue.(time.Time).Format(timeFormatShort)
				}
		}

		stringifiedMap[key] = stringValue
	}

	return stringifiedMap
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func preparePlayerMap(players interface{}) interface{} {
	playerSlice := make([]interface{}, 0)

	for _, player := range players.([]interface{}) {
		playerMap := player.(map[string]interface{})
		fixedPlayerMap := playerMap["player"].(map[string]interface{})
		grid := fixedPlayerMap["grid"]
		if grid != nil {
			fixedPlayerMap["grid"] = prepareGrid(fixedPlayerMap["grid"].(string))
		}

		playerSlice = append(playerSlice, fixedPlayerMap)
	}

	return playerSlice
}

func prepareGrid(gridString string) map[string]interface{} {
	gridMap := make(map[string]interface{})

	if gridString == "" {
		return gridMap
	}

	gridSlice := strings.Split(gridString, ":")
	if (len(gridSlice) != 2) {
		return gridMap
	}
	gridMap["row"] = gridSlice[0]
	gridMap["column"] = gridSlice[1]

	return gridMap
}