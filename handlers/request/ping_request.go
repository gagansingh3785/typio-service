package request

import (
	"context"
	"net/http"
)

type PingRequest struct{}

func NewPingRequest(ctx context.Context, r *http.Request) (*PingRequest, error) {
	return &PingRequest{}, nil
}

func (r *PingRequest) Validate() error {
	return nil
}
