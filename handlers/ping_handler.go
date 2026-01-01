package handlers

import (
	"context"
	"net/http"

	"github.com/gagansingh3785/typio-service/handlers/request"
	"github.com/gagansingh3785/typio-service/handlers/response"
)

func PingHandler() http.HandlerFunc {
	return Handler(
		func(ctx context.Context, r *http.Request) (*request.PingRequest, error) {
			return request.NewPingRequest(ctx, r)
		},
		func(ctx context.Context, r *request.PingRequest) (any, error) {
			return nil, nil
		},
		func(ctx context.Context, _ any) *response.PingResponse {
			return response.NewPingResponse()
		},
	)
}
