package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(HomeHandler)

	http.Handle("/", r)
	fmt.Println("Server is running on port 9090.  http://localhost:9090")
	http.ListenAndServe(":9090", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
