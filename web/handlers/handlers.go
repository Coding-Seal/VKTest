package handlers

import (
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

func Handlers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /user/create", createUser())
	mux.HandleFunc("POST /quest/create", createQuest())
	mux.HandleFunc("POST /quest/{quest_id}/complete", markQuestComplete())
	mux.HandleFunc("GET /user/{user_id}/history", getUserHistory())
	return mux
}
func createUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New()
		slog.Debug("created user", slog.String("user_id", id.String()))
	}
}
func createQuest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New()
		slog.Debug("created quest", slog.String("quest_id", id.String()))
	}
}
func markQuestComplete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New()
		slog.Debug("marked quest complete", slog.String("quest_id", r.PathValue("quest_id")), slog.String("user_id", id.String()))
	}
}
func getUserHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Debug("retrieved user history", slog.String("user_id", r.PathValue("user_id")))
	}
}
