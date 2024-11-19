package main

import (
	"fmt"
	"os"
)

func GetSamplersFromDB() ([]ServiceSampler) {
	samplers := make([]ServiceSampler, 2)
	samplers[0] = HttpSampler{
		baseInfo: BaseMonitorInformation{},
		httpInfo: HttpMonitorInformation{url: "https://9eff769fd174490ea8a667d41c7bb7ff.api.mockbin.io/"},
	}
	samplers[1] = PingSampler{
		baseInfo: BaseMonitorInformation{},
		pingInfo: PingMonitorInformation{address: "127.0.0.1"},
	}
	return samplers
}

func main() {
	samplers := GetSamplersFromDB()
	for _, sampler := range samplers {
		fmt.Fprintln(os.Stdout, []any{"\n"}...)
		res, err := sampler.Sample()
		if err != nil {
			fmt.Fprintln(os.Stdout, []any{"ERR in sampler %v", sampler}...)
		} else {
			res.UpdateDB()
		}
	}
}
