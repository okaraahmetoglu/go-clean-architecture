package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

// Logger için yapı (struct)
type Logger struct {
	File *os.File
}

var (
	l    *Logger
	once sync.Once
)

func NewLogger() (*Logger, error) {
	var err error
	once.Do(func() {
		file, e := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if e != nil {
			err = e
			return
		}
		log.SetOutput(file)
		log.Println("Logger başlatıldı.")
		l = &Logger{File: file}
	})
	return l, err
}

//var l *Logger

// GormLogger'ı başlat
func NewGormLogger() (*Logger, error) {
	file, err := os.OpenFile("gorm.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	log.SetOutput(file)
	log.Println("Logger başlatıldı.")

	return &Logger{File: file}, nil
}

// Logger'ı başlat
/*func NewLogger() (*Logger, error) {
	if l == nil {
		file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}

		log.SetOutput(file)
		log.Println("Logger başlatıldı.")

		l = &Logger{File: file}

	}
	return l, nil
}*/

func (l *Logger) Println(message string) {
	log.Println("[INFO]: " + message)
}

// Fatalf log (formatlı Fatal log yazma ve çıkış)
func (l *Logger) Printf(format string, args ...interface{}) {
	// Formatlı mesajı oluştur
	message := fmt.Sprintf(format, args...)

	// Fatal log mesajını yaz
	log.Println("[FATAL]: " + message)

}

// Log yazma metodu
func (l *Logger) Info(message string) {
	log.Println("[INFO]: " + message)
}

// Fatalf log (formatlı Fatal log yazma ve çıkış)
func (l *Logger) Fatalf(format string, args ...interface{}) {
	// Formatlı mesajı oluştur
	message := fmt.Sprintf(format, args...)

	// Fatal log mesajını yaz
	log.Println("[FATAL]: " + message)

	// Dosyayı kapat
}

// Log dosyasını kapatma
func (l *Logger) Close() {
	if l.File != nil {
		l.File.Close()
	}
}
