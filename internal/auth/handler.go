// basic work of handler is to get data from the request and send data to the server all logics happends in service
package auth

import (
	"encoding/json"
	"net/http"

	"github.com/AniketDubey199/models"
)

type Handler struct {
	service *Service
}

func Newhandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req models.User
	var res models.UserReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Not valid structure")
		return
	}
	err := h.service.Register(r.Context(), req, &res)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&res)
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	var login models.UserLogin
	err := h.service.Login(r.Context(), &login, username, password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&login)
}
