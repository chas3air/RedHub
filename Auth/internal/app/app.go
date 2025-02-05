package app

import (
	"auth/internal/config"
	"auth/internal/controllers/authcontroller"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

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
	auth_controller := authcontroller.New(a.log, &http.Client{Timeout: a.cfg.Timeout})

	r := mux.NewRouter()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.HandleFunc("/login", auth_controller.Login)
	r.HandleFunc("/register", auth_controller.Register)

	a.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", a.cfg.Port),
		Handler: r,
	}

	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		if err := a.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
