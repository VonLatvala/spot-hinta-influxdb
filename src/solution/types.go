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

type TodayResponse []HourParams

func (todayResp TodayResponse) String() string {
	return strings.Join(
		Map(
			todayResp,
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
}
