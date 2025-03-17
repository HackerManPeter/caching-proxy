package server

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/hackermanpeter/caching-proxy/internal/cache"
	"github.com/hackermanpeter/caching-proxy/internal/client"
)

type httpHandler struct {
	origin string
}

func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// open file
	c, err := cache.Default()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	cacheDataPtr, err := c.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fileData := *cacheDataPtr

	// check if request is in cache
	url := fmt.Sprintf("%v%v", h.origin, r.URL)
	bodyKey := fmt.Sprintf("%v-[body]", url)
	if fileData[bodyKey] != nil {
		// if request is there return request

		headerKey := fmt.Sprintf("%v-[headers]", url)
		var headers http.Header
		json.Unmarshal(fileData[headerKey], &headers)

		setHeaders(w, headers)

		w.Write(fileData[bodyKey])
		slog.Info("X-CACHE: HIT")
		return
	}

	// else make request
	res, err := client.MakeRequest(r, h.origin)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
		return
	}

	data, err := c.Update(cacheDataPtr, url, res)
	if err != nil {
		fmt.Printf("Unable to update cache: %v", err)
	}

	// set headers
	headers := res.Header
	setHeaders(w, headers)

	// respond
	w.Write(data)
	slog.Info("X-CACHE: MISS")

}

func Start(port int, origin string) {
	s, err := buildServer(port, origin)
	if err != nil {
		fmt.Printf("Unable to start server: %v", err)
	}

	fmt.Printf("Listening on %v\n", port)
	if err = s.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}

	}

}
