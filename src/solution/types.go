package spothintainfluxdb

import (
	"fmt"
	"strings"
)

type HourParams struct {
	Rank int
	DateTime string
	PriceNoTax float32
	PriceWithTax float32
}

func (hourParams HourParams) String() string {
	return fmt.Sprintf("%d\t%s\t%f\t%f",
		hourParams.Rank,
		hourParams.DateTime,
		hourParams.PriceNoTax,
		hourParams.PriceWithTax,
	)
}

type HourParamsResponse []HourParams
type TodayResponse = HourParamsResponse
// Alias as this is just a longer version
type TodayAndDayForwardResponse = HourParamsResponse

func (hourParamsResp HourParamsResponse) String() string {
	return strings.Join(
		Map(
			hourParamsResp,
			func(hourParams HourParams) string {
				return hourParams.String()
			},
		),
		",",
	)
}

type UpstreamApiConfig struct {
	proto string
	fqdn string
}

type InfluxDatabaseConfig struct {
	proto string
	host string
	port uint16
	name string
	username string
	password string
}

type RuntimeConfig struct {
	influxDatabaseConfig InfluxDatabaseConfig
	upstreamApiConfig UpstreamApiConfig
	execIntervalMinutes int
	includeTomorrow bool
}
