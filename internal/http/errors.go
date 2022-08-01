package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrUserNotFound           = errors.New("no user found")
	ErrAlreadyFriends         = errors.New("pq: duplicate key value violates unique constraint \"idx_joint\"")
	ErrRequestAlreadyAccepted = errors.New("this request has already been accepted")
)

type errorResponse struct {
	Msg string `json:"message"`
}

func writeJsonError(msg string, w http.ResponseWriter) error {
	jsonBody := errorResponse{Msg: msg}
	return json.NewEncoder(w).Encode(jsonBody)
}

func ErrAlreadyExists(id string) error {
	return errors.New(fmt.Sprintf("Neo4jError: Neo.ClientError.Schema.ConstraintValidationFailed (Node(11) already exists with label `Person` and property `id` = '%s')", id))
}
