package samplers

import "time"

type MonitorResponse interface {
}

type BaseMonitorResponse struct {
	Id string
}

type BaseMonitorInformation struct {
	MonitorID string
	Interval time.Duration
}

type ServiceSampler interface {
	Sample() (MonitorResponse, error)
	// IsOk(Response) (bool)
	// Notify() (error)
}