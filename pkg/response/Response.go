package response

import (
	"encoding/json"
	"net/http"
)

// Success response
func JSON(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    data,
	}

	json.NewEncoder(w).Encode(res)
}

// Error response
func Error(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := map[string]interface{}{
		"status":  status,
		"message": message,
	}

	json.NewEncoder(w).Encode(res)
}
