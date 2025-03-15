package server

import (
	"fmt"
	"io"
	"net/http"

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
	setHeaders(w, res)

	// respond
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("%v", err)
	}
	w.Write(data)

}

func Start(port int, origin string) {
	s := buildServer(port, origin)

	fmt.Printf("Listening on %v\n", port)
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}

}
