//go:build windows
// +build windows

package screenshot

import (
	"fmt"
	"image"

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
