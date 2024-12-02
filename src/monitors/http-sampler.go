package monitors

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type HttpMonitor struct {
	BaseInfo BaseMonitorInformation `gorm:"embedded"`
	HttpInfo HttpMonitorInformation `gorm:"embedded"`
}

type HttpMonitorInformation struct {
	Url string
}

type HttpResponse struct {
	baseMonitorRes BaseMonitorResponse `gorm:"embedded"`
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

func (httpMonitor HttpMonitor) GetBaseInformation() (BaseMonitorInformation) {
	return httpMonitor.BaseInfo
}

func (httpMonitor HttpMonitor) readFields(httpsRes *http.Response) (HttpResponse, error) {
	response := HttpResponse{}
	data, err := io.ReadAll(httpsRes.Body)
	if err != nil {
		return response, err
	}
	response.data = data
	response.statusCode = httpsRes.StatusCode
	return response, nil
}

func (httpMonitor HttpMonitor) Monitor() (MonitorResponse, error) {
	log.Println("Http Monitor with id:", httpMonitor.BaseInfo.Model.ID)
	startTime := time.Now()
	httpRes, err := http.Get(httpMonitor.HttpInfo.Url)
	latency := time.Since(startTime)
	if err != nil {
		return HttpResponse{}, err
	}
	response, err := httpMonitor.readFields(httpRes)
	if err != nil {
		return HttpResponse{}, err
	}
	response.latency = latency
	return response, nil
}