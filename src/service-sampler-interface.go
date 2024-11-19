package main

import "time"

type MonitorResponse interface {
	UpdateDB() error
}

type BaseMonitorInformation struct {
	monitorID string
	interval time.Duration
}

type ServiceSampler interface {
	Sample() (MonitorResponse, error)
	// IsOk(Response) (bool)
	// Notify() (error) // Goes to db with monitorID to get the notify details?
}