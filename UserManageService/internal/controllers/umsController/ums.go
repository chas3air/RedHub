package umscontroller

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"userManageService/internal/lib/logger/sl"

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

	resp, err := ums.client.Get("URL/users")
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
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		ums.log.Error("error encoding response:", sl.Err(err))
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ums.log.Info("/users (GET) completed process")
}

func (ums *UserManageService) GetById(w http.ResponseWriter, r *http.Request) {
	ums.log.Info("/users/{id} (GET) start process...")

	id_s := mux.Vars(r)["id"]

	resp, err := ums.client.Get("URL/users/" + id_s)
	if err != nil {
		// TODO: реализовать
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// TODO: реализовать
		return
	}
	defer resp.Body.Close()

	_ = body
}

func (ums *UserManageService) GetByLoginAndPassword(w http.ResponseWriter, r *http.Request) {}
func (ums *UserManageService) Update(w http.ResponseWriter, r *http.Request)                {}
func (ums *UserManageService) Delete(w http.ResponseWriter, r *http.Request)                {}
