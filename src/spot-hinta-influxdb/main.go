package main

import "github.com/VonLatvala/spot-hinta-influxdb/spothintainfluxdb"


func main() {
	logger, runtimeConfig := spothintainfluxdb.Init()
	spothintainfluxdb.Run(logger, runtimeConfig)
}
