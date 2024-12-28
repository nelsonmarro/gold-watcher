package client

import (
	"net/http"
	"time"
)

var Base_url = "https://data-asg.goldprice.org/dbXRates/"

type HttpClient struct {
	http http.Client
}

func NewHttpClient(timeout time.Duration) *HttpClient {
	return &HttpClient{
		http: http.Client{
			Timeout: timeout,
		},
	}
}

func NewHttpClientWithTransport(timeout time.Duration, transport http.RoundTripper) *HttpClient {
	return &HttpClient{
		http: http.Client{
			Timeout:   timeout,
			Transport: transport,
		},
	}
}
