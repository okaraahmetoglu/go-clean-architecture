package mediator

import (
	"fmt"
	"reflect"

	"go.uber.org/dig"
)

// IRequest represents a generic request type
type IRequest interface{}

type Response[T any] struct {
	Data T
}

// IRequestHandler processes requests of a specific type
type IRequestHandler interface {
	Handle(request IRequest) (interface{}, error)
}

/*
// HandlerRegistry
type HandlerRegistry struct {
	container *dig.Container

}

// Yeni bir HandlerRegistry oluştur
func NewHandlerRegistry(container *dig.Container) *HandlerRegistry {
	return &HandlerRegistry{
		container: container,
		handlers:  make(map[reflect.Type]reflect.Type),
	}
}

// Handler kaydet


/*func (r *HandlerRegistry) Register(requestType interface{}, handlerType interface{}) {
	r.handlers[reflect.TypeOf(requestType)] = reflect.TypeOf(handlerType).Elem()
	r.container.Provide(func() RequestHandler {
		handler := reflect.New(reflect.TypeOf(handlerType).Elem()).Interface()
		return handler.(RequestHandler)
	})
}*/

/*func (m *Mediator) Register[T IRequest](constructor func() IRequestHandler[T]) error {
	return m.container.Provide(constructor)
}*/

func (m *Mediator) Register(requestType interface{}, handlerType interface{}) {
	m.handlers[reflect.TypeOf(requestType)] = reflect.TypeOf(handlerType).Elem()
	/*r.container.Provide(func() interface{} {
		return reflect.New(reflect.TypeOf(handlerType).Elem()).Interface()
	})*/
}

// Generic olmayan bir method.
func (m *Mediator) Send(request IRequest, typeName string) (interface{}, error) {

	var handler IRequestHandler
	err := m.container.Invoke(func(h IRequestHandler) {
		handler = h
	})

	if err != nil {
		return nil, fmt.Errorf("handler not found: %w", err)
	}

	return handler.Handle(request)
}

// Mediator sınıfı
type Mediator struct {
	container *dig.Container
	handlers  map[reflect.Type]reflect.Type
	//Registry *HandlerRegistry
}

// Yeni bir Mediator oluştur
func NewMediator(container *dig.Container) *Mediator {
	return &Mediator{container: container, handlers: make(map[reflect.Type]reflect.Type)}
}
