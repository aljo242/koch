package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/mux"

	"github.com/rs/zerolog/log"
)

// Server ...
type Server struct {
	http.Server
	secure       bool
	ip           string
	port         string
	hostname     string
	shutdownCode int
	userShutdown bool
	wg           *sync.WaitGroup
	quit         chan struct{}
	isRunning    bool
}

func serverShutdownCallback() {
	fmt.Printf("\n")
	log.Printf("shutting down server...")
}

func newTLSConfig(certFile, keyFile, rootCA string) (*tls.Config, error) {

	cer, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return &tls.Config{MinVersion: tls.VersionTLS12}, os.ErrNotExist
	}

	rootCAPool := x509.NewCertPool()

	// read rootCA file into byte
	f, err := os.Open(rootCA)
	if err != nil {
		return &tls.Config{MinVersion: tls.VersionTLS12}, os.ErrNotExist
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return &tls.Config{MinVersion: tls.VersionTLS12}, fmt.Errorf("error reading Root CA file %v : %w", rootCA, err)
	}

	ok := rootCAPool.AppendCertsFromPEM(b)
	if !ok {
		return &tls.Config{MinVersion: tls.VersionTLS12}, fmt.Errorf("error appending Root CA cert %v : %w", rootCA, err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cer},
		RootCAs:      rootCAPool,
		MinVersion:   tls.VersionTLS12,
	}, nil
}

// NewServer ...
func NewServer(secure bool, ip string, port string, certFile string, keyFile string, rootCA string,
	hostname string, shutdownCode int, cmdEnable bool, r *mux.Router) (*Server, error) {
	var err error
	tlsCfg := &tls.Config{MinVersion: tls.VersionTLS12}
	// check if IP is valid
	if net.ParseIP(ip) == nil && ip != "localhost" {
		return nil, ErrInvalidIP
	}
	addr := ip + ":" + port

	if secure {
		tlsCfg, err = newTLSConfig(certFile, keyFile, rootCA)
		if err != nil {
			// log.Fatal().Err(err).Msg("error getting TLS config")
			return nil, err
		}
	}

	quit := make(chan struct{})

	srv := &Server{
		http.Server{
			Handler:           r,
			Addr:              addr,
			WriteTimeout:      15 * time.Second,
			ReadTimeout:       15 * time.Second,
			ReadHeaderTimeout: 15 * time.Second,
			MaxHeaderBytes:    1 << 20,
			TLSConfig:         tlsCfg,
		},
		secure,
		ip,
		port,
		hostname,
		shutdownCode,
		cmdEnable,
		&sync.WaitGroup{},
		quit,
		false,
	}

	srv.RegisterOnShutdown(serverShutdownCallback)

	return srv, nil
}

// Quit sends closes the server quit channel if the server is running
// signaling the server to begin shutting down
// if the server is not running, Quit will return an error
func (srv *Server) Quit() error {
	if srv.isRunning {
		close(srv.quit)
		srv.isRunning = !srv.isRunning
		return nil
	}

	return errors.New("server not running; cannot shutdown")
}

// Run starts the running loop of the server and will fire a message to
// running once running has "fully begun"
func (srv *Server) Run(running chan struct{}) {

	srv.wg.Add(1)

	go func() {
		defer srv.wg.Done() // let main know we are done cleaning up
		// always returns error. ErrServerClosed on graceful close
		if srv.secure {
			// listen for HTTP traffic and redirect to HTTPS
			go func(hostName string) {
				httpAddr := srv.ip + ":80"
				httpsHost := "https://" + hostName
				log.Printf("redirecting all traffic to http://%v/* to %v/*", httpAddr, httpsHost)
				if err := http.ListenAndServe(httpAddr, http.HandlerFunc(RedirectHTTPS(httpsHost))); err != nil {
					log.Fatal().Err(err).Msg("ListenAndServe error")
				}
			}(srv.hostname)

			if err := srv.ListenAndServeTLS("", ""); err != http.ErrServerClosed {
				// unexpected error
				log.Fatal().Err(err).Msg("ListenAndServeTLS() NOT IMPLEMENTED")
			}
		} else if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error
			log.Fatal().Err(err).Msg("ListenAndServe()")
		}
	}()

	// once we have run ListenAdnServe*, we are officially running
	srv.isRunning = true
	close(running)

	getUserInput := func() {
		var code int
		for {
			fmt.Printf("provide shutdown code: \n")
			_, err := fmt.Scanln(&code)
			if err != nil {
				fmt.Printf("error getting input: %v", err)
			}
			if code == srv.shutdownCode {
				break
			}

			fmt.Printf("invalid code.\n")
		}

		err := srv.Quit()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to quit server")
		}
	}

	if srv.userShutdown {
		go getUserInput()
	}

	// wait on a quit
	<-srv.quit
	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}

	// wait for goroutine to stop
	srv.wg.Wait()

	log.Printf("main: done. exiting...")
}
