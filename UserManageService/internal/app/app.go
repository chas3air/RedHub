package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"userManageService/internal/config"
	"userManageService/internal/controllers/interfaces"
	umscontroller "userManageService/internal/controllers/umsController"

	"github.com/gorilla/mux"
)

type App struct {
	log *slog.Logger
	cfg *config.Config
	srv *http.Server
	wg  sync.WaitGroup
}

func New(logger *slog.Logger, config *config.Config) *App {
	return &App{
		log: logger,
		cfg: config,
	}
}

func (a *App) StartServer() error {
	var userManageController interfaces.UserManager = umscontroller.New(a.log, &http.Client{Timeout: a.cfg.Timeout})

	r := mux.NewRouter()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.HandleFunc("/users", userManageController.Get).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}/", userManageController.GetById).Methods(http.MethodGet)
	r.HandleFunc("/users/login", userManageController.GetByLoginAndPassword).Methods(http.MethodPost)
	r.HandleFunc("/users", userManageController.Insert).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", userManageController.Update).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", userManageController.Delete).Methods(http.MethodDelete)

	a.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", a.cfg.Port),
		Handler: r,
	}

	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		if err := a.srv.ListenAndServe(); err != nil {
			a.log.Error("error of starting server", "error", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	return a.Stop()
}

func (a *App) Stop() error {
	a.log.Info("Stoping server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("error while stoping server: %v", err)
	}

	a.wg.Wait()
	a.log.Info("Server is stoped")
	return nil
}
