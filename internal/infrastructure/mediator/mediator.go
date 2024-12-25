package mediator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type Mediator struct {
	handlers map[reflect.Type]interface{}
}

func NewMediator() *Mediator {
	return &Mediator{
		handlers: make(map[reflect.Type]interface{}),
	}
}

// RegisterHandler bir request türü ile handler'ını kaydeder.
func (m *Mediator) RegisterHandler(request interface{}, handler interface{}) {
	requestType := reflect.TypeOf(request)
	m.handlers[requestType] = handler
}

// Send bir request'i uygun handler'a gönderir.
func (m *Mediator) Send(request interface{}) (interface{}, error) {
	requestType := reflect.TypeOf(request)
	handler, exists := m.handlers[requestType]
	if !exists {
		return nil, fmt.Errorf("no handler registered for %s", requestType)
	}

	// Handler'ı çalıştır
	handlerFunc := reflect.ValueOf(handler)
	response := handlerFunc.Call([]reflect.Value{reflect.ValueOf(request)})

	if len(response) > 1 && !response[1].IsNil() {
		return response[0].Interface(), response[1].Interface().(error)
	}

	return response[0].Interface(), nil
}

// findRequestType examines the methods of a handler to find a method with a single request parameter.
// It returns the type of that request parameter.
func findRequestType(handler interface{}) reflect.Type {
	handlerType := reflect.TypeOf(handler)

	// Eğer pointer referansı alıyorsak, handlerType'ın pointer'dan çıkartılması gerekebilir.
	if handlerType.Kind() == reflect.Ptr {
		handlerType = handlerType.Elem()
	}

	// Handler metodlarını dolaşarak tek parametreli metodu bulmaya çalış
	for i := 0; i < handlerType.NumMethod(); i++ {
		method := handlerType.Method(i)
		// Eğer metodun ismi Handle ya da benzeri bir şeyse
		// ve metodun parametre sayısı 1 ise, request tipini alabiliriz.
		if strings.HasPrefix(method.Name, "Handle") && method.Type.NumIn() == 2 {
			// Parametre türünü döndür
			return method.Type.In(1)
		}
	}

	return nil // Eğer metod bulunmazsa
}

// AutoRegisterHandlers scans only 'handler' directories for handlers
func (m *Mediator) AutoRegisterHandlers(rootDir string) error {
	// Recursive olarak klasörü tara
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Sadece 'handler' klasörlerindeki .go dosyalarını tara
		if info.IsDir() && strings.HasSuffix(path, "/handler") {
			log.Printf("Scanning handler directory: %s", path)
			return m.scanHandlerDirectory(path)
		}

		return nil
	})

	return err
}

// scanHandlerDirectory scans .go files in a handler directory for structs ending with "Handler"
func (m *Mediator) scanHandlerDirectory(handlerDir string) error {
	err := filepath.Walk(handlerDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Sadece .go dosyalarını tara
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			if err := m.registerHandlersFromFile(path); err != nil {
				log.Printf("Failed to register handlers from file %s: %v", path, err)
			}
		}
		return nil
	})

	return err
}

// registerHandlersFromFile parses a Go file to find structs with "Handler" suffix and registers them
func (m *Mediator) registerHandlersFromFile(filePath string) error {
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filePath, nil, parser.AllErrors)
	if err != nil {
		return err
	}

	// Dosya içindeki struct'ları ara
	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			// Struct adı "Handler" ile bitiyorsa
			if strings.HasSuffix(typeSpec.Name.Name, "Handler") {
				handlerInstance := m.createHandlerInstance(typeSpec.Name.Name)
				if handlerInstance != nil {
					// Request türünü belirleyip register et
					requestType := findRequestType(handlerInstance)
					if requestType != nil {
						m.RegisterHandler(requestType, handlerInstance)
						log.Printf("Registered handler: %s", typeSpec.Name.Name)
					}
				}
			}
		}
	}

	return nil
}

// createHandlerInstance creates an instance of a handler struct using reflection
func (m *Mediator) createHandlerInstance(structName string) interface{} {
	// Struct adından bir örnek oluştur (mock örneği için)
	// Burada structName'i kullanarak bir mapping mekanizması ekleyebilirsiniz.
	return reflect.New(reflect.TypeOf(structName)).Interface()
}
