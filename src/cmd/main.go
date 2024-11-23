package main

import (
	"log"
	"service-sentinel/db"
)

func main() {			
	err := db.Init()
	if err != nil {
		log.Fatal("ERROR! Failed to init Database, exiting", err)
		panic(err)
	}
				
	monitors, err := db.GetAllMonitors()
	if err != nil {
		log.Fatal(err)
	}
	for _, monitor := range monitors {
		res, _ := monitor.Sample()
		log.Println(res)
	}

	err = db.ShutDown()
	if err != nil {
		log.Fatal("ERROR! Failed to shut down database", err)
	}
}
