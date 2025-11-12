package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

// CaptureScreen captura toda la pantalla principal
func CaptureScreen() (image.Image, error) {
	// Obtener el número de pantallas
	n := screenshot.NumActiveDisplays()
	if n == 0 {
		return nil, fmt.Errorf("no se encontraron pantallas activas")
	}

	// Capturar la pantalla principal (índice 0)
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return nil, fmt.Errorf("error capturando pantalla: %w", err)
	}

	return img, nil
}

// SaveScreenshot guarda la imagen en un archivo
func SaveScreenshot(img image.Image, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("error creando archivo: %w", err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return fmt.Errorf("error guardando imagen: %w", err)
	}

	return nil
}
