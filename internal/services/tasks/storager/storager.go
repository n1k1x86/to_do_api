package storager

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TasksRepo struct {
	pool *pgxpool.Pool
}

func (t *TasksRepo) CreateNewTask(ctx context.Context, task *Task) error {
	query := "insert into tasks(title,description,author,estimation,created_at,updated_at) values($1,$2,$3,$4,$5,$6)"
	_, err := t.pool.Exec(ctx, query, task.Title, task.Description, task.Author, task.Estimation, time.Now(), time.Now)
	if err != nil {
		log.Printf("error while inserting record into tasks: %v", err)
		return err
	}
	return nil
}

func (t *TasksRepo) DeleteTaskByID(ctx context.Context, id int64) error {
	query := "delete from tasks where id = $1"
	_, err := t.pool.Exec(ctx, query, id)
	if err != nil {
		log.Printf("error while deleting a record from tasks: %v", err)
		return err
	}
	return nil
}

func (t *TasksRepo) GetAllUsersTasks(ctx context.Context, userID int64) ([]*FullTask, error) {
	result := make([]*FullTask, 0)
	query := "select id, title, description, author, estimation, created_at, updated_at from tasks where author = $1"
	rows, err := t.pool.Query(ctx, query, userID)
	if err != nil {
		log.Printf("error while getting all users tasks: %v", err)
		return nil, err
	}
	for rows.Next() {
		var id int64
		var title string
		var description sql.NullString
		var estimation sql.NullInt64
		var createdAt time.Time
		var updatedAt time.Time
		err = rows.Scan(&id, &title, &description, &estimation, &createdAt, &updatedAt)
		if err != nil {
			log.Printf("error while scanning values: %v", err)
			return nil, err
		}
		result = append(result, &FullTask{
			ID:          id,
			Title:       title,
			Description: description.String,
			Estimation:  estimation.Int64,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		})
	}
	return result, nil
}

func (t *TasksRepo) GetAllTasks(ctx context.Context) ([]*FullTask, error) {
	result := make([]*FullTask, 0)
	query := "select id, title, description, author, estimation, created_at, updated_at from tasks"
	rows, err := t.pool.Query(ctx, query)
	if err != nil {
		log.Printf("error while getting all users tasks: %v", err)
		return nil, err
	}
	for rows.Next() {
		var id int64
		var title string
		var description sql.NullString
		var estimation sql.NullInt64
		var createdAt time.Time
		var updatedAt time.Time
		err = rows.Scan(&id, &title, &description, &estimation, &createdAt, &updatedAt)
		if err != nil {
			log.Printf("error while scanning values: %v", err)
			return nil, err
		}
		result = append(result, &FullTask{
			ID:          id,
			Title:       title,
			Description: description.String,
			Estimation:  estimation.Int64,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		})
	}
	return result, nil
}

func (t *TasksRepo) BuildUpdatePipline(task *Task) ([]interface{}, []string) {
	query := make([]string, 0)
	args := make([]interface{}, 0)
	argInd := 1
	if task.Title != "" {
		query = append(query, fmt.Sprintf("title=$%d", argInd))
		args = append(args, task.Title)
		argInd++
	}
	if task.Description != "" {
		query = append(query, fmt.Sprintf("description=$%d", argInd))
		args = append(args, task.Description)
		argInd++
	}
	if task.Author != 0 {
		query = append(query, fmt.Sprintf("author=$%d", argInd))
		args = append(args, task.Author)
		argInd++
	}
	if task.Estimation != 0 {
		query = append(query, fmt.Sprintf("estimation=$%d", argInd))
		args = append(args, task.Estimation)
		argInd++
	}
	return args, query
}

func (t *TasksRepo) UpdateTaskByID(ctx context.Context, task *Task, id int64) error {
	args, setQuery := t.BuildUpdatePipline(task)
	if len(args) == 0 {
		log.Println("nothing for updating in tasks")
		return nil
	}
	query := fmt.Sprintf("update tasks set %s, updated_at = NOW() where id = $%d", strings.Join(setQuery, ", "), len(args))
	_, err := t.pool.Exec(ctx, query, args, id)
	if err != nil {
		log.Printf("error while updating tasks with id = %d: %v", id, err)
		return err
	}
	return nil
}

func NewTasksRepo(pool *pgxpool.Pool) *TasksRepo {
	return &TasksRepo{pool: pool}
}
