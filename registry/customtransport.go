package registry

import (
	"net/http"
)

// ErrorTransport defines the data structure for returning errors from the round tripper.
type CustomTransport struct {
	Transport http.RoundTripper
	Headers   map[string]string
}

// RoundTrip defines the round tripper for the error transport.
func (t *CustomTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	if len(t.Headers) != 0 {
		for header, value := range t.Headers {
			request.Header.Add(header, value)
		}
	}

	resp, err := t.Transport.RoundTrip(request)

	return resp, err
}
