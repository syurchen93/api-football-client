package client

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/syurchen93/api-football-client/request"
	"github.com/syurchen93/api-football-client/request/league"
	"github.com/syurchen93/api-football-client/request/misc"
	"github.com/syurchen93/api-football-client/request/team"
	"github.com/syurchen93/api-football-client/response"
	//"time"
)

var responseFolder string = "../test/response/"
var resultFolder string = "../test/result/"
type testRequestStruct struct {
	RequestStruct request.RequestInterface 
	SnapshotName string
	RequestUrlWithoutHost string
	ExpectError bool
}

func TestNewClient(t *testing.T) {
	apiKey := "test"
	client := NewClient(apiKey)
	if client.apiKey != apiKey {
		t.Errorf("Expected %s, got %s", apiKey, client.apiKey)
	}
}

func TestSetBaseURL(t *testing.T) {
	baseURL := "https://test.com"
	client := NewClient("test")
	client.SetBaseURL(baseURL)
	if client.baseURL != baseURL {
		t.Errorf("Expected %s, got %s", baseURL, client.baseURL)
	}
}

func TestSetApiHost(t *testing.T) {
	apiHost := "test.com"
	client := NewClient("test")
	client.SetApiHost(apiHost)
	if client.apiHost != apiHost {
		t.Errorf("Expected %s, got %s", apiHost, client.apiHost)
	}
}

func TestDoRequest(t *testing.T) {
	var requestsToTest = []testRequestStruct {
		{
			RequestStruct: league.Country{},
			SnapshotName: "countries-full.json",
			RequestUrlWithoutHost: "/countries",
		},
		{
			RequestStruct: league.Country{Code: "US"},
			SnapshotName: "countries-one.json",
			RequestUrlWithoutHost: "/countries?code=US",
		},
		{
			RequestStruct: league.League{Current: "true", CountryCode: "DE"},
			SnapshotName: "leagues-current-de.json",
			RequestUrlWithoutHost: "/leagues?code=DE&current=true",
		},
		{
			RequestStruct: league.Season{},
			SnapshotName: "seasons.json",
			RequestUrlWithoutHost: "/leagues/seasons",
		},
		{
			RequestStruct: team.Team{Search: "Manchester"},
			SnapshotName: "team-manchester.json",
			RequestUrlWithoutHost: "/teams?search=Manchester",
		},
		{
			RequestStruct: team.Team{League: 39, Search: "Leipzig"},
			SnapshotName: "team-leipzig-error.json",
			RequestUrlWithoutHost: "/teams?league=39&search=Leipzig",
			ExpectError: true,
		},
		{
			RequestStruct: team.Statistics{League: 78, Team: 173, Season: 2022},
			SnapshotName: "team-stats-leipzig.json",
			RequestUrlWithoutHost: "/teams/statistics?league=78&season=2022&team=173",
		},
		{
			RequestStruct: team.Statistics{League: 78, Team: 173, Season: 2022, 
				LimitDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)},
			SnapshotName: "team-stats-leipzig-2022.json",
			RequestUrlWithoutHost: "/teams/statistics?date=2022-01-01&league=78&season=2022&team=173",
		},
		{
			RequestStruct: misc.Timezone{},
			SnapshotName: "timezones.json",
			RequestUrlWithoutHost: "/timezone",
		},
		{
			RequestStruct: team.Season{Team: 33},
			SnapshotName: "team-seasons-mu.json",
			RequestUrlWithoutHost: "/teams/seasons?team=33",
		},
		{
			RequestStruct: team.Country{},
			SnapshotName: "team-countries.json",
			RequestUrlWithoutHost: "/teams/countries",
		},
	}

	for _, requestToTest := range requestsToTest {
		var actualResponseJson []byte
		resultFilePath := resultFolder + requestToTest.SnapshotName 
		actualResponse, err := mockRequest(t, requestToTest)
		if err != nil {
			if requestToTest.ExpectError {
				actualResponseJson = []byte(err.Error())
			} else {
				t.Fatalf("Error from DoRequest: %s", err)
			}
		} else {
			sortSliceByHash(actualResponse)
			actualResponseJson, err = json.Marshal(actualResponse)
			if err != nil {
				t.Fatalf("Error serializing response: %s", err)
			}
		}

		if _, err := os.Stat(resultFilePath); os.IsNotExist(err) {
			err := os.WriteFile(resultFilePath, actualResponseJson, 0644)
			if err != nil {
				t.Fatalf("Not able to write snapshot: %s", err)
			}
			t.Skip("Snapshot generated for", requestToTest.SnapshotName)
		}

		expectedResponseContent, err := os.ReadFile(resultFilePath)

		if err != nil {
			t.Fatalf("Error reading result file: %s", err)
		}

		if string(expectedResponseContent) != string(actualResponseJson) {
			t.Errorf("Expected %s\n\n\nGot %s", expectedResponseContent, actualResponseJson)
		}
	}
}

func mockRequest(t *testing.T, testRequesData testRequestStruct) ([]response.ResponseInterface, error) {
	responseContent, err := os.ReadFile(responseFolder + testRequesData.SnapshotName)
	if err != nil {
		t.Fatalf("Error reading file: %s", err)
	}
	responseContentString := string(responseContent)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, responseContentString)
	}))
	defer ts.Close()

	apiClient := NewClient("test")
	apiClient.baseURL = ts.URL + "/"
	apiClient.httpClient = ts.Client()

	actualUrl, err := apiClient.prepareUrlWithParams(testRequesData.RequestStruct)
	if err != nil {
		return nil, err
	}
	actualUrl = strings.Replace(actualUrl, ts.URL, "", 1)

	if testRequesData.RequestUrlWithoutHost != actualUrl {
		t.Fatalf("Expected url %s, got %s", testRequesData.RequestUrlWithoutHost, actualUrl)
	}

	return apiClient.DoRequest(testRequesData.RequestStruct)
}

func sortSliceByHash (slice []response.ResponseInterface) {
	getHash := func(r response.ResponseInterface) string {
		bytes, _ := json.Marshal(r)
		return fmt.Sprintf("%x", md5.Sum(bytes))
	}
	
	less := func(i, j int, slice []response.ResponseInterface) bool {
		return getHash(slice[i]) < getHash(slice[j])
	}
	
	sort.Slice(slice, func(i, j int) bool {
		return less(i, j, slice)
	})
}