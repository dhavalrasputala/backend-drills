package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	port := ":8080"
	mux.HandleFunc("/copy", CopyHandler)
	log.Printf("Server starting at %v", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}

func CopyHandler(w http.ResponseWriter, r *http.Request) {
	// Mistak 1. Use MultipartReader to stream, avoiding loading the whole file into RAM
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, "Invalid multipart form", http.StatusBadRequest)
		return
	}
	// 2. Loop through the parts of the multipart form
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, "Error reading form", http.StatusInternalServerError)
			return
		}
		// 3. If this part is our file upload field (named "data")
		if part.FormName() == "data" {
			tempfile, err := os.Create("tempfile.csv")
			if err != nil {
				http.Error(w, "Unable to create file on disk", http.StatusInternalServerError)
				return
			}
			defer tempfile.Close()
			// 4. io.Copy now streams directly from the network socket to the disk
			// It only uses a tiny 32KB buffer in memory, regardless of file size!
			_, err = io.Copy(tempfile, part)
			if err != nil {
				http.Error(w, "Unable to copy content", http.StatusInternalServerError)
				return
			}
			w.Write([]byte("File streamed successfully!"))
			return
		}
	}
	http.Error(w, "File field 'data' not found", http.StatusBadRequest)
}
