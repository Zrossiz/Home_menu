package helpers

import (
	"encoding/json"
	"net/http"
)

func SendSuccess(rw http.ResponseWriter, result interface{}, statusCode int) error {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	if err := json.NewEncoder(rw).Encode(result); err != nil {
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return err
	}
	return nil
}
