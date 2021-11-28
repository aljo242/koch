package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

const (
	// crypto public addresses
	btcAddr string = "3P9F4yxfpmwrLnSzg4c1SCFr177yBUx3WF"
	ltcAddr string = "MPS7pQ3eKyYqdeXDS79ezJ8T5dT31Hxzap"
	ethAddr string = "0xd78804789051626fa7dbf0f9bfba21a5c697c8b8"
	adaAddr string = "addr1q93g6h6cj6kdvjad9tuykr76f77r2pth4r6c9l2sq7amzypkxvjyelkkwprsxm0cpcqgagc9fzgpxqnmgk7ejkfm389sa56wmh"
)

var cryptoAddr map[string]string

func init() {
	cryptoAddr = make(map[string]string)
	initCryptoAddr()
}

func initCryptoAddr() {
	cryptoAddr["btc"] = btcAddr
	cryptoAddr["ltc"] = ltcAddr
	cryptoAddr["eth"] = ethAddr
	cryptoAddr["ada"] = adaAddr
}

// DonateHandler handles an incoming donation request and serves back a page or the crypto address as JSON
func DonateHandler(cacheMaxAge int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r != nil {
			currency := filepath.Base(r.URL.Path)
			log.Debug().Str("Handler", "DonateHandler").Str("Currency", currency).Msg("incoming request")
			w.Header().Set("Content-Type", "text/html; charset=utf-8") // normal header
			fmt.Fprintf(w, "<h1>%v</h1>\n", cryptoAddr[currency])
		}
	}
}
