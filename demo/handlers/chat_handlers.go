package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/aljo242/chef"
	"github.com/rs/zerolog/log"
)

// ChatHomeHandler is the route for the chat home where users can get assigned unique identifiers
func ChatHomeHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		// this page currently only serves html resources
		log.Debug().Str("Handler", "ChatHomeHandler").Msg("incoming request")

		if r.Method == http.MethodGet {
			defer func() {
				wantFile := filepath.Join(htmlDir, "chat.html")
				if _, err := os.Stat(wantFile); os.IsNotExist(err) {
					w.WriteHeader(http.StatusNotFound)
					log.Fatal().Err(err).Str("Filename", wantFile).Msg("Error finding file")
					return
				}

				w.Header().Set("Content-Type", "text/html; charset=UTF-8")
				w.Header().Set("Cache-Control", "max-age="+strconv.FormatInt(int64(cacheMaxAge), 10))
				http.ServeFile(w, r, wantFile)
			}()

			wantFile := cssDir + "chat.css"
			chatFilepath, _ := filepath.Abs(wantFile)
			wantFile = jsDir + "chat.js"
			jsFilepath, _ := filepath.Abs(wantFile)
			err := chef.PushFiles(w, chatFilepath, jsFilepath)
			if err != nil {
				log.Error().Err(err).Msg("Error pushing files")
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

// ChatSignUpHandler connects to the database and creates a new id for a chat user
func ChatSignUpHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// ChatSignInHandler connects to the database and signs a user in if they are in the database
func ChatSignInHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
