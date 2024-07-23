
package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"golevel/app/controllers"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.HomeController{}.Index)
}
