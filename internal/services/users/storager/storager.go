package storager

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserExist = errors.New("user is already exist")

type UsersRepo struct {
	pool *pgxpool.Pool
}

func (r *UsersRepo) BuildToken(login, hashedPass string) string {
	strToken := fmt.Sprintf("%s:%s", login, hashedPass)
	return base64.StdEncoding.EncodeToString([]byte(strToken))
}

func (r *UsersRepo) IsUserAuthenticated(ctx context.Context, login, password string) (string, bool, error) {
	query := "select login, password from users where login = $1;"
	row := r.pool.QueryRow(ctx, query, login)
	var pass sql.NullString
	var user sql.NullString
	err := row.Scan(&user, &pass)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", false, nil
		}
		return "", false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(pass.String), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", false, nil
		}
		return "", false, err
	}
	return user.String, true, nil
}

func (r *UsersRepo) IsUserAlreadyExist(ctx context.Context, u *User) (bool, error) {
	query := "select login from users where login = $1"
	row := r.pool.QueryRow(ctx, query, u.Login)
	var login sql.NullString
	err := row.Scan(&login)
	if err != nil && err != pgx.ErrNoRows {
		return false, err
	}
	if login.String != "" {
		return true, nil
	}
	return false, nil
}

func (r *UsersRepo) RegUser(ctx context.Context, u *User) (string, error) {
	ok, err := r.IsUserAlreadyExist(ctx, u)
	if err != nil {
		return "", err
	}
	if ok {
		return "", ErrUserExist
	}

	query := "insert into users(login, password, created_at, updated_at) values($1,$2,$3,$4)"

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
		return "", err
	}

	_, err = r.pool.Exec(ctx, query, u.Login, string(hashedPass), time.Now(), time.Now())
	if err != nil {
		log.Printf("error while executing query: %v", err)
		return "", err
	}
	token := r.BuildToken(u.Login, u.Password)
	return token, nil
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
