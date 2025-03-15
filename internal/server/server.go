package server

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hackermanpeter/caching-proxy/internal/client"
)

type httpHandler struct {
	origin string
}

func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// make request
	res, err := client.MakeRequest(r, h.origin)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// set headers
	headers := w.Header()
	for key, value := range res.Header {
		for _, v := range value {
			headers.Set(key, v)
		}
	}

	// respond
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("%v", err)
	}
	w.Write(data)

}

func Start(port int, origin string) {
	addr := fmt.Sprintf(":%v", port)

	s := http.Server{
		Addr:         addr,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler: httpHandler{
			origin,
		},
	}

	fmt.Printf("Listening on %v\n", addr)
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}

}
