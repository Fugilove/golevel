package routes

import (
	"golevel/app/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.HomeController{}.Index)
}
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Ваш код обработки HTTP-запроса
}
