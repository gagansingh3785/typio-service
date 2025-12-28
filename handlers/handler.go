package handlers

import (
	"context"
	"net/http"

	"github.com/gagansingh3785/typio-service/handlers/request"
	"github.com/gagansingh3785/typio-service/handlers/response"
)

type (
	requestProcessor  func(ctx context.Context, r *http.Request) (request.RequestType, error)
	processor         func(ctx context.Context, request request.RequestType) (any, error)
	responseProcessor func(ctx context.Context, domainObj any) response.ResponseType
)

func Handler(
	reqProcessor requestProcessor,
	domainProcessor processor,
	respProcessor responseProcessor,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		processedReq, err := reqProcessor(ctx, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := processedReq.Validate(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		domainObj, err := domainProcessor(ctx, processedReq)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp := respProcessor(ctx, domainObj)
		w.WriteHeader(resp.GetStatus())
	}
}
