package server

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

const (
	sampleConfigFile          = "./sample/sample_config.json"
	sampleConfigFileTLS       = "./sample/sample_config_tls.json"
	sampleConfigFileNoRootTLS = "./sample/sample_config_tls_no_root.json"
	sampleHTML                = "./sample/test.html"
	incorrectConfigFile       = "incorrect.wrong"
)

var (
	client *http.Client
)

func init() {
	os.Setenv("GODEBUG", "x509ignoreCN=0")
}

func pushAttemptHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	err := PushFiles(w, sampleHTML)
	if err != nil {
		log.Error().Err(err).Msg("UNABLE TO PUSH")
	}

	w.WriteHeader(http.StatusOK)
}

func validHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func invalidHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func TestMain(m *testing.M) {
	if !strings.Contains(os.Getenv("GODEBUG"), "x509ignoreCN=0") {
		fmt.Println("Please set GODEBUG=\"x509ignoreCN=0\" or testing will not work")
		os.Exit(1)
	}

	runningChan := make(chan struct{})

	cfg, err := LoadConfig(sampleConfigFile)
	if err != nil {
		os.Exit(-1)
	}

	r := mux.NewRouter()
	// attach basic handler
	r.HandleFunc("/valid", validHandler)
	r.HandleFunc("/invalid", invalidHandler)
	r.HandleFunc("/pushAttempt", pushAttemptHandler)

	srv := NewServer(cfg, r)
	go func(ch chan struct{}) {
		srv.Run(ch)
	}(runningChan)

	// wait until running message
	<-runningChan

	client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}
	time.Sleep(3 * time.Second)

	exitCode := m.Run()

	err = srv.Quit()
	if err != nil {
		os.Exit(-1)
	}

	if srv.isRunning {
		os.Exit(-1)
	}
	os.Exit(exitCode)
}

func TestConfig(t *testing.T) {
	// provide nonexistent file to get incorrect file error
	_, err := LoadConfig(incorrectConfigFile)
	require.ErrorIs(t, err, os.ErrNotExist)

	_, err = LoadConfig(sampleHTML)
	require.ErrorIs(t, err, ErrConfigNotJSON)
}

func TestTLSConfig(t *testing.T) {
	// test loading default config with no TLS
	cfg, err := LoadConfig(sampleConfigFile)
	if err != nil {
		t.Error(err)
	}

	// will throw error since no key pair is not present in config
	_, err = getTLSConfig(cfg)
	if err != os.ErrNotExist { // should be returned if no PEM files found in getTLSConfig
		t.Error(err)
	}

	///////////////////////////////////////////////////////////////

	// test loading default config with  TLS
	cfg, err = LoadConfig(sampleConfigFileTLS)
	require.NoError(t, err)

	f, err := os.Open(cfg.KeyFile)
	require.NoError(t, err)
	f.Close()

	f, err = os.Open(cfg.CertFile)
	require.NoError(t, err)
	f.Close()

	_, err = tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
	require.NoError(t, err)

	_, err = getTLSConfig(cfg)
	require.NoError(t, err)

	// test loading default config with TLS but no root CA specified

	// test loading default config with  TLS
	cfg, err = LoadConfig(sampleConfigFileNoRootTLS)
	require.NoError(t, err)
}

func TestValidGetRequest(t *testing.T) {
	wantStatus := "200 OK"

	r, err := client.Get("http://localhost/valid")
	if err != nil {
		t.Errorf("Error with valid get request to server : %v", err)
	}
	defer r.Body.Close()

	assert.Equal(t, wantStatus, r.Status)
}

func TestInvalidGetRequest(t *testing.T) {
	wantStatus := "404 Not Found"

	r, err := client.Get("http://localhost/invalid")
	if err != nil {
		t.Errorf("Error with invalid get request to server : %v", err)
	}
	defer r.Body.Close()

	assert.Equal(t, wantStatus, r.Status)
}

func TestPushAttemptRequest(t *testing.T) {
	wantStatus := "200 OK"

	r, err := client.Get("http://localhost/pushAttempt")
	if err != nil {
		t.Errorf("Error with invalid get request to server : %v", err)
	}
	defer r.Body.Close()

	assert.Equal(t, wantStatus, r.Status)
}
