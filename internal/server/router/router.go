package router

import (
	"context"
	tasksStorager "todoapi/internal/services/tasks/storager"
	usersStorager "todoapi/internal/services/users/storager"

	"github.com/go-chi/chi/v5"
)

const ContentTypeJSON = "application/json"

func NewRouter(ctx context.Context, usersRepo *usersStorager.UsersRepo, tasksRepo *tasksStorager.TasksRepo) chi.Router {
	r := chi.NewRouter()
	r.Post("/api/reg-user", RegUser(ctx, usersRepo))
	r.Route("/api", func(r chi.Router) {
		r.Use(CheckToken(ctx, usersRepo))
		r.Get("/test", TestRoute(ctx))
	})
	return r
}
