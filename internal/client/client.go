package client

import (
	"fmt"
	"net/http"
	"time"
)

func new() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
	}
}

func MakeRequest(r *http.Request, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(r.Context(), r.Method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to form request: %v", err)
	}

	req.Header = r.Header
	req.Body = r.Body

	client := new()
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res, nil
}
