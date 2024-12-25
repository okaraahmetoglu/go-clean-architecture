package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ServerConfig struct {
	Port int `json:"port"`
}

type DatabaseConfig struct {
	URL string `json:"url"`
}

type LogConfig struct {
	Level int `json:"level"`
}

type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
	Log      LogConfig      `json:"log"`
}

func Load() (*Config, error) {
	// JSON konfigürasyon dosyasını açma
	file, err := os.Open("config.json") // config.json dosyasını açıyoruz
	if err != nil {
		return nil, fmt.Errorf("config dosyası açılırken hata: %v", err)
	}
	defer file.Close()

	// JSON verisini çözümleme
	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("config dosyası çözümleme hatası: %v", err)
	}

	return &config, nil
}
