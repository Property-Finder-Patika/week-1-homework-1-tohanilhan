package main

// Call http://localhost:8080/ali
import (
	"fmt"
	"net/http"
)

func main() {
	// Create a new router
	http.HandleFunc("/", GreetingServer)

	// Start the server on port 8082
	http.ListenAndServe(":8082", nil)
}

// Run the program and open the browser at http://localhost:8080/tohan
// to see the result.
// $ go run greeting/web/greet.go
// result will be: Hello, tohan!
func GreetingServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
