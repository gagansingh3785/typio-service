package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/gagansingh3785/typio-service/appcontext"
	"github.com/gagansingh3785/typio-service/router"
	zlog "github.com/rs/zerolog/log"
)

func StartHTTPServer() error {
	zlog.Info().Msg("Starting HTTP Server")

	routerOptions := []router.RouterOption{
		router.WithPingRoute(),
	}

	sigCtx, cancelFunc := NewSignalContext()
	defer cancelFunc()

	r := router.NewRouterWithOptions(routerOptions...)

	// TODO: Check Slowloris attack
	// nolint:gosec
	s := &http.Server{
		Addr:    appcontext.GetConfig().GetServerAddr(),
		Handler: r,
	}

	errCh := make(chan error)
	go func() {
		errCh <- s.ListenAndServe()
	}()

	select {
	case <-sigCtx.Done():
		return shutdownServer(s)
	case err := <-errCh:
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	return nil
}

func shutdownServer(s *http.Server) error {
	zlog.Info().Msg("Shutting down HTTP Server")

	// TODO: Explore if we need to have
	// timeout on graceful shutdown of server
	if err := s.Shutdown(context.Background()); err != nil {
		return err
	}

	zlog.Info().Msg("Shutdown successful")

	return nil
}
