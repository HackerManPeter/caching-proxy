package client

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func new() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
	}
}

func MakeRequest(r *http.Request, origin string) (*http.Response, error) {
	slog.Info("Request Received")

	req, err := http.NewRequestWithContext(r.Context(), r.Method, origin, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to form request: %v", err)
	}

	req.Header = r.Header
	req.Body = r.Body

	slog.Info("Response Sent")
	client := new()
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res, nil
}
