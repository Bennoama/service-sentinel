package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type HttpSampler struct {
	baseInfo BaseMonitorInformation
	httpInfo HttpMonitorInformation
}

type HttpMonitorInformation struct {
	url string
}

type HttpResponse struct {
	ssl           string
	sslCertExpiry string
	data          []byte
	statusCode    int
	validity      int
	latency       time.Duration
}

func (httpResponse HttpResponse) UpdateDB() (error) {
	fmt.Fprintln(os.Stdout, []any{"Http updating db -> %v", httpResponse}...)
	return nil
}

func (httpSampler HttpSampler) readFields(httpsRes *http.Response) (HttpResponse, error) {
	response := HttpResponse{}
	data, err := io.ReadAll(httpsRes.Body)
	if err != nil {
		return response, err
	}
	response.data = data
	response.statusCode = httpsRes.StatusCode
	return response, nil
}

func (httpSampler HttpSampler) Sample() (MonitorResponse, error) {
	startTime := time.Now()
	httpRes, err := http.Get(httpSampler.httpInfo.url)
	latency := time.Since(startTime)
	if err != nil {
		return HttpResponse{}, err
	}
	response, err := httpSampler.readFields(httpRes)
	if err != nil {
		return HttpResponse{}, err
	}
	response.latency = latency
	return response, nil
}