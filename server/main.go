package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
