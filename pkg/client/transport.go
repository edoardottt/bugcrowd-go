package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/edoardottt/bugcrowd-go/pkg/api"

	"github.com/avast/retry-go"
)

const (
	Attempts = 10
	Delay    = time.Millisecond * 100
	Error400 = 400
	Error500 = 500
)

type transport struct {
	underlying http.Transport
	username   string
	password   string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	var response *http.Response

	options := []retry.Option{
		retry.Delay(Delay),
		retry.DelayType(retry.BackOffDelay),
		retry.Attempts(Attempts),
		retry.LastErrorOnly(true),
	}

	if err := retry.Do(func() error {
		resp, err := t.underlying.RoundTrip(req)
		if err != nil {
			return err
		}

		if resp.StatusCode >= Error400 {
			defer func() { _ = resp.Body.Close() }()
			var apiError api.Error
			if err := json.NewDecoder(resp.Body).Decode(&apiError); err != nil || apiError.Status == 0 {
				return fmt.Errorf("server error: %v %v - %w", ErrStatus, resp.StatusCode, err)
			}
			if resp.StatusCode < Error500 {
				return retry.Unrecoverable(&apiError)
			}
			return &apiError
		}

		response = resp
		return nil
	}, options...); err != nil {
		return nil, err
	}

	return response, nil
}
