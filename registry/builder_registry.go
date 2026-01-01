package registry

import (
	"github.com/gagansingh3785/typio-service/builders"
)

type BuilderRegistry struct {
	ParagraphBuilder builders.ParagraphsBuilder
}

func NewBuilderRegistry(repoRegistry *RepositoryRegistry) *BuilderRegistry {
	return &BuilderRegistry{
		ParagraphBuilder: builders.NewParagraphsBuilder(repoRegistry.ParagraphRepository),
	}
}
