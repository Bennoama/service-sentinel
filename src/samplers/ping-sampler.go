package samplers

import (
	"net"
	"time"
	"github.com/tatsushid/go-fastping"
)

type PingMonitorInformation struct {
	Address string
	Network string
}

type PingSampler struct {
	BaseInfo BaseMonitorInformation `gorm:"embedded"`
	PingInfo PingMonitorInformation `gorm:"embedded"`
}

type PingResponse struct {
	baseMonitorRes BaseMonitorResponse `gorm:"embedded"`
	latency time.Duration
}


func (pingSampler PingSampler) Sample() (MonitorResponse, error) {
	pinger := fastping.NewPinger()
	ipAddress, err := net.ResolveIPAddr(pingSampler.PingInfo.Network, pingSampler.PingInfo.Address)
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
