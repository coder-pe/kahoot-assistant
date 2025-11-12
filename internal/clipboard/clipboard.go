package clipboard

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"

	"golang.design/x/clipboard"

	_ "image/jpeg" // Soporte para JPEG
	_ "image/png"  // Soporte para PNG
)

// ClipboardContent representa el contenido del portapapeles
type ClipboardContent struct {
	Type      string // "text" o "image"
	Text      string // Contenido si es texto
	ImagePath string // Ruta temporal si es imagen
}

// Init inicializa el sistema de clipboard
func Init() error {
	return clipboard.Init()
}

// GetContent obtiene el contenido del clipboard
func GetContent() (*ClipboardContent, error) {
	// Intentar primero obtener imagen
	imgData := clipboard.Read(clipboard.FmtImage)
	if len(imgData) > 0 {
		// Es una imagen, guardarla temporalmente
		img, _, err := image.Decode(bytes.NewReader(imgData))
		if err == nil {
			// Crear archivo temporal
			tempDir := os.TempDir()
			tempFile := filepath.Join(tempDir, "kahoot_clipboard.png")

			file, err := os.Create(tempFile)
			if err != nil {
				return nil, fmt.Errorf("error creando archivo temporal: %w", err)
			}
			defer file.Close()

			if err := png.Encode(file, img); err != nil {
				return nil, fmt.Errorf("error guardando imagen: %w", err)
			}

			return &ClipboardContent{
				Type:      "image",
				ImagePath: tempFile,
			}, nil
		}
	}

	// Si no es imagen, intentar obtener texto
	textData := clipboard.Read(clipboard.FmtText)
	if len(textData) > 0 {
		return &ClipboardContent{
			Type: "text",
			Text: string(textData),
		}, nil
	}

	return nil, fmt.Errorf("el portapapeles está vacío o contiene un formato no soportado")
}

// Cleanup limpia archivos temporales
func Cleanup(content *ClipboardContent) {
	if content != nil && content.Type == "image" && content.ImagePath != "" {
		os.Remove(content.ImagePath)
	}
}
