package main

import (
	"log"
	"net/http"
	"time"
)

type responsewriter struct {
	http.ResponseWriter
	StatusCode int
}

func main() {
	mux := http.NewServeMux()
	fhandler := AuthMiddleware(LogginMiddleware(mux))
	server := &http.Server{
		Addr:         ":8080",
		Handler:      fhandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Println("server Starting at :", server.Addr)
	err := server.ListenAndServe() //Use server.ListenAndServer when defined new server
	if err != nil {
		log.Fatal(err)
	}
}

func LogginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &responsewriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapped, r)
		log.Printf(
			"%v %d %s %s %v",
			r.RemoteAddr,
			wrapped.StatusCode,
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("x-api-key")
		if key == "" {
			http.Error(w, "Enter Valid key", http.StatusUnauthorized)
		}
		auth, ok := ClientFromKey(key)
		if !ok {
			http.Error(w, "Invalid key", http.StatusUnauthorized)
			return //mistake:- make sure you return back after
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("welcome " + auth + " you are authorized"))
		next.ServeHTTP(w, r)
	})
}

func ClientFromKey(key string) (string, bool) {
	if key == "secretkey123" {
		return "client-a", true
	}
	return "Unauthorized", false
}
