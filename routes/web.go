package routes

import (
	"golevel/app/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.HomeController{}.Index)
}
