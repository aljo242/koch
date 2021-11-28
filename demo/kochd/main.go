package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aljo242/koch/demo/handlers"
	"github.com/aljo242/koch/server"
	"github.com/aljo242/koch/template"
	"github.com/aljo242/koch/util/file_util"
	"github.com/aljo242/koch/util/ip_util"
	"github.com/aljo242/koch/x/chat"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	// DefaultConfigFile is the name of the user's config JSON file
	DefaultConfigFile string = "sample/sample_config.json"

	// TemplateBaseDir is where HTML template files are located to be
	// executed and copied to the res dir
	TemplateBaseDir string = "./web_res"

	// TemplateOutputDir is the directory all outputs of SetupTemplates will fall under
	TemplateOutputDir string = "./static"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "c", DefaultConfigFile, "Full path to JSON configuration file")
}

func setupLogger(cfg server.Config) {
	if cfg.DebugLog {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("log level is DEBUG")
	} else {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		log.Error().Msg("log level is ERROR")

	}
}

// SetupTemplates builds the template output directory, executes HTML templates,
// and copies all web resource files to the template output directory (.js, .ts, .js.map, .css, .html)
func SetupTemplates(cfg server.Config) ([]string, error) {
	files := make([]string, 0)
	log.Debug().Msg("setting up templates")

	log.Debug().Msg("cleaning output directory")
	// clean static output dir
	err := os.RemoveAll(TemplateOutputDir)
	if err != nil {
		return nil,
			fmt.Errorf("error cleaning output directory %v : %w", TemplateOutputDir, err)
	}

	log.Debug().Str("OutputDir", TemplateOutputDir).Msg("creating new output directories")
	// Create/ensure output directory
	if err = file_util.EnsureDir(TemplateOutputDir); err != nil {
		return nil, err
	}

	// Create subdirs
	htmlOutputDir := filepath.Join(TemplateOutputDir, "html")
	if err = file_util.EnsureDir(htmlOutputDir); err != nil {
		return nil, err
	}

	jsOutputDir := filepath.Join(TemplateOutputDir, "js")
	if err = file_util.EnsureDir(jsOutputDir); err != nil {
		return nil, err
	}

	cssOutputDir := filepath.Join(TemplateOutputDir, "css")
	if err = file_util.EnsureDir(cssOutputDir); err != nil {
		return nil, err
	}

	tsOutputDir := filepath.Join(TemplateOutputDir, "src")
	if err = file_util.EnsureDir(tsOutputDir); err != nil {
		return nil, err
	}

	imgOutputDir := filepath.Join(TemplateOutputDir, "img")
	if err = file_util.EnsureDir(imgOutputDir); err != nil {
		return nil, err
	}

	modelOutputDir := filepath.Join(TemplateOutputDir, "model")
	if err = file_util.EnsureDir(modelOutputDir); err != nil {
		return nil, err
	}

	miscFilesOutputDir := filepath.Join(TemplateOutputDir, "files")
	if err = file_util.EnsureDir(miscFilesOutputDir); err != nil {
		return nil, err
	}

	log.Debug().Str("BaseDir", TemplateBaseDir).Msg("ensuring template base directory exists")
	// Ensure base template directory exists
	if !file_util.Exists(TemplateBaseDir) {
		return nil,
			fmt.Errorf("base Dir %v does not exist", TemplateBaseDir)
	}

	// walk through all files in the template resource dir
	err = filepath.Walk(TemplateBaseDir,
		func(path string, info os.FileInfo, err error) error {
			// skip certain directories
			if info.IsDir() && info.Name() == "node_modules" {
				return filepath.SkipDir
			}

			handleCopyFileErr := func(err error) {
				if err != nil {
					log.Fatal().Err(err).Msg("error copying file")
				}
			}

			handleExecuteTemlateErr := func(err error) {
				if err != nil {
					log.Fatal().Err(err).Msg("error executing HTML template")
				}
			}

			switch filepath.Ext(path) {
			case ".html":
				newPath := filepath.Join(htmlOutputDir, filepath.Base(path))
				log.Debug().Str("fromPath", path).Str("toPath", newPath).Msg("moving static web resources")
				handleExecuteTemlateErr(template.ExecuteTemplateHTML(cfg, path, newPath))
			case ".js", ".map":
				newPath := filepath.Join(jsOutputDir, filepath.Base(path))
				if filepath.Base(path) == "serviceWorker.js" || filepath.Base(path) == "serviceWorker.js.map" {
					newPath = filepath.Join("./", filepath.Base(path))
				}
				log.Debug().Str("fromPath", path).Str("toPath", newPath).Msg("moving static web resources")
				handleCopyFileErr(file_util.CopyFile(path, newPath))
			case ".css":
				newPath := filepath.Join(cssOutputDir, filepath.Base(path))
				log.Debug().Str("fromPath", path).Str("toPath", newPath).Msg("moving static web resources")
				handleCopyFileErr(file_util.CopyFile(path, newPath))
			case ".ts":
				newPath := filepath.Join(tsOutputDir, filepath.Base(path))
				log.Debug().Str("fromPath", path).Str("toPath", newPath).Msg("moving static web resources")
				handleCopyFileErr(file_util.CopyFile(path, newPath))
			case ".ico", ".png", ".jpg", ".svg", ".gif":
				newPath := filepath.Join(imgOutputDir, filepath.Base(path))
				log.Debug().Str("fromPath", path).Str("toPath", newPath).Msg("moving static web resources")
				handleCopyFileErr(file_util.CopyFile(path, newPath))
			case ".pdf", ".doc", ".docx", ".xml":
				newPath := filepath.Join(miscFilesOutputDir, filepath.Base(path))
				log.Debug().Str("fromPath", path).Str("toPath", newPath).Msg("moving static web resources")
				handleCopyFileErr(file_util.CopyFile(path, newPath))
			case ".dae", ".obj", ".gltf":
				newPath := filepath.Join(modelOutputDir, filepath.Base(path))
				log.Debug().Str("fromPath", path).Str("toPath", newPath).Msg("moving static web resources")
				handleCopyFileErr(file_util.CopyFile(path, newPath))
			}

			return nil
		})
	if err != nil {
		return nil,
			fmt.Errorf("error walking %v : %w", TemplateBaseDir, err)
	}

	log.Debug().Msg("template setup complete.")
	return files, nil
}

func initServer() *server.Server {
	log.Printf("loading configuration in file: %v", configFile)
	cfg, err := server.LoadConfig(configFile)
	if err != nil {
		log.Fatal().Err(err).Msg("error loading config")
		return nil
	}
	setupLogger(cfg)

	cfg.Print()

	var hostIP string
	if cfg.ChooseIP {

		h, err := ip_util.HostInfo()
		if err != nil {
			log.Fatal().Err(err).Msg("error creating Host Struct")
			return nil
		}

		hostIP, err = ip_util.SelectHost(h.InternalIPs)
		if err != nil {
			log.Fatal().Err(err).Msg("error chosing host IP")
			return nil
		}
	} else {
		hostIP = cfg.IP
	}

	_, err = SetupTemplates(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("error setting up templates")
		return nil
	}

	hub := chat.NewHub()
	go hub.Run()

	addr := hostIP + ":" + cfg.Port

	// generate/execute resource templates

	// create new gorilla mux router
	r := mux.NewRouter()
	// attach pather with handler
	r.HandleFunc("/home", handlers.HomeHandler(cfg.CacheMaxAge))
	r.HandleFunc("/", handlers.RedirectHome())
	r.HandleFunc("/static/js/{scriptname}", handlers.ScriptsHandler(cfg.CacheMaxAge))
	r.HandleFunc("/static/css/{filename}", handlers.CSSHandler(cfg.CacheMaxAge))
	r.HandleFunc("/static/html/{filename}", handlers.HTMLHandler(cfg.CacheMaxAge))
	r.HandleFunc("/static/src/{filename}", handlers.TypeScriptHandler(cfg.CacheMaxAge))
	r.HandleFunc("/static/img/{filename}", handlers.ImageHandler(cfg.CacheMaxAge))
	r.HandleFunc("/static/model/{filename}", handlers.ModelHandler(cfg.CacheMaxAge))
	r.HandleFunc("/manifest.json", handlers.ManifestHandler(cfg.CacheMaxAge))
	r.HandleFunc("/serviceWorker.js", handlers.ServiceWorkerHandler(cfg.CacheMaxAge))
	r.HandleFunc("/serviceWorker.js.map", handlers.ServiceWorkerHandler(cfg.CacheMaxAge))
	r.HandleFunc("/tunes/home", handlers.RedirectConstructionHandler())
	r.HandleFunc("/shop/home", handlers.RedirectConstructionHandler())
	// r.HandleFunc("/chat/{name}", handlers.ChatHomeHandler("", cfg.DebugLog))
	// CHAT HANDLERs
	r.HandleFunc("/chat/home", handlers.ChatHomeHandler(cfg.CacheMaxAge))
	r.HandleFunc("/chat/ws", chat.ServeWs(hub))
	r.HandleFunc("/chat/signup", handlers.RedirectConstructionHandler())
	r.HandleFunc("/chat/signin", handlers.RedirectConstructionHandler())
	// file handler
	r.HandleFunc("/files/{filename}", handlers.MiscFileHandler(cfg.CacheMaxAge))

	// RESUME HANDLER
	r.HandleFunc("/resume/home", handlers.ResumeHomeHandler(cfg.CacheMaxAge))

	// UNDER CONSTRUCTION
	r.HandleFunc("/under-construction", handlers.ConstructionHandler(cfg.CacheMaxAge))

	// HALL OF ART
	r.HandleFunc("/hall-of-art/home", handlers.HallofArtHomeHandler(cfg.CacheMaxAge))

	// DONATE PAGES
	r.HandleFunc("/donate/{cryptoname}", handlers.DonateHandler(cfg.CacheMaxAge))

	fmt.Printf("\n")
	log.Printf("starting Server at: %v...", addr)
	srv := server.NewServer(cfg, r)

	return srv
}

func main() {
	flag.Parse()
	log.Printf("main: starting HTTP server...")
	srv := initServer()
	running := make(chan struct{})
	srv.Run(running)
}
