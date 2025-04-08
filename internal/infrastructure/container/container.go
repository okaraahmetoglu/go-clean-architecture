package container

import (
	"github.com/okaraahmetoglu/go-clean-architecture/internal/infrastructure/logger"
	"go.uber.org/dig"
)

func BuildContainer(logger *logger.Logger) (*dig.Container, error) {
	container := dig.New()

	if err := RegisterDependencies(container, logger); err != nil {
		logger.Fatalf("Failed to register dependencies: %v", err)
		return container, nil
	}

	return container, nil
}
