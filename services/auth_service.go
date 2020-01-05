package servicesimport (
 "api_auth_with_jwt/core/authentication"
 "api_auth_with_jwt/services/models"
 "encoding/json"
 jwt "github.com/dgrijalva/jwt-go"
 "net/http"
)func Login(requestUser *models.User) (int, []byte) {
 authBackend := authentication.InitJWTAuthenticationBackend()if authBackend.Authenticate(requestUser) {
  token, err := authBackend.GenerateToken(requestUser.UUID)
  if err != nil {
   return http.StatusInternalServerError, []byte("")
  } else {
   return http.StatusOK, response
  }
 }return http.StatusUnauthorized, []byte("")
}func RefreshToken(requestUser *models.User) []byte {
 authBackend := authentication.InitJWTAuthenticationBackend()
 token, err := authBackend.GenerateToken(requestUser.UUID)
 if err != nil {
  panic(err)
 }
 if err != nil {
  panic(err)
 }
 return response
}func Logout(req *http.Request) error {
 authBackend := authentication.InitJWTAuthenticationBackend()
 tokenRequest, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
  return authBackend.PublicKey, nil
 })
 if err != nil {
  return err
 }
 tokenString := req.Header.Get("Authorization")
 return authBackend.Logout(tokenString, tokenRequest)
}
