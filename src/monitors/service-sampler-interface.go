package monitors

import (
	"time"

	"gorm.io/gorm"
)

type MonitorResponse interface {
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
	// IsOk(Response) (bool)
	// Notify() (error)
}
