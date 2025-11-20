package router

import (
	tasksStorager "todoapi/internal/services/tasks/storager"
	usersStorager "todoapi/internal/services/users/storager"

	"github.com/go-chi/chi/v5"
)

func NewRouter(usersRepo *usersStorager.UsersRepo, tasksRepo *tasksStorager.TasksRepo) chi.Router {
	r := chi.NewRouter()
	return r
}
