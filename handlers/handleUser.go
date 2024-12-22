package handlers

import (
	"encoding/json"
	// "errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/intogit/go-blog-backend/internal/auth"
	"github.com/intogit/go-blog-backend/internal/database"
)

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserName string `json:"user_name"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Print("invalid request body - Couldn't decode parameters", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	user, err := ApiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		UserID:    uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserName:  params.UserName,
	})
	if err != nil {
		log.Print("unable to add user ", err)
		respondWithError(w, http.StatusInternalServerError, "unable to add user")
		return
	}
	respondWithJSON(w, http.StatusOK, (user))

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(200)
	// json.NewEncoder(w).Encode(user)
	// log.Print("name added in db successfully", user)
}

func GetAuthenticatedUserByApiKey(w http.ResponseWriter, r *http.Request) (database.User, string, error) {
	apikey, err := auth.GetApiKey(r.Header)
	if err != nil {
		log.Println(err)
		errDetail := "Authentication error: "
		// err = errors.join(anothererr, err)
		return database.User{}, errDetail, err
	}
	user, err := ApiCfg.DB.GetUserByApiKey(r.Context(), apikey)
	if err != nil {
		log.Println(err)
		errDetail := "Database error: Could not get user "
		return database.User{}, errDetail, err
	}
	return user, "", nil
}

func HandleGetUserByApiKey(w http.ResponseWriter, r *http.Request) {
	user, errDetail, err := GetAuthenticatedUserByApiKey(w, r)
	if err != nil {
		log.Print(err)
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf(errDetail, " \n, %v ", err))
		return
	}
	// apikey, err := auth.GetApiKey(r.Header)
	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Autherization error: %v", err))
	// 	log.Println(err)
	// 	return
	// }
	// user, err := ApiCfg.DB.GetUserByApiKey(r.Context(), apikey)
	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not get user : %v", err))
	// 	log.Println(err)
	// 	return
	// }
	respondWithJSON(w, http.StatusOK, user)
}
