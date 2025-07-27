package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ilkeradgzl/shopstream/services/user-service/internal/model"
)

func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	user := model.User{ID: userID, Name: "Jane Doe", Email: "jane@example.com"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
