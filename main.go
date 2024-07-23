package main

import (
	"golevel/config"
	"golevel/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
