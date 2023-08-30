package gooseneck

import (
	"net/http"
)

func MakeResponseFunc(status int, body string) func(w http.ResponseWriter) {
	return func(w http.ResponseWriter) {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(status)
		w.Write([]byte(body))
	}
}
