package handlers

import (
	"context"
	"net/http"

	"github.com/gagansingh3785/typio-service/domain"
	"github.com/gagansingh3785/typio-service/handlers/request"
	"github.com/gagansingh3785/typio-service/handlers/response"
	"github.com/gagansingh3785/typio-service/registry"
)

func ParagraphsV1Handler(svcRegistry *registry.ServiceRegistry) http.HandlerFunc {
	return Handler(
		func(ctx context.Context, r *http.Request) (*request.GetParagraphsV1Request, error) {
			return request.NewGetParagraphsV1Request(ctx, r)
		},
		func(ctx context.Context, _ *request.GetParagraphsV1Request) (*domain.Paragraph, error) {
			return svcRegistry.ParagraphService.GetRandomParagraph(ctx)
		},
		func(ctx context.Context, paragraph *domain.Paragraph) *response.GetParagraphsV1Response {
			return response.NewGetParagraphsV1Response(paragraph)
		},
	)
}
