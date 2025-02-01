package authcontroller

import (
	"auth/internal/lib/logger/sl"
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"

	redhub "github.com/chas3air/Domain/Redhub"
)

type AuthController struct {
	log    *slog.Logger
	client *http.Client
}

// TODO: переделать на интерфейс
func New(logger *slog.Logger, client *http.Client) *AuthController {
	return &AuthController{
		log:    logger,
		client: client,
	}
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	ac.log.Info("/Login start process")

	//TODO: убрать "URL"
	// Request: логин и пароль пользователя в формате json
	// Response: пользователь в формате json, или ошибка если такого нет
	resp, err := ac.client.Post("URL/login", "application/json", r.Body)
	if err != nil {
		ac.log.Error("error of sending request:", sl.Err(err))
		http.Error(w, "Failed to send request: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		ac.log.Error("status code is undetermined:", slog.Int("status_code", resp.StatusCode))
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}

	respondeUser := redhub.User{}
	err = json.NewDecoder(resp.Body).Decode(&respondeUser)
	if err != nil {
		ac.log.Error("error reading response body:", sl.Err(err))
		http.Error(w, "Failed to read response: "+err.Error(), http.StatusBadRequest)
		return
	}

	// тут нужно провести генерацию jwt токена
	jwt_token := "token"
	// отоговая структура ответа:
	responseStruct := struct {
		User      redhub.User `json:"user"`
		Jwt_token string      `json:"jwt_token"`
	}{
		User:      respondeUser,
		Jwt_token: jwt_token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responseStruct); err != nil {
		ac.log.Error("error encoding response:", sl.Err(err))
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ac.log.Info("/login completed process")
}

func (ac *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	ac.log.Info("/Register start process")

	var user redhub.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		ac.log.Error("error reading request body:", sl.Err(err))
		http.Error(w, "Failed to read request: "+err.Error(), http.StatusBadRequest)
		return
	}

	bs, err := json.Marshal(user)
	if err != nil {
		ac.log.Error("error marshaling user:", sl.Err(err))
		http.Error(w, "Failed to prepare request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := ac.client.Post("URL/users", "application/json", bytes.NewBuffer(bs))
	if err != nil {
		ac.log.Error("error of sending request:", sl.Err(err))
		http.Error(w, "Failed to send request: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		ac.log.Error("status code is undetermined:", slog.Int("status_code", resp.StatusCode))
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}

	// TODO: тут генерация jwt токена
	jwt_token := ""
	respStruct := struct {
		JwtToken string `json:"jwt_token"`
	}{
		JwtToken: jwt_token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(respStruct); err != nil {
		ac.log.Error("error encoding response:", sl.Err(err))
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ac.log.Info("/register completed process")
}
