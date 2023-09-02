package gooseneck

import (
	"net/http"
)

// TODO: add "github.com/julienschmidt/httprouter"

func MakeResponseFunc(status int, body string) func(w http.ResponseWriter) {
	return func(w http.ResponseWriter) {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(status)
		w.Write([]byte(body))
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	MakeResponseFunc(http.StatusOK, "{}")(w)
}
