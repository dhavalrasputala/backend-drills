package main

/* Day 1 Task :-
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
*/
/*
Day 2:- Task: Using the standard http.ServeMux (Go 1.22+), create routes for GET /users/{id} and POST /users. Extract the id using r.PathValue("id").
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/users/{id}/{name}", AddUserHandler)
    port := ":8080"

    fmt.Printf("Server Starting at %v\n", port)
    if err := http.ListenAndServe(port, mux); err != nil {
        log.Fatal(err)
    }
}
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "invalid Method", http.StatusMethodNotAllowed)
        return
    }
    id := r.PathValue("id")
    name := r.PathValue("name")
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "ID: %s, Name: %s\n", id, name)
}
*/
/* Day 3:- Task: Create a struct with a field of type json.RawMessage. Unmarshal a JSON payload into it, then conditionally unmarshal that raw message into two different structs based on a "type" field.
package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    contents := []byte(`
    {
        "name":"value",
        "Id":"value2"
    }
`)
    mapStructure := make(map[string]string)
    _ = json.Unmarshal(contents, &mapStructure)

    fmt.Println("From Map:")
    fmt.Println(mapStructure["name"])
    fmt.Println(mapStructure["Id"])

    type Custom struct {
        Name string          // Must be exported ###Caution-struct field names should start With CAPITAL LETTER
        ID   json.RawMessage // Must be exported.
    }
    customStructure := &Custom{"Name:JSON","ID:404"}
    err := json.Unmarshal(contents, customStructure) ###Caution-customStructure is already an pointer to Custom struct no need to pass it as pointer again
    if err != nil {
        fmt.Println("Error:", err)
    }

    fmt.Println("\nFrom Struct:")
    fmt.Println("Name:", customStructure.Name)

    fmt.Println("Raw ID:", string(customStructure.ID))
}
*/
