package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hackermanpeter/caching-proxy/internal/cache"
)

func buildServer(port int, origin string) (*http.Server, error) {
	addr := fmt.Sprintf(":%v", port)
	return &http.Server{
		Addr:         addr,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler: httpHandler{
			origin,
		},
	}, nil

}

func setHeaders(w http.ResponseWriter, h http.Header) {
	headers := w.Header()
	for key, value := range h {
		for _, v := range value {
			headers.Set(key, v)
		}
	}
}

func getHeaders(url string, fileData map[string][]byte) (http.Header, error) {
	headerKey := cache.GetHeaderKey(url)
	var headers http.Header
	if err := json.Unmarshal(fileData[headerKey], &headers); err != nil {
		return nil, err
	}
	return headers, nil
}
