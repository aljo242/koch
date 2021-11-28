package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/aljo242/chef"
	"github.com/rs/zerolog/log"
)

// RedirectHome redirects to the {HOST}/home url with a 301 status
func RedirectHome() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Str("Handler", "RedirectHome").Str("Request URL", r.URL.Path).Msg("incoming request")
		http.Redirect(w, r, r.URL.Host+"/home", http.StatusPermanentRedirect)
	}
}

// RedirectConstructionHandler redirects to the {HOST}/under-construction url (construction handler) with a 307 (temporary moved) status
func RedirectConstructionHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Str("Handler", "RedirectConstructionHandler").Str("Request URL", r.URL.Path).Msg("incoming request")
		http.Redirect(w, r, r.URL.Host+"/under-construction", http.StatusTemporaryRedirect)
	}
}

func ConstructionHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			log.Debug().Str("Handler", "ConstructionHandler").Msg("incoming request")
			defer func() {
				wantFile := filepath.Join(htmlDir, "construction.html")
				if _, err := os.Stat(wantFile); os.IsNotExist(err) {
					w.WriteHeader(http.StatusNotFound)
					log.Debug().Err(err).Str("Filename", wantFile).Msg("Error finding file")
					return
				}

				w.Header().Set("Content-Type", "text/html; charset=UTF-8")
				w.Header().Set("Cache-Control", "max-age="+strconv.FormatInt(int64(cacheMaxAge), 10))
				http.ServeFile(w, r, wantFile)
			}()
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

// HomeHandler serves the home.html file
func HomeHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			log.Debug().Str("Handler", "HomeHandler").Msg("incoming request")
			defer func() {
				wantFile := filepath.Join(htmlDir, "home.html")
				if _, err := os.Stat(wantFile); os.IsNotExist(err) {
					w.WriteHeader(http.StatusNotFound)
					log.Debug().Err(err).Str("Filename", wantFile).Msg("Error finding file")
					return
				}

				w.Header().Set("Content-Type", "text/html; charset=UTF-8")
				w.Header().Set("Cache-Control", "max-age="+strconv.FormatInt(int64(cacheMaxAge), 10))
				http.ServeFile(w, r, wantFile)
			}()

			wantFile := cssDir + "home.css"
			chatFilepath, _ := filepath.Abs(wantFile)
			wantFile = jsDir + "app.js"
			jsFilepath, _ := filepath.Abs(wantFile)
			wantFile = imgDir + "1favicon.ico"
			faviconFilepath, _ := filepath.Abs(wantFile)
			wantFile = imgDir + "horse.jpg"
			backgroundImage, _ := filepath.Abs(wantFile)
			wantFile = modelDir + "kasa_obake.gltf"
			coolModel, _ := filepath.Abs(wantFile)
			err := chef.PushFiles(w, chatFilepath, jsFilepath, faviconFilepath, backgroundImage, coolModel)
			if err != nil {
				log.Error().Err(err).Msg("Error pushing files")
			}

		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

// ResumeHomeHandler takes a script name and
func ResumeHomeHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		filename := filepath.Base(r.URL.Path)
		log.Debug().Str("Handler", "ResumeHomeHandler").Str("Filename", filename).Msg("incoming request")

		if r.Method == http.MethodGet {
			defer func() {
				wantFile := filepath.Join(htmlDir, "resume.html")
				if _, err := os.Stat(wantFile); os.IsNotExist(err) {
					w.WriteHeader(http.StatusNotFound)
					log.Fatal().Err(err).Str("Filename", wantFile).Msg("Error finding file")
					return
				}

				w.Header().Set("Content-Type", "text/html; charset=UTF-8")
				w.Header().Set("Cache-Control", "max-age="+strconv.FormatInt(int64(cacheMaxAge), 10))
				http.ServeFile(w, r, wantFile)
			}()

			wantFile := cssDir + "resume.css"
			chatFilepath, _ := filepath.Abs(wantFile)
			wantFile = imgDir + "1favicon.ico"
			faviconFilepath, _ := filepath.Abs(wantFile)
			wantFile = imgDir + "cactus.jpg"
			backgroundImage, _ := filepath.Abs(wantFile)

			err := chef.PushFiles(w, chatFilepath, faviconFilepath, backgroundImage)
			if err != nil {
				log.Error().Err(err).Msg("Error pushing files")
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

// TunesHomeHandler takes a script name and
func TunesHomeHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		filename := filepath.Base(r.URL.Path)
		log.Debug().Str("Handler", "ResumeHomeHandler").Str("Filename", filename).Msg("incoming request")

		if r.Method == http.MethodGet {
			defer func() {
				wantFile := filepath.Join(htmlDir, "resume.html")
				if _, err := os.Stat(wantFile); os.IsNotExist(err) {
					w.WriteHeader(http.StatusNotFound)
					log.Fatal().Err(err).Str("Filename", wantFile).Msg("Error finding file")
					return
				}

				w.Header().Set("Content-Type", "text/html; charset=UTF-8")
				w.Header().Set("Cache-Control", "max-age="+strconv.FormatInt(int64(cacheMaxAge), 10))
				http.ServeFile(w, r, wantFile)
			}()

			wantFile := cssDir + "resume.css"
			chatFilepath, _ := filepath.Abs(wantFile)

			err := chef.PushFiles(w, chatFilepath)
			if err != nil {
				log.Error().Err(err).Msg("Error pushing files")
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

// HallofArtHomeHandler takes a script name and
func HallofArtHomeHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		filename := filepath.Base(r.URL.Path)
		log.Debug().Str("Handler", "HallofArtHomeHandler").Str("Filename", filename).Msg("incoming request")

		if r.Method == http.MethodGet {
			defer func() {
				wantFile := filepath.Join(htmlDir, "shadow.html")
				if _, err := os.Stat(wantFile); os.IsNotExist(err) {
					w.WriteHeader(http.StatusNotFound)
					log.Fatal().Err(err).Str("Filename", wantFile).Msg("Error finding file")
					return
				}

				w.Header().Set("Content-Type", "text/html; charset=UTF-8")
				w.Header().Set("Cache-Control", "max-age="+strconv.FormatInt(int64(cacheMaxAge), 10))
				http.ServeFile(w, r, wantFile)
			}()

			wantFile := cssDir + "home.css"
			chatFilepath, _ := filepath.Abs(wantFile)

			err := chef.PushFiles(w, chatFilepath)
			if err != nil {
				log.Error().Err(err).Msg("Error pushing files")
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
