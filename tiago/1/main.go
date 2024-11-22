package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("Hello root"))

			return
		case "/users":
			w.Write([]byte("Hello users"))

			return
		default:
			w.Write([]byte("method not allowed"))

			return
		}

	}
}

func main() {
	api := &api{addr: ":8080"}

	srv := &http.Server{
		Addr:    api.addr,
		Handler: api,
	}

	srv.ListenAndServe()

	// if err := http.ListenAndServe(s.addr, s); err != nil {
	// 	log.Fatal("Error starting server: ", err)
	// }
}
