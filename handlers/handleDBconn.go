package handlers

import (
	"database/sql"

	"github.com/intogit/go-blog-backend/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

var ApiCfg apiConfig

func NewConn(dbConn *sql.DB) {
	ApiCfg = apiConfig{DB: database.New(dbConn)}
}

// conn is *sql.DB
// but sqlc is generalised and we need to convert the conn to *database.Queries
// we will use new() fxn of database package

// var ApiCfg *database.Queries
// func NewConn(dbConn *sql.DB) {
// 	ApiCfg = database.New(dbConn)
// }
