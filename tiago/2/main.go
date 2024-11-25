package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type api struct {
	addr string
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var users = []User{}

func (a *api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	newUsers := []User{{FirstName: "Tiago", LastName: "Silva"} 
	users = append(users, newUsers...)

	fmt.Println("Users: ", users)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

/*
	curl -X POST http://localhost:8080/users \
     -H "Content-Type: application/json" \
     -d '{"name": "John Doe"}'
*/

func (a *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var payloadUser User

	if err := json.NewDecoder(r.Body).Decode(&payloadUser); err != nil {
		http.Error(w, fmt.Sprintf("error decoding payload: %s", err), http.StatusInternalServerError)

		return
	}

	u := User{
		FirstName: payloadUser.FirstName,
		LastName:  payloadUser.LastName,
	}

	users = append(users, u)

	w.WriteHeader(http.StatusCreated)
}

func main() {
	api := &api{addr: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUserHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Error starting server: ", err)
	}

}
