package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Slizzr/slizzr-user-service/internal/config"
	"github.com/Slizzr/slizzr-user-service/internal/database"
)

type Server struct {
	Config  config.Config
	Context context.Context
	Mongo   *database.Mongo
}

func (s *Server) ListenAndServe(address string) error {
	return http.ListenAndServe(address, s.newRouter())
}

func encodeJsonResponse(v interface{}, w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = writeJsonError("internal server error", w)
	}
}
