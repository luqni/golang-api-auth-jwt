package routers

import (
	"github.com/gorilla/mux"
)
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router := SetHelloRouter(router)
	router := SetAuthenticationRouter(router)
	return router
}
