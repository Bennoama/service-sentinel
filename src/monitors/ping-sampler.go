package monitors

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/tatsushid/go-fastping"
)

type PingMonitorInformation struct {
	Address string
	Network string
}

type PingMonitor struct {
	BaseInfo BaseMonitorInformation `gorm:"embedded"`
	PingInfo PingMonitorInformation `gorm:"embedded"`
}

type PingResponse struct {
	baseMonitorRes BaseMonitorResponse `gorm:"embedded"`
	latency time.Duration
}

func (pingMonitor PingMonitor) GetBaseInformation() (BaseMonitorInformation) {
	return pingMonitor.BaseInfo
}

func (pingMonitor PingMonitor) Monitor() (MonitorResponse, error) {
	log.Println("Ping Monitor with id:", pingMonitor.BaseInfo.Model.ID)
	pinger := fastping.NewPinger()
	ipAddress, err := net.ResolveIPAddr(pingMonitor.PingInfo.Network, pingMonitor.PingInfo.Address)
	if err != nil {
		return PingResponse{}, err
	}
	pinger.AddIPAddr(ipAddress)
	response := PingResponse{}
	pinger.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		response.latency = (rtt)
	}
	err = pinger.Run()
	if err != nil {
		return PingResponse{}, err
	}
	return response, nil
}

func (pingRes PingResponse) UpdateDB() (error) {
	fmt.Fprintln(os.Stdout, []any{"Ping updating db -> %v", pingRes}...)
	return nil
}

func (PingMonitor) GetType () (ServiceMonitorType) {
	return PingMonitorType
}

