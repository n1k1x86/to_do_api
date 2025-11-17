package storager

import "time"

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      int64  `json:"author"`
	Estimation  int64  `json:"estimation"`
}

type FullTask struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      int64     `json:"author"`
	Estimation  int64     `json:"estimation"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
