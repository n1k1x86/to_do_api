package router

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorMsg struct {
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func HandleBadRequest(w http.ResponseWriter, err error) {
	e := ErrorMsg{
		Status: http.StatusBadRequest,
		Detail: err.Error(),
	}
	data, err := json.Marshal(&e)
	if err != nil {
		log.Printf("error while handling bad request error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write(data)
}

func HandleInternalError(w http.ResponseWriter, err error) {
	e := ErrorMsg{
		Status: http.StatusInternalServerError,
		Detail: err.Error(),
	}
	data, err := json.Marshal(&e)
	if err != nil {
		log.Printf("error while handling bad request error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(data)
}
