package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/intogit/go-blog-backend/internal/database"
)

var apiCfg *database.Queries

func NewConn(dbConn *sql.DB) {
	apiCfg = database.New(dbConn)
}

// conn is *sql.DB
// but sqlc is generalised and we need to convert the conn to *database.Queries
// we will use new() fxn of database package

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Print("invalid request body", err)
	}
	user, err := apiCfg.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		log.Print("unable to add user ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
	log.Print("name added in db successfully", user)
}
