package handlers

import (
	"net/http"

	zlog "github.com/rs/zerolog/log"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("PONG"))
	if err != nil {
		zlog.Err(err).Msg("Error while sending ping response")
	}
}
