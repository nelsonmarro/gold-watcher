package test

import (
	"net/http"
)

// fakeTransport simula el comportamiento de un transporte HTTP
type FakeTransport struct {
	Response *http.Response
	Err      error
}

func (f *FakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return f.Response, f.Err
}
