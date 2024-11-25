package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// init
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Logger)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	r.MethodNotAllowed(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed "))
	}))

	//main routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")

		if idParam == "404" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))

			return
		}

		w.Write([]byte(fmt.Sprintf("Hello, %s!", idParam)))
	})

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

	// server
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Couldn't listen on port 3000: ", err)
	}
}

/*
	one way to create sub routes

	todoRoutes := chi.NewRouter()

	todoRoutes.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Get all todos"))
	})

	r.Mount("/todo", todoRoutes)

*/
