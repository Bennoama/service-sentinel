package monitoring

import (
	"service-sentinel/db"
)

func StartMonitoring() error {
	monitorsList, err := db.GetAllMonitors()
	if err != nil {
		panic(err)
	}
	for _, currMonitor := range monitorsList {
		go currMonitor.Monitor()
	}
	return nil
}

