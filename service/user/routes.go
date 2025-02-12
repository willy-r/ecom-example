package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.HandleLogin).Methods(http.MethodPost)
	router.HandleFunc("/register", h.HandleRegister).Methods(http.MethodPost)
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// handle login
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	// handle register
}
