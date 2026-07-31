//day4:- Client Disconnects (r.Context()): Write a handler with a 10-second sleep. Print a message when the client disconnects using r.Context().Done().
// Notice: How the server knows to stop work if the user closes their browser.

package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	port := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloHandler)
	log.Printf("Server starting at %v", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-time.After(10 * time.Second):
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("HELLO WORLD!"))
	case <-ctx.Done():
		return
	}
}
