package response

import (
	"net/http"

	"github.com/gagansingh3785/typio-service/domain"
)

type GetParagraphsV1Response struct {
	Content string `json:"content"`
}

func NewGetParagraphsV1Response(paragraph *domain.Paragraph) *GetParagraphsV1Response {
	return &GetParagraphsV1Response{
		Content: paragraph.Content,
	}
}

func (resp *GetParagraphsV1Response) GetStatus() int {
	return http.StatusOK
}
