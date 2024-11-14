package samplers

import "time"

type Response struct {
	ssl           string
	sslCertExpiry string
	data          []byte
	statusCode    int
	validity      int
	latency       time.Duration
}

type ServiceSampler interface {
	Sample() (Response, error)
	// IsOk(Response) (bool)
	// Notify() (error)
}