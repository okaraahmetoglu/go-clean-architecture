package container

import (
	"log"

	"go.uber.org/dig"
)

func BuildContainer() (*dig.Container, error) {
	container := dig.New()

	if err := RegisterDependencies(container); err != nil {
		log.Fatalf("Failed to register dependencies: %v", err)
		return container, nil
	}

	return container, nil
}
