package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Group(func(r chi.Router) {
		r.Get("/", helloWorldHandler)
	})
	r.Group(func(r chi.Router) {
		// r.Use(AuthMiddleware)
		r.Route("/todo", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Get all todos"))
			})

			r.Route("/{todoID}", func(r chi.Router) {
				r.Get("/", func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("Get single todo"))
				})
				r.Put("/", func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("Update todo"))
				})
				r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("Delete dodo"))
				})
			})
		})
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Could not start server:", err)
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func getAdminProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello admin!"))
}
