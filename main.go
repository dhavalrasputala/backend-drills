package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Custom Header"))
}
func main() {
	handler := MyHandler{}
	port := ":8080"
	fmt.Printf("Server Running on %s", port)
	err := http.ListenAndServe(port, handler)
	if err != nil {
		fmt.Printf("Server Failed :%v", err)
	}
}
