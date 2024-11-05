package spothintainfluxdb

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func queryApi(logger *log.Logger, upstreamApiConfig UpstreamApiConfig, endpoint string) []byte  {
	var apiBaseUrl = fmt.Sprintf("%s://%s", upstreamApiConfig.proto, upstreamApiConfig.fqdn)

	apiRequest, err := http.Get(fmt.Sprintf("%s/%s", apiBaseUrl, endpoint))
	if err != nil {
		logger.Fatalf("Unable to GET '/%s': %s", endpoint, err.Error())
	}

	defer apiRequest.Body.Close()

	bodyBytes, err := io.ReadAll(apiRequest.Body)
	if err != nil {
		logger.Fatal(err)
	}

	if apiRequest.StatusCode != http.StatusOK {
		logger.Fatalf("HTTP %d: %s", apiRequest.StatusCode, string(bodyBytes))
	}

	return bodyBytes
}

func queryToday(logger *log.Logger, upstreamApiConfig UpstreamApiConfig) HourParamsResponse {
	var bodyBytes = queryApi(logger, upstreamApiConfig, todayEndpoint)

	var todayResponse TodayResponse

	err := json.Unmarshal(bodyBytes, &todayResponse)
	if err != nil {
		logger.Fatal(err)
	}

	return todayResponse
}

func queryTodayAndDayForward(logger *log.Logger, upstreamApiConfig UpstreamApiConfig) HourParamsResponse {
	var bodyBytes = queryApi(logger, upstreamApiConfig, todayAndDayForwardEndpoint)

	var todayAndDayForwardResponse TodayAndDayForwardResponse

	err := json.Unmarshal(bodyBytes, &todayAndDayForwardResponse)
	if err != nil {
		logger.Fatal(err)
	}

	return todayAndDayForwardResponse
}
