package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// ExtractTextFromImage extrae texto de una imagen usando Tesseract OCR
func ExtractTextFromImage(imagePath string, config *Config) (string, error) {
	// Determinar el path de tesseract
	tesseractCmd := "tesseract"
	if config.OCR.TesseractPath != "" {
		tesseractCmd = config.OCR.TesseractPath
	}

	// Construir el comando
	// tesseract imagen.png stdout -l eng
	args := []string{
		imagePath,
		"stdout", // Salida a stdout
		"-l", config.OCR.Language,
	}

	cmd := exec.Command(tesseractCmd, args...)

	// Capturar la salida
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Ejecutar el comando
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error ejecutando tesseract: %w, stderr: %s", err, stderr.String())
	}

	// Obtener el texto
	text := stdout.String()

	// Limpiar el texto
	text = strings.TrimSpace(text)
	if text == "" {
		return "", fmt.Errorf("no se pudo extraer texto de la imagen")
	}

	return text, nil
}
