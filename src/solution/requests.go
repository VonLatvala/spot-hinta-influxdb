package spothintainfluxdb

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func queryToday(logger *log.Logger, upstreamApiConfig UpstreamApiConfig) TodayResponse {
	var apiBaseUrl = fmt.Sprintf("%s://%s", upstreamApiConfig.proto, upstreamApiConfig.fqdn)

	apiRequest, err := http.Get(fmt.Sprintf("%s/%s", apiBaseUrl, todayEndpoint))
	if err != nil {
		logger.Fatalf("Unable to GET '/%s': %s", todayEndpoint, err.Error())
	}

	defer apiRequest.Body.Close()

	bodyBytes, err := io.ReadAll(apiRequest.Body)
	if err != nil {
		logger.Fatal(err)
	}

	if apiRequest.StatusCode != http.StatusOK {
		logger.Fatalf("HTTP %d: %s", apiRequest.StatusCode, string(bodyBytes))
	}

	var todayResponse TodayResponse

	err = json.Unmarshal(bodyBytes, &todayResponse)
	if err != nil {
		logger.Fatal(err)
	}

	return todayResponse
}
