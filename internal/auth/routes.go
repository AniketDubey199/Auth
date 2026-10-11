// all the routes would get taken care of here only
package auth

import "net/http"

func Routes(service *Service) *http.ServeMux {
	mux := http.NewServeMux()
	handler := Newhandler(service)

	mux.HandleFunc("POST /auth/register", handler.RegisterUser)
	mux.HandleFunc("POST /auth/login", handler.LoginUser)
	return mux
}
