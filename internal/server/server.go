package server

import (
	"context"
	"log"
	"net/http"
	"todoapi/internal/config"
	"todoapi/internal/server/router"
	tasksStorager "todoapi/internal/services/tasks/storager"
	usersStorager "todoapi/internal/services/users/storager"
)

type HTTPServer struct {
	server *http.Server
}

func (h *HTTPServer) Shutdown(ctx context.Context) error {
	err := h.server.Shutdown(ctx)
	if err != nil {
		return err
	}
	log.Println("server was closed successfully")
	return nil
}

func (h *HTTPServer) Run() error {
	log.Println("server is running")
	err := h.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func NewHTTPServer(ctx context.Context, cfg *config.Server, usersRepo *usersStorager.UsersRepo, tasksRepo *tasksStorager.TasksRepo) *HTTPServer {
	r := router.NewRouter(ctx, usersRepo, tasksRepo)
	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: r,
	}
	return &HTTPServer{
		server: server,
	}
}
