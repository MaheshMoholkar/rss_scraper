package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MaheshMoholkar/rss_scraper/internal/database"
)

func (apiConfig *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID int32 `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON:%v", err))
		return
	}

	feedFollow, err := apiConfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Unable to create feed follow:%v", err))
		return
	}

	responseWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiConfig *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollow, err := apiConfig.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Unable to get feed follows:%v", err))
		return
	}

	responseWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feedFollow))
}
