package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/intogit/go-blog-backend/internal/database"
)

func HandleCreateFeedFollow(w http.ResponseWriter, r *http.Request) {
	user, errDetail, err := GetAuthenticatedUserByApiKey(w, r)
	if err != nil {
		log.Print("invalid user trying to follow a feed ", err)
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf(errDetail, " \n, %v ", err))
		return
	}
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Print("invalid request body - Couldn't decode parameters", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	feed_follow, err := ApiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		FeedFollowsID: uuid.New(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		UserID:        user.UserID,
		FeedID:        params.FeedID,
	})

	if err != nil {
		log.Print("unable to follow a feed by user", err)
		respondWithError(w, http.StatusInternalServerError, "unable to follow a feed by user")
		return
	}
	respondWithJSON(w, http.StatusOK, feed_follow)
}

func HandleGetAllFeedFollow(w http.ResponseWriter, r *http.Request) {
	user, errDetail, err := GetAuthenticatedUserByApiKey(w, r)
	if err != nil {
		log.Print("invalid user trying to get feed follow ", err)
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf(errDetail, " \n, %v ", err))
		return
	}
	feed_follow, err := ApiCfg.DB.GetAllFeedFollow(r.Context(), user.UserID)

	if err != nil {
		log.Print("unable to get followed feed for a user", err)
		respondWithError(w, http.StatusInternalServerError, "unable to get followed feed for a user")
		return
	}
	respondWithJSON(w, http.StatusOK, feed_follow)
}

// -- reason to use user_id along with feed_follows_id is that
// -- what if user B get feed_follows_id of A anyhow, then B will be able to delete the A feed feed_follows
// -- so, if we have check that A is only deleting eed_follows_id of A only.

func HandleDeleteFeedFollow(w http.ResponseWriter, r *http.Request) {
	user, errDetail, err := GetAuthenticatedUserByApiKey(w, r)
	if err != nil {
		log.Print("invalid user trying to delete a feed_follows ", err)
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf(errDetail, " \n, %v ", err))
		return
	}
	feed_follows_id_str := chi.URLParam(r, "feed_follow_id")
	feed_follows_id, err := uuid.Parse(feed_follows_id_str)
	if err != nil {
		log.Print("Couldn't parse feed follow id", err)
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't parse feed follow id %v ", err))
		return
	}

	err = ApiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		FeedFollowsID: feed_follows_id,
		UserID:        user.UserID,
	})

	if err != nil {
		log.Print("unable to (delete) unfollow a feed by users", err)
		respondWithError(w, http.StatusInternalServerError, "unable to get followed feed for a user")
		return
	}
	respondWithJSON(w, http.StatusOK, struct{}{})

	// TODO:
	// when we try to delete in sql and query condition doesn't holds true,
	// it deletes nothing from db, but returns success query execution
	// due to this we do not get to know if it really deleted the record...
	// so that's why we need to count the pre-delete records and post-delete records
	// and if they are equal, means no delete happend and we can send err to client
}
