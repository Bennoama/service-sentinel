package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/tatsushid/go-fastping"
)

type Response struct { // go (haha) to service-sampler-interface.go file
	ssl           string
	sslCertExpiry string
	data          []byte
	statusCode    int
	validity      int
	latency       time.Duration
}

type ServiceSampler interface {  // go (haha) to service-sampler-interface.go file
	Sample() (Response, error)
	// IsOk(Response) (bool)
	// Notify() (error)
}


type HttpSampler struct { // go (haha) to http-sampler.go file
	url string
}

func (httpSampler HttpSampler) readFields(httpsRes *http.Response) (Response, error) { // go (haha) to http-sampler.go file
	response := Response{}
	data, err := io.ReadAll(httpsRes.Body)
	if err != nil {
		return response, err
	}
	response.data = data
	response.statusCode = httpsRes.StatusCode
	return response, nil
}

func (httpSampler HttpSampler) Sample() (Response, error) { // go (haha) to http-sampler.go file
	startTime := time.Now()
	httpRes, err := http.Get(httpSampler.url)
	latency := time.Since(startTime)
	if err != nil {
		return Response{}, err
	}
	response, err := httpSampler.readFields(httpRes)
	if err != nil {
		return Response{}, err
	}
	response.latency = latency
	return response, nil
}

type PingSampler struct { // go (haha) to ping-sampler.go file
	address string
	network string
}

func (pingSampler PingSampler) Sample() (Response, error) {// go (haha) to ping-sampler.go file
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

func (response Response) String() string { // go (haha) to service-sampler-interface.go file
	return fmt.Sprintf("\nssl: %v \nsslCertExpiry: %v\ndata: %v\nstatusCode: %v\nvalidity: %v \nlatency: %v[us]", response.ssl, response.sslCertExpiry, string(response.data), response.statusCode, response.validity, response.latency)
}

func main() {
	samplers := make([]ServiceSampler, 2)
	samplers[0] = HttpSampler{url: "https://9eff769fd174490ea8a667d41c7bb7ff.api.mockbin.io/"}
	samplers[1] = PingSampler{address: "127.0.0.1"}
	for _, sampler := range samplers {
		fmt.Fprintln(os.Stdout, []any{"\n"}...)
		res, err := sampler.Sample()
		if err != nil {
			fmt.Fprintln(os.Stdout, []any{"ERR in sampler %v", sampler}...)
		} else {
			var a []any = []any{reflect.TypeOf(sampler), res}
			fmt.Fprintln(os.Stdout, a...)
		}
	}
}
