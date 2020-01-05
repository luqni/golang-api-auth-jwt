package routersimport (
    "api_auth_with_jwt/controllers"
    "api_auth_with_jwt/core/authentication"
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
)func SetHelloRoutes(router *mux.Router) *mux.Router {
    router.Handle(
        "/test/hello",
        negroni.New(
            negroni.HandlerFunc(controllers.HelloController),
        )).Methods("GET")
    
    return router
}
