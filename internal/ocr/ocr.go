package ocr

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"kahoot-assistant/internal/config"
)

// ExtractTextFromImage extrae texto de una imagen usando Tesseract OCR
func ExtractTextFromImage(imagePath string, cfg *config.Config) (string, error) {
	// Determinar el path de tesseract
	tesseractCmd := "tesseract"
	if cfg.OCR.TesseractPath != "" {
		tesseractCmd = cfg.OCR.TesseractPath
	}

	// Construir el comando
	// tesseract imagen.png stdout -l eng
	args := []string{
		imagePath,
		"stdout", // Salida a stdout
		"-l", cfg.OCR.Language,
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
