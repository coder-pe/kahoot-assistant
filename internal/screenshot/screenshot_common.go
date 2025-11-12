package screenshot

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

// SaveScreenshot guarda la imagen en un archivo (común para todos los OS)
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
