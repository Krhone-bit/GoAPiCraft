package utils

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(resp http.ResponseWriter, code int, message string) {
	respondWithJSON(resp, code, map[string]string{"error": message})
}

func respondWithJSON(resp http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(code)
	resp.Write(response)
}
