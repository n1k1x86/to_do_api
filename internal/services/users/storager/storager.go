package storager

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepo struct {
	pool *pgxpool.Pool
}

func (r *UsersRepo) InsertUser(ctx context.Context, u *User) error {
	query := "insert into users(login, password, created_at, updated_at) value($1,$2,$3,$4)"

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
		return err
	}

	_, err = r.pool.Exec(ctx, query, u.Login, hashedPass, time.Now(), time.Now())
	if err != nil {
		log.Printf("error while executing query: %v", err)
		return err
	}
	return nil
}

func (r *UsersRepo) DeleteUser(ctx context.Context, login string) error {
	query := "delete from users where login = $1"
	_, err := r.pool.Exec(ctx, query, login)
	if err != nil {
		log.Printf("error while executing query: %v", err)
		return err
	}
	return nil
}

func (r *UsersRepo) GetUserByLogin(ctx context.Context, login string) error {
	query := "select id, login, password, created_at, updated_at from users where login = $1"
	row := r.pool.QueryRow(ctx, query, login)
	var user FullUserInfo
	err := row.Scan(&user.ID, &user.Login, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Printf("error while scanning values into struct: %v", err)
		return err
	}
	return nil
}

func NewUsersRepo(pool *pgxpool.Pool) *UsersRepo {
	return &UsersRepo{pool: pool}
}
