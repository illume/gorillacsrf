package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"),
		csrf.Path("/api"),
		csrf.TrustedOrigins([]string{"http://localhost:8080"}),
	)

	api := r.PathPrefix("/api").Subrouter()
	api.Use(csrfMiddleware)

	// Set up CORS middleware
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-CSRF-Token"}),
	)
	api.Use(corsOptions)

	api.HandleFunc("/user/{id}", GetUser).Methods("GET")

	// handle a post request for creating a number
	api.HandleFunc("/number", CreateNumber).Methods("POST")

	http.ListenAndServe(":9090", r)
}

func CreateNumber(w http.ResponseWriter, r *http.Request) {
	// return {"number": 22}
	var number = struct {
		Number int `json:"number"`
	}{
		Number: 22,
	}
	w.Header().Set("X-CSRF-Token", csrf.Token(r))

	b, err := json.Marshal(number)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(b)

}

// Sets the token in the X-CSRF-Token header of the response and returns the user json.
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Authenticate the request, get the id from the route params,
	// and fetch the user from the DB, etc.
	var user = struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{
		ID:   "123",
		Name: "Alice",
	}

	// Get the token and pass it in the CSRF header. Our JSON-speaking client
	// or JavaScript framework can now read the header and return the token in
	// in its own "X-CSRF-Token" request header on the subsequent POST.
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
	b, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(b)
}
