package client

import (
	"os"
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"sort"
	"strings"
	"crypto/md5"
	"encoding/json"
	"github.com/syurchen93/api-football-client/request/league"
	"github.com/syurchen93/api-football-client/response"
	"github.com/syurchen93/api-football-client/request"
)

var responseFolder string = "../test/response/"
var resultFolder string = "../test/result/"
type testRequestStruct struct {
	RequestStruct request.RequestInterface 
	SnapshotName string
	RequestUrlWithoutHost string
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
	}

	for _, requestToTest := range requestsToTest {
		resultFilePath := resultFolder + requestToTest.SnapshotName 
		actualResponse, err := mockRequest(t, requestToTest)
		if err != nil {
			t.Fatalf("Error from DoRequest: %s", err)
		}

		sortSliceByHash(actualResponse)
		actualResponseJson, err := json.Marshal(actualResponse)
		if err != nil {
			t.Fatalf("Error serializing response: %s", err)
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
		t.Errorf("Expected url %s, got %s", testRequesData.RequestUrlWithoutHost, actualUrl)
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