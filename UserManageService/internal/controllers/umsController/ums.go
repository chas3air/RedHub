package umscontroller

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"userManageService/internal/lib/logger/sl"

	redhub "github.com/chas3air/Domain/Redhub"
	"github.com/gorilla/mux"
)

type UserManageService struct {
	log    *slog.Logger
	client *http.Client
}

func New(logger *slog.Logger, client *http.Client) *UserManageService {
	return &UserManageService{
		log:    logger,
		client: client,
	}
}

func (ums *UserManageService) Get(w http.ResponseWriter, r *http.Request) {
	ums.log.Info("/users (GET) start process")

	// TODO: тут получение jwt токена из запроса

	req, err := http.NewRequest(http.MethodGet, "URL/users", nil)
	if err != nil {
		ums.log.Error("error of creating request:", sl.Err(err))
		http.Error(w, "Failed to create request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: тут вставка токена в запрос и прокидывание его дальше

	resp, err := ums.client.Do(req)
	if err != nil {
		ums.log.Error("error of sending request:", sl.Err(err))
		http.Error(w, "Failed to send request: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		ums.log.Error("status code is undetermined:", slog.Int("status_code", resp.StatusCode))
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ums.log.Error("error reading response body:", sl.Err(err))
		http.Error(w, "Failed to read response: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		ums.log.Error("error encoding response:", sl.Err(err))
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ums.log.Info("/users (GET) completed process")
}

func (ums *UserManageService) GetById(w http.ResponseWriter, r *http.Request) {
	ums.log.Info("/users/{id}/ (GET) start process...")

	id_s := mux.Vars(r)["id"]

	// TODO: тут получение jwt токена из запроса

	req, err := http.NewRequest(http.MethodGet, "URL/users/"+id_s+"/", nil)
	if err != nil {
		ums.log.Error("error of creating request:", sl.Err(err))
		http.Error(w, "Failed to create request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: тут вставка токена в запрос и прокидывание его дальше

	resp, err := ums.client.Do(req)
	if err != nil {
		ums.log.Error("error of sending request:", sl.Err(err))
		http.Error(w, "Failed to send request: "+err.Error(), http.StatusBadRequest)
		return
	}

	if resp.StatusCode >= 400 {
		ums.log.Error("status code is undetermined:", slog.Int("status_code", resp.StatusCode))
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ums.log.Error("error reading response body:", sl.Err(err))
		http.Error(w, "Failed to read response: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		ums.log.Error("error encoding response:", sl.Err(err))
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ums.log.Info("/users/id/ (GET) completed process")
}

func (ums *UserManageService) GetByLoginAndPassword(w http.ResponseWriter, r *http.Request) {
	ums.log.Info("/users/login (POST) start process...")

	loginAndPassword := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&loginAndPassword); err != nil {
		ums.log.Error("error of reading request body:", sl.Err(err))
		http.Error(w, "Failed to read request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: тут получение jwt токена из запроса

	req, err := http.NewRequest(http.MethodPost, "URL/users/login", r.Body)
	if err != nil {
		ums.log.Error("error of creating request:", sl.Err(err))
		http.Error(w, "Failed to create request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: тут вставка токена в запрос и прокидывание его дальше

	resp, err := ums.client.Do(req)
	if err != nil {
		ums.log.Error("error of sending request:", sl.Err(err))
		http.Error(w, "Failed to send request: "+err.Error(), http.StatusBadRequest)
		return
	}

	if resp.StatusCode >= 400 {
		ums.log.Error("status code is undetermined:", slog.Int("status_code", resp.StatusCode))
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ums.log.Error("error reading response body:", sl.Err(err))
		http.Error(w, "Failed to read response: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		ums.log.Error("error encoding response:", sl.Err(err))
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ums.log.Info("/users/login (POST) completed process")
}

func (ums *UserManageService) Insert(w http.ResponseWriter, r *http.Request) {
	ums.log.Info("/users (POST) start process...")

	var user_checker redhub.User
	if err := json.NewDecoder(r.Body).Decode(&user_checker); err != nil {
		ums.log.Error("error of reading request body:", sl.Err(err))
		http.Error(w, "Failed to read request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: тут получение jwt токена из запроса

	req, err := http.NewRequest(http.MethodPost, "URL/users/login", r.Body)
	if err != nil {
		ums.log.Error("error of creating request:", sl.Err(err))
		http.Error(w, "Failed to create request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: тут вставка токена в запрос и прокидывание его дальше

	resp, err := ums.client.Do(req)
	if err != nil {
		ums.log.Error("error of sending request:", sl.Err(err))
		http.Error(w, "Failed to send request: "+err.Error(), http.StatusBadRequest)
		return
	}

	if resp.StatusCode >= 400 {
		ums.log.Error("status code is undetermined:", slog.Int("status_code", resp.StatusCode))
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	ums.log.Info("/users/login (POST) completed process")
}

func (ums *UserManageService) Update(w http.ResponseWriter, r *http.Request) {
	ums.log.Info("/users/id (PUT) start process...")

	id_s := mux.Vars(r)["id"]

	var user_checker redhub.User
	if err := json.NewDecoder(r.Body).Decode(&user_checker); err != nil {
		ums.log.Error("error of reading request body:", sl.Err(err))
		http.Error(w, "Failed to read request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: тут получение jwt токена из запроса

	req, err := http.NewRequest(http.MethodPut, "URL/users/"+id_s, r.Body)
	if err != nil {
		ums.log.Error("error of creating request:", sl.Err(err))
		http.Error(w, "Failed to create request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: тут вставка токена в запрос и прокидывание его дальше

	resp, err := ums.client.Do(req)
	if err != nil {
		ums.log.Error("error of sending request:", sl.Err(err))
		http.Error(w, "Failed to send request: "+err.Error(), http.StatusBadRequest)
		return
	}

	if resp.StatusCode >= 400 {
		ums.log.Error("status code is undetermined:", slog.Int("status_code", resp.StatusCode))
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	ums.log.Info("/users/id (PUT) completed process")
}

func (ums *UserManageService) Delete(w http.ResponseWriter, r *http.Request) {
	ums.log.Info("/users/{id} (DELETE) start process...")

	id_s := mux.Vars(r)["id"]

	// TODO: тут получение jwt токена из запроса

	req, err := http.NewRequest(http.MethodDelete, "URL/users/"+id_s, nil)
	if err != nil {
		ums.log.Error("error of creating request:", sl.Err(err))
		http.Error(w, "Failed to create request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: тут вставка токена в запрос и прокидывание его дальше

	resp, err := ums.client.Do(req)
	if err != nil {
		ums.log.Error("error of sending request:", sl.Err(err))
		http.Error(w, "Failed to send request: "+err.Error(), http.StatusBadRequest)
		return
	}

	if resp.StatusCode >= 400 {
		ums.log.Error("status code is undetermined:", slog.Int("status_code", resp.StatusCode))
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ums.log.Error("error reading response body:", sl.Err(err))
		http.Error(w, "Failed to read response: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		ums.log.Error("error encoding response:", sl.Err(err))
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ums.log.Info("/users/id (DELETE) completed process")
}
