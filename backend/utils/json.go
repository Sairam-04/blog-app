package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	type errResp struct {
		Error   string `json:"error"`
		Success bool   `json:"success"`
	}
	RespondWithJSON(w, code, errResp{
		Error:   msg,
		Success: false,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("failed to marshal JSON response %w", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content_type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
