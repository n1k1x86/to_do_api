package storager

import "time"

type FullUserInfo struct {
	ID        int64     `json:"id"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserDeleteRequest struct {
	Login string `json:"login"`
	Token string `json:"token"`
}
