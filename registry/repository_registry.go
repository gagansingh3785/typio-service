package registry

import (
	"github.com/gagansingh3785/typio-service/database"
	"github.com/gagansingh3785/typio-service/repository"
)

type RepositoryRegistry struct {
	ParagraphRepository repository.Repository
}

func NewRepositoryRegistry(db *database.Database) *RepositoryRegistry {
	return &RepositoryRegistry{
		ParagraphRepository: repository.NewRepository(db),
	}
}
