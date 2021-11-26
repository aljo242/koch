package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

// Config is the general struct holds parsed JSON config info
type Config struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	IP           string `json:"IP"`
	ChooseIP     bool   `json:"chooseIP"`
	HTTPS        bool   `json:"secure"`
	DebugLog     bool   `json:"debugLog"`
	CacheMaxAge  int    `json:"cacheMaxAge"`
	ShutdownCode int    `json:"shutdownCode"`
	UserShutdown bool   `json:"userShutdown"`
	CertFile     string `json:"certFile"`
	KeyFile      string `json:"keyFile"`
	RootCA       string `json:"rootCA"`

	// TODO add more
}

// errors declarations
var ErrConfigNotJSON = errors.New("config file not JSON")

// LoadConfig returns a config struct given a valid config.json file
func LoadConfig(filename string) (Config, error) {
	cfg := Config{}
	filePath := filepath.Clean(filename)
	cfgFile, err := os.Open(filePath)
	if err != nil {
		return Config{}, os.ErrNotExist
	}
	defer func() {
		err := cfgFile.Close()
		if err != nil {
			log.Error().Err(err).Str("filename", filePath).Msg("error closing the file")
		}
	}()

	jsonParser := json.NewDecoder(cfgFile)
	err = jsonParser.Decode(&cfg)
	if err != nil {
		return Config{},
			ErrConfigNotJSON
	}

	return cfg, nil
}

// Print provides a pretty formatted print of a ServerConfig
func (cfg *Config) Print() {
	fmt.Printf("\n-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-\n")
	fmt.Printf("Config:\n")
	fmt.Printf("\tHost:\t\t%v\n", cfg.Host)
	fmt.Printf("\tPort:\t\t%v\n", cfg.Port)
	fmt.Printf("\tIP:\t\t%v\n", cfg.IP)
	fmt.Printf("\tChooseIP:\t%t\n", cfg.ChooseIP)
	fmt.Printf("\tHTTPS:\t\t%t\n", cfg.HTTPS)
	fmt.Printf("\tDebugLog:\t%t\n", cfg.DebugLog)
	fmt.Printf("\tDebugLog:\t%t\n", cfg.DebugLog)
	fmt.Printf("\tShutdownCode:\t%d\n", cfg.ShutdownCode)
	fmt.Printf("\tCertFile:\t%v\n", cfg.CertFile)
	fmt.Printf("\tKeyFile:\t%v\n", cfg.KeyFile)
	fmt.Printf("\tRootCA:\t\t%v\n", cfg.RootCA)
	fmt.Printf("-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-\n\n")
}
