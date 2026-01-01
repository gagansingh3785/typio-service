package registry

import "github.com/gagansingh3785/typio-service/services"

type ServiceRegistry struct {
	ParagraphService services.ParagraphService
}

func NewServiceRegistry(builderRegistry *BuilderRegistry) *ServiceRegistry {
	return &ServiceRegistry{
		ParagraphService: services.NewParagraphService(builderRegistry.ParagraphBuilder),
	}
}
