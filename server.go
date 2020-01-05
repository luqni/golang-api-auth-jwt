package mainimport (
 "api_auth_with_jwt/routers"
 "github.com/codegangsta/negroni"
 "net/http"
)func main() {
 router := routers.InitRoutes()
 n := negroni.Classic()
 n.UseHandler(router)
 http.ListenAndServe(":5000", n)
}
