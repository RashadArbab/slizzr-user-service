package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Slizzr/slizzr-user-service/internal/database"
	"github.com/gorilla/mux"
)

type mongoGetUser interface {
	GetUser(id string) (*database.MongoUser, error)
}

func getUser(mongo mongoGetUser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			writeJsonError("missing user id", w)
			return
		}

		result, err := mongo.GetUser(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeJsonError(err.Error(), w)
			return
		}

		w.WriteHeader(http.StatusOK)
		encodeJsonResponse(result, w)
	}
}

type mongoGetMultipleUser interface {
	GetMultipleUser(ids []string) ([]*database.MongoUser, error)
}

type RequestBody struct {
	Ids []string `json:"ids"`
}

func getMultipleUser(mongo mongoGetMultipleUser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Checkpoint 1")
		var body RequestBody
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(body.Ids)

		result, err := mongo.GetMultipleUser(body.Ids)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeJsonError(err.Error(), w)
			return
		}

		w.WriteHeader(http.StatusOK)
		encodeJsonResponse(result, w)
	}
}
