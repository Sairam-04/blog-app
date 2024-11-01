package utils

import (
	"encoding/json"
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

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	return json.NewEncoder(w).Encode(payload)
}
