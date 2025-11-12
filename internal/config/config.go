package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	GeminiAPIKey string `yaml:"gemini_api_key"`
	Email        struct {
		Enabled  bool   `yaml:"enabled"`
		SMTPHost string `yaml:"smtp_host"`
		SMTPPort int    `yaml:"smtp_port"`
		From     string `yaml:"from"`
		Password string `yaml:"password"`
		To       string `yaml:"to"`
	} `yaml:"email"`
	Output struct {
		Console bool `yaml:"console"`
		Email   bool `yaml:"email"`
		LogFile bool `yaml:"log_file"`
	} `yaml:"output"`
	OCR struct {
		Language      string `yaml:"language"`
		TesseractPath string `yaml:"tesseract_path"`
	} `yaml:"ocr"`
	Timeouts struct {
		MaxProcessingTime int `yaml:"max_processing_time"`
		GeminiTimeout     int `yaml:"gemini_timeout"`
		EmailTimeout      int `yaml:"email_timeout"`
	} `yaml:"timeouts"`
	LogFilePath string `yaml:"log_file_path"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error leyendo config: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error parseando config: %w", err)
	}

	// Validaciones básicas
	if config.GeminiAPIKey == "" || config.GeminiAPIKey == "TU_TOKEN_AQUI" {
		return nil, fmt.Errorf("debes configurar gemini_api_key en config.yaml")
	}

	if config.Output.Email && config.Email.Enabled {
		if config.Email.From == "" || config.Email.To == "" || config.Email.Password == "" {
			return nil, fmt.Errorf("debes configurar los datos de email en config.yaml")
		}
	}

	return &config, nil
}
