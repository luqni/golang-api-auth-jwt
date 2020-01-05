package routers

import (
	"api_auth_with_jwt/controllers"
	"api_auth_with_jwt/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetAuthenticationRouter(router *mux.Router) *mux.Router {
	router.HandleFunc(
		"/token-auth",
		controllers.Login
	).Methods("POST")

	router.Handle(
		"/refresh-token-auth",
		negroni.New(
			negroni.HandleFunc(controllers.RefreshToken),
	)).Methods("GET")

	router.Handle(
		"/logout",
		negroni.New(
			negroni.HandleFunc(
				authentication.RequireTokenAuthentication
			),
			negroni.HandleFunc(controllers.Logout),
		)).Methods("GET")
	return router
}
