// all the routes would get taken care of here only
package auth

import "net/http"

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /auth/register", RegisterUser)

	return mux
}
