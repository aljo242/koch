package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/aljo242/koch/util/file_util"
	"github.com/rs/zerolog/log"
)

const (
	htmlDir      string = "html/"
	jsDir        string = "js/"
	cssDir       string = "css/"
	tsDir        string = "src/"
	imgDir       string = "img/"
	modelDir     string = "model/"
	miscFilesDir string = "files"
	rootDir      string = "/"
)

// ScriptsHandler takes a script name and
func ScriptsHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r != nil {
			filename := filepath.Base(r.URL.Path)
			log.Debug().Str("Handler", "ScriptsHandler").Str("Filename", filename).Msg("incoming request")
			if r.Method == http.MethodGet {
				dir := filepath.Join(file_util.OutputDir, jsDir)
				wantFile := filepath.Join(dir, filename)
				if _, err := os.Stat(wantFile); os.IsNotExist(err) {
					w.WriteHeader(http.StatusNotFound)
					log.Debug().Err(err).Str("Filename", wantFile).Msg("Error finding file")
					return
				}

				switch filepath.Ext(wantFile) {
				case ".js":
					w.Header().Set("Content-Type", "application/javascript; charset=UTF-8")
				case ".js.map":
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				}

				w.Header().Set("Cache-Control", "max-age="+strconv.FormatInt(int64(cacheMaxAge), 10))
				http.ServeFile(w, r, wantFile)

			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}
}

// CSSHandler takes a script name and
func CSSHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r != nil {

			filename := filepath.Base(r.URL.Path)
			log.Debug().Str("Handler", "CSSHandler").Str("Filename", filename).Msg("incoming request")

			if r.Method == http.MethodGet {
				dir := filepath.Join(file_util.OutputDir, cssDir)
				wantFile := filepath.Join(dir, filename)
				if _, err := os.Stat(wantFile); os.IsNotExist(err) {
					w.WriteHeader(http.StatusNotFound)
					log.Debug().Err(err).Str("Filename", wantFile).Msg("Error finding file")
					return
				}

				w.Header().Set("Content-Type", "text/css; charset=UTF-8")
				w.Header().Set("Cache-Control", "max-age="+strconv.FormatInt(int64(cacheMaxAge), 10))
				http.ServeFile(w, r, wantFile)

			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}
}

// HTMLHandler takes a script name and
func HTMLHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r != nil {
			filename := filepath.Base(r.URL.Path)
			log.Debug().Str("Handler", "HTMLHandler").Str("Filename", filename).Msg("incoming request")

			if r.Method == http.MethodGet {
				dir := filepath.Join(file_util.OutputDir, htmlDir)
				wantFile := filepath.Join(dir, filename)
				if _, err := os.Stat(wantFile); os.IsNotExist(err) {
					w.WriteHeader(http.StatusNotFound)
					log.Debug().Err(err).Str("Filename", wantFile).Msg("Error finding file")
					return
				}

				w.Header().Set("Content-Type", "text/html; charset=UTF-8")
				w.Header().Set("Cache-Control", "max-age="+strconv.FormatInt(int64(cacheMaxAge), 10))
				http.ServeFile(w, r, wantFile)

			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}
}

// TypeScriptHandler takes a script name and returns a HandleFunc
func TypeScriptHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r != nil {
			filename := filepath.Base(r.URL.Path)
			log.Debug().Str("Handler", "TypeScriptHandler").Str("Filename", filename).Msg("incoming request")

			if r.Method == http.MethodGet {
				dir := filepath.Join(file_util.OutputDir, tsDir)
				wantFile := filepath.Join(dir, filename)
				if _, err := os.Stat(wantFile); os.IsNotExist(err) {
					w.WriteHeader(http.StatusNotFound)
					log.Debug().Err(err).Str("Filename", wantFile).Msg("Error finding file")
					return
				}

				w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
				w.Header().Set("Cache-Control", "max-age="+strconv.FormatInt(int64(cacheMaxAge), 10))
				http.ServeFile(w, r, wantFile)

			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}
}

// ManifestHandler serves manifest.json
func ManifestHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r != nil {
			filename := filepath.Base(r.URL.Path)

			if r.Method == http.MethodGet {
				log.Debug().Str("Handler", "ManifestHandler").Str("Filename", filename).Msg("incoming request")
				dir := filepath.Join(file_util.OutputDir, rootDir)
				wantFile := filepath.Join(dir, filename)
				if _, err := os.Stat(wantFile); os.IsNotExist(err) {
					w.WriteHeader(http.StatusNotFound)
					log.Debug().Err(err).Str("Filename", wantFile).Msg("Error finding file")
					return
				}

				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.Header().Set("Cache-Control", "max-age="+strconv.FormatInt(int64(cacheMaxAge), 10))
				http.ServeFile(w, r, wantFile)

			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}
}

// ServiceWorkerHandler serves serviceWorker.js
func ServiceWorkerHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r != nil {
			filename := filepath.Base(r.URL.Path)

			if r.Method == http.MethodGet {
				log.Debug().Str("Handler", "ServiceWorkerHandler").Str("Filename", filename).Msg("incoming request")
				dir := filepath.Join(file_util.OutputDir, rootDir)
				wantFile := filepath.Join(dir, filename)
				if _, err := os.Stat(wantFile); os.IsNotExist(err) {
					w.WriteHeader(http.StatusNotFound)
					log.Debug().Err(err).Str("Filename", wantFile).Msg("Error finding file")
					return
				}
				switch filepath.Ext(wantFile) {
				case ".js":
					w.Header().Set("Content-Type", "application/javascript; charset=UTF-8")
				case ".js.map":
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				}
				w.Header().Set("Cache-Control", "max-age="+strconv.FormatInt(int64(cacheMaxAge), 10))
				http.ServeFile(w, r, wantFile)

			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}
}

// ImageHandler returns a HandleFunc to serve image files
func ImageHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r != nil {
			filename := filepath.Base(r.URL.Path)

			if r.Method == http.MethodGet {
				dir := filepath.Join(file_util.OutputDir, imgDir)
				wantFile := filepath.Join(dir, filename)
				log.Debug().Str("Handler", "ImageHandler").Str("Filename", wantFile).Msg("incoming request")

				if _, err := os.Stat(wantFile); os.IsNotExist(err) {
					w.WriteHeader(http.StatusNotFound)
					log.Debug().Err(err).Str("Filename", wantFile).Msg("Error finding file")
					return
				}

				switch filepath.Ext(wantFile) {
				case ".jpg", ".jpeg":
					w.Header().Set("Content-Type", "image/jpeg")
				case ".png":
					w.Header().Set("Content-Type", "image/png")
				case ".gif":
					w.Header().Set("Content-Type", "image/gif")
				case ".ico":
					w.Header().Set("Content-Type", "image/x-icon")
				}
				w.Header().Set("Cache-Control", "max-age="+strconv.FormatInt(int64(cacheMaxAge), 10))
				http.ServeFile(w, r, wantFile)

			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}
}

// ModelHandler returns a HandleFunc to serve model files
func ModelHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r != nil {
			filename := filepath.Base(r.URL.Path)

			if r.Method == http.MethodGet {
				dir := filepath.Join(file_util.OutputDir, modelDir)
				wantFile := filepath.Join(dir, filename)
				log.Debug().Str("Handler", "ModelHandler").Str("Filename", wantFile).Msg("incoming request")

				if _, err := os.Stat(wantFile); os.IsNotExist(err) {
					w.WriteHeader(http.StatusNotFound)
					log.Debug().Err(err).Str("Filename", wantFile).Msg("Error finding file")
					return
				}

				switch filepath.Ext(wantFile) {
				case ".dae":
					w.Header().Set("Content-Type", "model/dae")
				case ".obj":
					w.Header().Set("Content-Type", "model/obj")
				case ".gltf":
					w.Header().Set("Content-Type", "model/gltf")
				}
				w.Header().Set("Cache-Control", "max-age="+strconv.FormatInt(int64(cacheMaxAge), 10))
				http.ServeFile(w, r, wantFile)

			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}
}

// MiscFileHandler serves file requests
func MiscFileHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r != nil {
			filename := filepath.Base(r.URL.Path)

			if r.Method == http.MethodGet {
				dir := filepath.Join(file_util.OutputDir, miscFilesDir)
				wantFile := filepath.Join(dir, filename)
				log.Debug().Str("Handler", "MiscFileHandler").Str("requested file", filename).Msg("incoming request")

				if _, err := os.Stat(wantFile); os.IsNotExist(err) {
					w.WriteHeader(http.StatusNotFound)
					log.Debug().Err(err).Str("Filename", wantFile).Msg("Error finding file")
					return
				}

				if filepath.Ext(wantFile) == ".pdf" {
					w.Header().Set("Content-Type", "application/pdf")
				}
				w.Header().Set("Cache-Control", "max-age="+strconv.FormatInt(int64(cacheMaxAge), 10))
				http.ServeFile(w, r, wantFile)

			} else {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}
}
