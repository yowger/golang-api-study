package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const port = ":8080"

/*
without tags: ID int

	output:
		{"ID":1,"Name":"Apple","Price":100}
*/

/*
with tags: ID int `json:"id"`

	output:
		{"id":1,"name":"Apple","price":100}
*/

/*
	to remove fields during serialization
		Price int `json:"-"`

		output:
			{"id":1,"name":"Apple"}
*/

/*
	to rename field during serialization
		Price int `json:"cost"`

		output:
			{"id":1,"name":"Apple","cost":100}
*/

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var items = []Item{
	{ID: 1, Name: "Laptop", Price: 1000},
	{ID: 2, Name: "Phone", Price: 500},
	{ID: 3, Name: "Tablet", Price: 300},
}

/*
	interface{} is equivalent to any in TS

	func respondWithJSON(response http.ResponseWriter, code int, payload interface{}) {

		cons:
			compiler won't check payload
*/

/*
	or use generics

	func respondWithJSON[T any](response http.ResponseWriter, code int, payload T) {

*/

func respondWithJSON[T any](response http.ResponseWriter, code int, payload T) {
	response.Header().Set("Content-Type", "application.json")
	response.WriteHeader(code)
	json.NewEncoder(response).Encode(payload)
}

func getItems(response http.ResponseWriter) {
	respondWithJSON(response, http.StatusOK, items)
}

func getItem(response http.ResponseWriter, request *http.Request) {
	idStr := strings.TrimPrefix(request.URL.Path, "/")

	fmt.Println("id: ", idStr)
	respondWithJSON(response, http.StatusOK, "test item")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getItems(response)
		default:
			message := map[string]string{"error": "Method not allowed"}
			respondWithJSON(response, http.StatusMethodNotAllowed, message)
		}
	})

	if serverError := http.ListenAndServe(port, mux); serverError != nil {
		log.Fatalf("server error: %v", serverError)
	}

}

/*
Yes, Go has always been capable of building REST APIs without third-party libraries because it includes a robust `net/http` package in its standard library. With the release of Go 1.22, you still use the same `net/http` package, but the language's simplicity and built-in functionality make it an excellent choice for building REST APIs.

Here’s a beginner example of creating a basic REST API in Go without any additional libraries.

---

### Example: A Simple REST API

This example demonstrates an API for managing a list of items with basic CRUD operations (Create, Read, Update, Delete).

#### Code:
```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Item represents a basic data structure
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Data store
var items = []Item{
	{ID: 1, Name: "Laptop", Price: 1000},
	{ID: 2, Name: "Phone", Price: 500},
}

// Helper function to write JSON responses
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

// Helper function to find an item by ID
func findItemByID(id int) (*Item, int) {
	for i, item := range items {
		if item.ID == id {
			return &item, i
		}
	}
	return nil, -1
}

// Handlers
func getItems(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/items/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid item ID"})
		return
	}

	item, _ := findItemByID(id)
	if item == nil {
		respondWithJSON(w, http.StatusNotFound, map[string]string{"error": "Item not found"})
		return
	}
	respondWithJSON(w, http.StatusOK, item)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}
	newItem.ID = len(items) + 1 // Simple ID generation
	items = append(items, newItem)
	respondWithJSON(w, http.StatusCreated, newItem)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/items/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid item ID"})
		return
	}

	_, index := findItemByID(id)
	if index == -1 {
		respondWithJSON(w, http.StatusNotFound, map[string]string{"error": "Item not found"})
		return
	}

	items = append(items[:index], items[index+1:]...)
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Item deleted"})
}

// Main function
func main() {
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getItems(w, r)
		case http.MethodPost:
			createItem(w, r)
		default:
			respondWithJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		}
	})

	mux.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getItem(w, r)
		case http.MethodDelete:
			deleteItem(w, r)
		default:
			respondWithJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		}
	})

	// Start server
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", mux)
}
```

---

### How It Works:
1. **Data Store**:
   - `items` is a slice that serves as an in-memory database.
2. **Routes**:
   - `/items`:
     - `GET`: Retrieves all items.
     - `POST`: Adds a new item.
   - `/items/{id}`:
     - `GET`: Retrieves a single item by ID.
     - `DELETE`: Deletes an item by ID.
3. **HTTP Methods**:
   - Routes differentiate operations using HTTP methods (e.g., `GET`, `POST`, `DELETE`).
4. **JSON Responses**:
   - `respondWithJSON` sends responses as JSON.

---

### Running the Server
1. Save the code in a file, e.g., `main.go`.
2. Run the server:
   ```sh
   go run main.go
   ```
3. Test the API:
   - **Get all items**:
     ```sh
     curl http://localhost:8080/items
     ```
   - **Create an item**:
     ```sh
     curl -X POST -H "Content-Type: application/json" -d '{"name":"Tablet","price":300}' http://localhost:8080/items
     ```
   - **Get a specific item**:
     ```sh
     curl http://localhost:8080/items/1
     ```
   - **Delete an item**:
     ```sh
     curl -X DELETE http://localhost:8080/items/2
     ```

---

### Why No Libraries?
- Go’s `net/http` package provides everything necessary for HTTP servers.
- For beginner APIs, you don’t need external libraries like `gin` or `echo`.

For more advanced routing or middleware, you can explore frameworks later as your project grows.
*/
