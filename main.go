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
Day 2 Task:- Task: Using the standard http.ServeMux (Go 1.22+), create routes for GET /users/{id} and POST /users. Extract the id using r.PathValue("id").
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
