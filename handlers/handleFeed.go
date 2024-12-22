package handlers

import (
	"fmt"
	"net/http"

	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/intogit/go-blog-backend/internal/database"
)

func HandleCreateFeed(w http.ResponseWriter, r *http.Request) {

	user, errDetail, err := GetAuthenticatedUserByApiKey(w, r)
	if err != nil {
		log.Print("invalid user trying to create feed ", err)
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf(errDetail, " \n, %v ", err))
		return
	}
	type parameters struct {
		FeedName string `json:"feed_name"`
		FeedUrl  string `json:"feed_url"`
	}
	params := parameters{}
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Print("invalid request body - Couldn't decode parameters", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	feed, err := ApiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		FeedID:    uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedName:  params.FeedName,
		FeedUrl:   params.FeedUrl,
		UserID:    user.UserID,
	})
	if err != nil {
		log.Print("unable to add feed of user ", err)
		respondWithError(w, http.StatusInternalServerError, "unable to add feed of user")
		return
	}
	respondWithJSON(w, http.StatusOK, feed)
}
