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

	"github.com/syurchen93/api-football-client/common"
	"github.com/syurchen93/api-football-client/request"
	"github.com/syurchen93/api-football-client/request/fixture"
	"github.com/syurchen93/api-football-client/request/league"
	"github.com/syurchen93/api-football-client/request/misc"
	"github.com/syurchen93/api-football-client/request/player"
	"github.com/syurchen93/api-football-client/request/standings"
	"github.com/syurchen93/api-football-client/request/team"
	"github.com/syurchen93/api-football-client/response"
)

var responseFolder string = "../test/response/"
var resultFolder string = "../test/result/"

type testRequestStruct struct {
	RequestStruct         request.RequestInterface
	SnapshotName          string
	RequestUrlWithoutHost string
	ExpectError           bool
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
	var requestsToTest = []testRequestStruct{
		{
			RequestStruct:         league.Country{},
			SnapshotName:          "countries-full.json",
			RequestUrlWithoutHost: "/countries",
		},
		{
			RequestStruct:         league.Country{Code: "US"},
			SnapshotName:          "countries-one.json",
			RequestUrlWithoutHost: "/countries?code=US",
		},
		{
			RequestStruct:         league.League{Current: true, CountryCode: "DE"},
			SnapshotName:          "leagues-current-de.json",
			RequestUrlWithoutHost: "/leagues?code=DE&current=true",
		},
		{
			RequestStruct:         league.Season{},
			SnapshotName:          "seasons.json",
			RequestUrlWithoutHost: "/leagues/seasons",
		},
		{
			RequestStruct:         team.Team{Search: "Manchester"},
			SnapshotName:          "team-manchester.json",
			RequestUrlWithoutHost: "/teams?search=Manchester",
		},
		{
			RequestStruct:         team.Team{League: 39, Search: "Leipzig"},
			SnapshotName:          "team-leipzig-error.json",
			RequestUrlWithoutHost: "/teams?league=39&search=Leipzig",
			ExpectError:           true,
		},
		{
			RequestStruct:         team.Statistics{League: 78, Team: 173, Season: 2022},
			SnapshotName:          "team-stats-leipzig.json",
			RequestUrlWithoutHost: "/teams/statistics?league=78&season=2022&team=173",
		},
		{
			RequestStruct: team.Statistics{League: 78, Team: 173, Season: 2022,
				LimitDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)},
			SnapshotName:          "team-stats-leipzig-2022.json",
			RequestUrlWithoutHost: "/teams/statistics?date=2022-01-01&league=78&season=2022&team=173",
		},
		{
			RequestStruct:         misc.Timezone{},
			SnapshotName:          "timezones.json",
			RequestUrlWithoutHost: "/timezone",
		},
		{
			RequestStruct:         team.TeamSeason{Team: 33},
			SnapshotName:          "team-seasons-mu.json",
			RequestUrlWithoutHost: "/teams/seasons?team=33",
		},
		{
			RequestStruct:         team.Country{},
			SnapshotName:          "team-countries.json",
			RequestUrlWithoutHost: "/teams/countries",
		},
		{
			RequestStruct:         misc.Venue{Search: "Nou", Country: "Spain"},
			SnapshotName:          "venue-camp-nou-error.json",
			RequestUrlWithoutHost: "/venues?country=Spain&search=Nou",
			ExpectError:           true,
		},
		{
			RequestStruct:         misc.Venue{Search: "Nou"},
			SnapshotName:          "venue-camp-nou.json",
			RequestUrlWithoutHost: "/venues?search=Nou",
		},
		{
			RequestStruct:         standings.Standings{Season: 2022, League: 39},
			SnapshotName:          "standings-epl-2022.json",
			RequestUrlWithoutHost: "/standings?league=39&season=2022",
		},
		{
			RequestStruct:         fixture.Round{League: 2, Season: 2021, Current: false},
			SnapshotName:          "rounds-2021-cl.json",
			RequestUrlWithoutHost: "/fixtures/rounds?current=false&league=2&season=2021",
		},
		{
			RequestStruct: fixture.Fixture{
				League:   78,
				Statuses: []common.FixtureStatus{common.NotStarted, common.Postponed},
				Season:   2023,
			},
			SnapshotName:          "fixtures-bundes-ns-pst.json",
			RequestUrlWithoutHost: "/fixtures?league=78&season=2023&status=NS-PST",
		},
		{
			RequestStruct: fixture.Fixture{
				IDs: []int{1149523, 1149519},
			},
			SnapshotName:          "fixtures-cl-penalties.json",
			RequestUrlWithoutHost: "/fixtures?ids=1149523-1149519",
		},
		{
			RequestStruct: fixture.HeadToHead{
				H2H:  []int{33, 50},
				From: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				To:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
			},
			SnapshotName:          "head-to-head-mu-mc-2021.json",
			RequestUrlWithoutHost: "/fixtures/headtohead?from=2021-01-01&h2h=33-50&to=2023-12-31",
		},
		{
			RequestStruct: fixture.TeamStatistics{
				FixtureID: 653808,
				TeamID:    33,
				Type:      common.BallPossession,
			},
			SnapshotName:          "fixture-stats-ball-possession.json",
			RequestUrlWithoutHost: "/fixtures/statistics?fixture=653808&team=33&type=Ball+Possession",
		},
		{
			RequestStruct:         fixture.Event{FixtureID: 653808},
			SnapshotName:          "fixture-events.json",
			RequestUrlWithoutHost: "/fixtures/events?fixture=653808",
		},
		{
			RequestStruct:         fixture.Event{FixtureID: 653808, Type: common.Goal},
			SnapshotName:          "fixture-events-goal.json",
			RequestUrlWithoutHost: "/fixtures/events?fixture=653808&type=Goal",
		},
		{
			RequestStruct: fixture.Lineup{
				FixtureID: 653808,
				Type:      fixture.StartingXI,
			},
			SnapshotName:          "fixture-lineup-startXI.json",
			RequestUrlWithoutHost: "/fixtures/lineups?fixture=653808&type=startxi",
		},
		{
			RequestStruct: fixture.Lineup{
				FixtureID: 653808,
			},
			SnapshotName:          "fixture-lineup-full.json",
			RequestUrlWithoutHost: "/fixtures/lineups?fixture=653808",
		},
		{
			RequestStruct:         fixture.PlayerStatistics{FixtureID: 592872},
			SnapshotName:          "fixture-player-stats.json",
			RequestUrlWithoutHost: "/fixtures/players?fixture=592872",
		},
		{
			RequestStruct: misc.Injuries{
				Date:     time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC),
				LeagueID: 78,
			},
			SnapshotName:          "injuries-date-bundes.json",
			RequestUrlWithoutHost: "/injuries?date=2024-03-31&league=78",
		},
		{
			RequestStruct:         misc.Injuries{},
			SnapshotName:          "injuries-error.json",
			RequestUrlWithoutHost: "/injuries",
			ExpectError:           true,
		},
		{
			RequestStruct: fixture.Predictions{
				FixtureID: 1049124,
			},
			SnapshotName:          "fixture-predictions-bvb.json",
			RequestUrlWithoutHost: "/predictions?fixture=1049124",
		},
		{
			RequestStruct:         player.PlayerSeason{PlayerID: 2280},
			SnapshotName:          "player-season-azpi.json",
			RequestUrlWithoutHost: "/players/seasons?player=2280",
		},
		{
			RequestStruct:         player.Squad{TeamID: 173},
			SnapshotName:          "player-squad-Leipzig.json",
			RequestUrlWithoutHost: "/players/squads?team=173",
		},
		{
			RequestStruct:         player.PlayerInfo{LeagueID: 78, Season: 2022, Page: 3},
			SnapshotName:          "player-info-bundes-2022.json",
			RequestUrlWithoutHost: "/players?league=78&page=3&season=2022",
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
			t.Errorf("Expected %s\n\n\nGot %s\nIn %s\n", expectedResponseContent, actualResponseJson, requestToTest.SnapshotName)
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

func sortSliceByHash(slice []response.ResponseInterface) {
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
