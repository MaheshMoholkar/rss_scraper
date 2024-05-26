package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MaheshMoholkar/rss_scraper/internal/auth"
	"github.com/MaheshMoholkar/rss_scraper/internal/database"
)

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON:%v", err))
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Unable to create user:%v", err))
		return
	}

	responseWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("Auth error:%v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKEY(r.Context(), apiKey)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Unable to fetch user: %v", err))
	}

	responseWithJSON(w, 200, databaseUserToUser(user))
}
