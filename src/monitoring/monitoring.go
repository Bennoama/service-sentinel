package monitoring

import (
	"log"
	"service-sentinel/db"
	"service-sentinel/monitors"
	"time"
)

func StartMonitoring() error {
	monitorsList, err := db.GetAllMonitors()
	if err != nil {
		panic(err)
	}
	for _, currMonitor := range monitorsList {
		go monitorWithInterval(currMonitor)
	}
	return nil
}

func monitorWithInterval(monitor monitors.ServiceMonitor) {
	for {
		var err error
		monitor.Monitor()
		monitorType := monitor.GetType()
		switch monitorType {
		case monitors.HttpMonitorType:
			monitor, err = db.GetMonitorByKey(monitor.GetBaseInformation().Model.ID, monitors.HttpMonitor{})
		case monitors.PingMonitorType:
			monitor, err = db.GetMonitorByKey(monitor.GetBaseInformation().Model.ID, monitors.PingMonitor{})
		default:
			log.Println("Unknown monitor type: ", monitorType)
		}
		if err != nil { // No monitor found in db
			break
		}
		time.Sleep(monitor.GetBaseInformation().Interval + 1 * time.Second)
	}
}