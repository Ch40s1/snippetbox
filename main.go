package main

import (
	"log"
	"net/http"
)

// define a home function which writes a byte slice containing
// hello from snippetbox as the responce body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func main() {
	// use http.NewServeMux function to initialize a new servemux
	// then register the home function as the handler and / as url pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Print a lof to say the server is starting
	log.Print("stating server on http://localhost:8080")
	// use the ListenandServe function to start the server
	// it uses two arguments tcp netwerk address this case its "8080"
	// and handler in this case it is servemux.If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	// Note: We are using ServeMux (mux) to handle our routing.
	// ServeMux maps URL patterns to their corresponding handlers.
	// In this case, we have mapped the "/" URL pattern to our home function.
	// When http.ListenAndServe(":8080", mux) is called, it starts the server on port 8080 and uses mux to handle incoming requests.
	// Mux then directs requests to the appropriate handler based on the URL pattern.
	// For example, a request to "http://localhost:8080/" will be handled by the home function
	// because we registered "/" with home using mux.HandleFunc("/", home).
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
