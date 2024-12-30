package client

import (
	"net/http"
	"time"
)

var base_url = ""

type HttpClient struct {
	http http.Client
}

func NewHttpClient(baseUrl string, timeout time.Duration) *HttpClient {
	base_url = baseUrl

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
