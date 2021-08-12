// +build !aws_lambda

package cmd

import (
	"os"
	"os/signal"
	"syscall"

	http "github.com/daesu/serverless-go/api/http"
	"github.com/rs/zerolog/log"
)

// Start the API server
func Start(server *http.Server) error {
	log.Info().Msg("Starting server")

	errCh := make(chan error, 1)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Error().Err(err).Msg("http server failed")
			errCh <- err
		}
	}()

	// catch errors and termination signals
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case <-interrupt:
		log.Info().Msg("shutdown is requested")
	case <-errCh:
	}
	return nil
}
