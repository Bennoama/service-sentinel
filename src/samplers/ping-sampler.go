package samplers

import (
	"net"
	"time"

	"github.com/tatsushid/go-fastping"
)

type PingSampler struct {
	address string
	network string
}

func (pingSampler PingSampler) Sample() (Response, error) {
	pinger := fastping.NewPinger()
	ipAddress, err := net.ResolveIPAddr(pingSampler.network, pingSampler.address)
	if err != nil {
		return Response{}, err
	}
	pinger.AddIPAddr(ipAddress)
	response := Response{}
	pinger.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		response.latency = (rtt)
	}
	err = pinger.Run()
	if err != nil {
		return Response{}, err
	}

	return response, nil
}