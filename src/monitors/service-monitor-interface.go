package monitors

import (
	"time"

	"gorm.io/gorm"
)

type MonitorResponse interface {
	UpdateDB()(error)
}

type BaseMonitorResponse struct {
	Id string
}

type BaseMonitorInformation struct {
	Model gorm.Model `gorm:"embedded"`
	Interval time.Duration
}

type ServiceMonitor interface {
	Monitor() (MonitorResponse, error)
	GetBaseInformation() (BaseMonitorInformation)
	GetType() (ServiceMonitorType)
	// IsOk(Response) (bool)
	// Notify() (error)
}

type ServiceMonitorType int

const (
	HttpMonitorType = iota
	PingMonitorType
)