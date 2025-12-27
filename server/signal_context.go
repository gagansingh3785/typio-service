package server

import (
	"context"
	"os"
	"os/signal"
)

var notifySignals = []os.Signal{os.Interrupt, os.Kill}

func NewSignalContext() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), notifySignals...)
}
