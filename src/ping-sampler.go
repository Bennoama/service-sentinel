package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"github.com/tatsushid/go-fastping"
)

type PingMonitorInformation struct {
	address string
	network string
}

type PingSampler struct {
	baseInfo BaseMonitorInformation
	pingInfo PingMonitorInformation
}

type PingResponse struct {
	latency time.Duration
}

func (pingResponse PingResponse) UpdateDB() error {
	fmt.Fprintln(os.Stdout, []any{"Ping updating db -> latency in us", pingResponse.latency.Microseconds()}...)
	return nil
}

func (pingSampler PingSampler) Sample() (MonitorResponse, error) {
	pinger := fastping.NewPinger()
	ipAddress, err := net.ResolveIPAddr(pingSampler.pingInfo.network, pingSampler.pingInfo.address)
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
