package spothintainfluxdb

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkEnv(logger *log.Logger) {
	logger.Println("Checking environment")
	var requiredVariables = []string{
		ENV_INFLUX_PROTO,
		ENV_INFLUX_HOST,
		ENV_INFLUX_PORT,
		ENV_INFLUX_DATABASE,
		ENV_INFLUX_USERNAME,
		ENV_INFLUX_PASSWORD,
		ENV_UPSTREAM_API_PROTO,
		ENV_UPSTREAM_API_FQDN,
		ENV_EXEC_INTERVAL_MINUTES,
		ENV_INCLUDE_TOMORROW,
	}
	var missingFields []string;
	for _, variable := range requiredVariables {
		value := os.Getenv(variable)
		if len(strings.TrimSpace(value)) == 0 {
			missingFields = append(missingFields, variable)
		}
	}
	logger.Printf("Found %d missing environment variables", len(missingFields))
	sort.Strings(missingFields) // eww mutating sorter :(
	if len(missingFields) > 0 {
		logger.Fatalf("Missing environment variables: %s", strings.Join(missingFields, ", "))
	}
	logger.Println("Environment OK")
}

func constructRuntimeConfig(logger *log.Logger) RuntimeConfig {
	checkEnv(logger)
	logger.Println("Constructing runtime config")
	influxPort, err := strconv.Atoi(os.Getenv(ENV_INFLUX_PORT))
	if err != nil {
		logger.Fatal(err)
	}
	execIntervalMinutes, err := strconv.Atoi(os.Getenv(ENV_EXEC_INTERVAL_MINUTES))
	if err != nil {
		logger.Fatal(err)
	}
	includeTomorrow, err := strconv.ParseBool(os.Getenv(ENV_INCLUDE_TOMORROW))
	if err != nil {
		logger.Fatal(err)
	}
	var runtimeConfig = RuntimeConfig{
		influxDatabaseConfig: InfluxDatabaseConfig{
			proto: os.Getenv(ENV_INFLUX_PROTO),
			host: os.Getenv(ENV_INFLUX_HOST),
			port: uint16(influxPort),
			name: os.Getenv(ENV_INFLUX_DATABASE),
			username: os.Getenv(ENV_INFLUX_USERNAME),
			password: os.Getenv(ENV_INFLUX_PASSWORD),
		}, upstreamApiConfig: UpstreamApiConfig{
			proto: os.Getenv(ENV_UPSTREAM_API_PROTO),
			fqdn: os.Getenv(ENV_UPSTREAM_API_FQDN),
		},
		execIntervalMinutes: execIntervalMinutes,
		includeTomorrow: includeTomorrow,
	}
	logger.Println("Constructed runtime config")
	return runtimeConfig
}
