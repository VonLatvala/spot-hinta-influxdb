package spothintainfluxdb

const solutionName = "spot-hinta-influxdb"
const bootMessage = "Initializing spot-hinta influxdb integration"

const todayEndpoint = "today"
const todayAndDayForwardEndpoint = "todayAndDayForward"

const influxMeasurementRankField = "rank"
const influxMeasurementPriceNoTaxField = "price_no_tax"
const influxMeasurementPriceWithTaxField = "price_with_tax"
const influxMeasurementSpotPrice = "spot_price"
const influxMeasurementPrecision = "h"

const ENV_INFLUX_PROTO = "INFLUX_PROTO"
const ENV_INFLUX_HOST = "INFLUX_HOST"
const ENV_INFLUX_PORT = "INFLUX_PORT"
const ENV_INFLUX_DATABASE = "INFLUX_DATABASE"
const ENV_INFLUX_USERNAME = "INFLUX_USERNAME"
const ENV_INFLUX_PASSWORD = "INFLUX_PASSWORD"
const ENV_UPSTREAM_API_PROTO = "UPSTREAM_API_PROTO"
const ENV_UPSTREAM_API_FQDN = "UPSTREAM_API_FQDN"
const ENV_EXEC_INTERVAL_MINUTES = "EXEC_INTERVAL_MINUTES"
const ENV_INCLUDE_TOMORROW = "INCLUDE_TOMORROW"
