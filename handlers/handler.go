package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gagansingh3785/typio-service/constants"
	"github.com/gagansingh3785/typio-service/errors"
	"github.com/gagansingh3785/typio-service/handlers/request"
	"github.com/gagansingh3785/typio-service/handlers/response"
	zlog "github.com/rs/zerolog/log"
)

type (
	requestProcessor  func(ctx context.Context, r *http.Request) (request.RequestType, error)
	domainProcessor   func(ctx context.Context, request request.RequestType) (any, error)
	responseProcessor func(ctx context.Context, domainObj any) response.ResponseType
)

func Handler(
	reqProcessor requestProcessor,
	domainProcessor domainProcessor,
	respProcessor responseProcessor,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := process(r, reqProcessor, domainProcessor, respProcessor)
		if err != nil {
			writeErrorResponse(w, err)
			return
		}

		writeResponse(w, resp)
	}
}

func process(
	r *http.Request,
	reqProcessor requestProcessor,
	domainProcessor domainProcessor,
	respProcessor responseProcessor,
) (response.ResponseType, errors.ServiceError) {
	ctx := r.Context()
	processedReq, err := reqProcessor(ctx, r)
	if err != nil {
		return nil, errors.As4xxError(err)
	}

	if err := processedReq.Validate(); err != nil {
		return nil, errors.As4xxError(err)
	}

	domainObj, err := domainProcessor(ctx, processedReq)
	if err != nil {
		return nil, errors.AsServiceError(err)
	}

	resp := respProcessor(ctx, domainObj)

	return resp, nil
}

func writeErrorResponse(w http.ResponseWriter, err errors.ServiceError) {
	resp := response.NewErrorResponse(err.Code(), err.Message())
	respBytes, marshalErr := json.Marshal(resp)
	if marshalErr != nil {
		zlog.Err(marshalErr).Msg("Failed to marshal error response")
		http.Error(w, marshalErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(constants.ContentTypeHeader, constants.ContentTypeJSON)
	w.WriteHeader(err.StatusCode())
	if _, err := w.Write(respBytes); err != nil {
		zlog.Err(err).Msg("Failed to write error response")
	}
}

func writeResponse(w http.ResponseWriter, resp response.ResponseType) {
	respBytes, marshalErr := json.Marshal(resp)
	if marshalErr != nil {
		zlog.Err(marshalErr).Msg("Failed to marshal response")
		http.Error(w, marshalErr.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(constants.ContentTypeHeader, constants.ContentTypeJSON)
	w.WriteHeader(resp.GetStatus())
	if _, err := w.Write(respBytes); err != nil {
		zlog.Err(err).Msg("Failed to write response")
	}
}
