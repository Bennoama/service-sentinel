package main

import (
	"log"
	"service-sentinel/db"
	"service-sentinel/monitoring"
	"service-sentinel/server"
)

func init () {
	err := db.Init("serviceSentinel")
	if err != nil {
		log.Fatal("ERROR! Failed to init Database, exiting", err)
		panic(err)
	}
	monitoring.StartMonitoring()
	server.Init()
}

func main() {		
	err := db.ShutDown()
	if err != nil {
		log.Fatal("ERROR! Failed to shut down database", err)
	}
}
