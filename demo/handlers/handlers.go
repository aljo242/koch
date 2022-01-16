package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/aljo242/koch/util/file_util"
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
		if r != nil {
			log.Debug().Str("Handler", "RedirectConstructionHandler").Str("Request URL", r.URL.Path).Msg("incoming request")
			http.Redirect(w, r, r.URL.Host+"/under-construction", http.StatusTemporaryRedirect)
		}
	}
}

func ConstructionHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r != nil {
			log.Debug().Str("Handler", "ConstructionHandler").Msg("incoming request")
			defer func() {
				dir := filepath.Join(file_util.OutputDir, htmlDir)
				wantFile := filepath.Join(dir, "construction.html")
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

		if r.Method == http.MethodGet && r != nil {
			log.Debug().Str("Handler", "HomeHandler").Msg("incoming request")
			defer func() {
				dir := filepath.Join(file_util.OutputDir, htmlDir)
				wantFile := filepath.Join(dir, "home.html")
				if _, err := os.Stat(wantFile); os.IsNotExist(err) {
					w.WriteHeader(http.StatusNotFound)
					log.Debug().Err(err).Str("Filename", wantFile).Str("BaseDir", dir).Msg("Error finding file")
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

// ResumeHomeHandler takes a script name and
func ResumeHomeHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		if r != nil {
			filename := filepath.Base(r.URL.Path)
			log.Debug().Str("Handler", "ResumeHomeHandler").Str("Filename", filename).Msg("incoming request")

			if r.Method == http.MethodGet {
				defer func() {
					dir := filepath.Join(file_util.OutputDir, htmlDir)
					wantFile := filepath.Join(dir, "resume.html")
					if _, err := os.Stat(wantFile); os.IsNotExist(err) {
						w.WriteHeader(http.StatusNotFound)
						log.Fatal().Err(err).Str("Filename", wantFile).Msg("Error finding file")
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
}

// TunesHomeHandler takes a script name and
func TunesHomeHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		if r != nil {
			filename := filepath.Base(r.URL.Path)
			log.Debug().Str("Handler", "ResumeHomeHandler").Str("Filename", filename).Msg("incoming request")

			if r.Method == http.MethodGet {
				defer func() {
					dir := filepath.Join(file_util.OutputDir, htmlDir)
					wantFile := filepath.Join(dir, "resume.html")
					if _, err := os.Stat(wantFile); os.IsNotExist(err) {
						w.WriteHeader(http.StatusNotFound)
						log.Fatal().Err(err).Str("Filename", wantFile).Msg("Error finding file")
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
}

// HallofArtHomeHandler takes a script name and
func HallofArtHomeHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		if r != nil {
			filename := filepath.Base(r.URL.Path)
			log.Debug().Str("Handler", "HallofArtHomeHandler").Str("Filename", filename).Msg("incoming request")

			if r.Method == http.MethodGet {
				defer func() {
					dir := filepath.Join(file_util.OutputDir, htmlDir)
					wantFile := filepath.Join(dir, "shadow.html")
					if _, err := os.Stat(wantFile); os.IsNotExist(err) {
						w.WriteHeader(http.StatusNotFound)
						log.Fatal().Err(err).Str("Filename", wantFile).Msg("Error finding file")
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
}
