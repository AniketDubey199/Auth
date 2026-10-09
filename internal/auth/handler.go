// basic work of handler is to get data from the request and send data to the server all logics happends in service
package auth

import (
	"encoding/json"
	"net/http"

	"github.com/AniketDubey199/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req models.User

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Not valid structure")
		return
	}
	err := Register(req)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(&req)
}
