package samplers

import (
	"io"
	"net/http"
	"time"
)

type HttpSampler struct {
	url string
}

func (httpSampler HttpSampler) readFields(httpsRes *http.Response) (Response, error) {
	response := Response{}
	data, err := io.ReadAll(httpsRes.Body)
	if err != nil {
		return response, err
	}
	response.data = data
	response.statusCode = httpsRes.StatusCode
	return response, nil
}

func (httpSampler HttpSampler) Sample() (Response, error) {
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