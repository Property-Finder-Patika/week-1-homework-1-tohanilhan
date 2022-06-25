package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Attach the route /greet/{name} to the function greet
	r.HandleFunc("/greet/{name}", greet).Methods(http.MethodGet)

	// Start the server on port 8001
	http.ListenAndServe(":8001", r)
}

// greet function prints the greeting with the given name
func greet(w http.ResponseWriter, r *http.Request) {
	// Get the name from the URL
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Get the name from the URL
	vars := mux.Vars(r)
	var name = vars["name"]

	// Print the greeting
	w.Write([]byte("Hello, " + name + "!"))
}
