package authcontroller

import (
	"log/slog"
	"net/http"
)

type AuthController struct {
	log *slog.Logger
}

func New(logger *slog.Logger) *AuthController {
	return &AuthController{
		log: logger,
	}
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	ac.log.Info("/Login is work")
}

func (ac *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	ac.log.Info("/Register is work")
}
