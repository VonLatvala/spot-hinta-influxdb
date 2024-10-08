package spothintainfluxdb

import (
	"fmt"
	"log"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

func connectInflux(logger *log.Logger, influxDatabaseConfig InfluxDatabaseConfig) client.Client {
	influxClient, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: fmt.Sprintf("%s://%s:%d", influxDatabaseConfig.proto, influxDatabaseConfig.host, influxDatabaseConfig.port),
		Username: influxDatabaseConfig.username,
		Password: influxDatabaseConfig.password,
	})
	if err != nil {
		logger.Fatal(err)
	}
	return influxClient
}

func insertTodayInflux(logger *log.Logger, influxClient client.Client, influxDatabaseConfig InfluxDatabaseConfig, todayResponse TodayResponse) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: influxDatabaseConfig.name,
		Precision: influxMeasurementPrecision,
	})
	if err != nil {
		logger.Fatal(err)
	}
	for _, hourParams := range todayResponse {
		tags := map[string]string{}
		fields := map[string]interface{}{
			influxMeasurementRankField: hourParams.Rank,
			influxMeasurementPriceNoTaxField: hourParams.PriceNoTax,
			influxMeasurementPriceWithTaxField: hourParams.PriceWithTax,
		}
		datetime, err := time.Parse(time.RFC3339, hourParams.DateTime)
		if err != nil {
			logger.Fatal(err)
		}
		pt, err := client.NewPoint(influxMeasurementSpotPrice, tags, fields, datetime)
		if err != nil {
			logger.Fatal(err)
		}
		bp.AddPoint(pt)
	}
	if err:= influxClient.Write(bp); err != nil {
		logger.Fatal(err)
	}
}

