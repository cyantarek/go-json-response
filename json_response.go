package jsonresp

import (
	"net/http"
	"encoding/json"
)

type Message struct {
	Message string `json:"message"`
}

func Resp(w http.ResponseWriter, msg string) {
	m := Message{Message:msg}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(m)
}
