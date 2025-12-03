package router

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	usersStorager "todoapi/internal/services/users/storager"
)

type RegUserResponse struct {
	Token string `json:"token"`
}

type contextKey string

const (
	userLogin contextKey = "login"
)

func CheckToken(ctx context.Context, usersRepo *usersStorager.UsersRepo) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			data, err := base64.StdEncoding.DecodeString(token)
			if err != nil {
				log.Printf("error while decoding auth token: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			tokenParams := strings.Split(string(data), ":")
			if len(tokenParams) != 2 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			login, ok, err := usersRepo.IsUserAuthenticated(ctx, tokenParams[0], tokenParams[1])
			if err != nil {
				log.Printf("error while checking user authentication: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			userCtx := context.WithValue(r.Context(), userLogin, login)

			next.ServeHTTP(w, r.WithContext(userCtx))
		})
	}
}

func RegUser(ctx context.Context, usersRepo *usersStorager.UsersRepo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("RegUser: error while reading body: %v", err)
			HandleInternalError(w, err)
			return
		}
		defer r.Body.Close()
		w.Header().Add("Content-Type", ContentTypeJSON)
		var u usersStorager.User
		err = json.Unmarshal(data, &u)
		if err != nil {
			log.Printf("RegUser: error while marshaling user obj: %v", err)
			HandleInternalError(w, err)
			return
		}
		token, err := usersRepo.RegUser(ctx, &u)
		if err != nil {
			log.Printf("error while reg user: %v", err)
			if errors.Is(err, usersStorager.ErrUserExist) {
				HandleBadRequest(w, usersStorager.ErrUserExist)
				return
			}
			HandleInternalError(w, err)
			return
		}
		resp := RegUserResponse{Token: token}
		body, err := json.Marshal(&resp)
		if err != nil {
			log.Printf("RegUser: error while reading body: %v", err)
			HandleInternalError(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func DeleteUser(ctx context.Context, usersRepo *usersStorager.UsersRepo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqCtx := r.Context()
		login := reqCtx.Value(userLogin).(string)

		err := usersRepo.DeleteUser(ctx, login)
		if err != nil {
			log.Printf("DeleteUser: error while deleting user: %v", err)
			HandleInternalError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func GetUserByLogin(ctx context.Context, usersRepo *usersStorager.UsersRepo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqCtx := r.Context()
		login := reqCtx.Value(userLogin).(string)
		user, err := usersRepo.GetUserByLogin(ctx, login)
		if err != nil {
			log.Printf("GetUserByLogin: error while getting user by login: %v", err)
			HandleInternalError(w, err)
			return
		}
		body, err := json.Marshal(&user)
		if err != nil {
			log.Printf("GetUserByLogin: error while marshaling user: %v", err)
			HandleInternalError(w, err)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func TestRoute(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		reqCtx := r.Context()
		login := reqCtx.Value(userLogin)
		w.Write([]byte(fmt.Sprintf("Hello, %s", login)))
	}
}
