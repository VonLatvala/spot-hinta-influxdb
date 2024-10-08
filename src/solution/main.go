package spothintainfluxdb

import (
	"log"
	"time"
)

func Init() (*log.Logger, RuntimeConfig) {
	logger := *log.New(log.Writer(), solutionName + " ", log.LstdFlags | log.Lmsgprefix)
	logger.Println(bootMessage)

	var runtimeConfig RuntimeConfig = constructRuntimeConfig(&logger)

	return &logger, runtimeConfig
}

func Run(logger *log.Logger, runtimeConfig RuntimeConfig) {
	logger.Println("Starting main loop")
	stop := make(chan bool)
	go func() {
		for{
			tick(logger, runtimeConfig)
			select {
			case <-time.After(time.Duration(time.Duration(runtimeConfig.execIntervalMinutes) * time.Minute)):
			case <-stop:
				return
			}
		}
	}()
	for true{
		time.Sleep(1)
	}
}

func tick(logger *log.Logger, runtimeConfig RuntimeConfig) {
	var influxClient = connectInflux(logger, runtimeConfig.influxDatabaseConfig)
	defer influxClient.Close()

	logger.Println("Querying upstream API for today's hourly parameters")
	today := queryToday(logger, runtimeConfig.upstreamApiConfig)

	logger.Printf("Inserting %d hourly parameters to InfluxDB database %s at %s:%d\n",
	len(today), runtimeConfig.influxDatabaseConfig.name,
	runtimeConfig.influxDatabaseConfig.host, runtimeConfig.influxDatabaseConfig.port)
	insertTodayInflux(logger, influxClient, runtimeConfig.influxDatabaseConfig, today)
	logger.Printf("Successfully inserted %d records\n", len(today))
}
