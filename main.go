
package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"golevel/app/controllers"
	"golevel/config"
	"golevel/routes"
)

func main() {
	config.LoadConfig()

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
