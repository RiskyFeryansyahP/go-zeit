package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ShowNama(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello, World!")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("/api").SubRouter()

	subrouter.HandleFunc("/hello", ShowNama).Methods(http.MethodGet)

	router.ServeHTTP(w, r)
}
