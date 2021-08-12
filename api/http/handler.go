package http

import (
	"io"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

type handler struct {
	name string
}

func newHandler(name string) *handler {
	return &handler{name}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	payload := &strings.Builder{}
	_, err := io.Copy(payload, r.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Info().Msg(payload.String())
	w.WriteHeader(http.StatusOK)
}
