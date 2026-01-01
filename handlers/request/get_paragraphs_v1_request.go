package request

import (
	"context"
	"net/http"
)

type GetParagraphsV1Request struct{}

func NewGetParagraphsV1Request(ctx context.Context, r *http.Request) (*GetParagraphsV1Request, error) {
	return &GetParagraphsV1Request{}, nil
}

func (r *GetParagraphsV1Request) Validate() error {
	return nil
}
