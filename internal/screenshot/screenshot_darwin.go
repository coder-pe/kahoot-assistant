//go:build darwin
// +build darwin

package screenshot

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	_ "image/png"
)

// GetLatestScreenshot encuentra el screenshot más reciente en el escritorio de macOS
func GetLatestScreenshot() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error obteniendo directorio home: %w", err)
	}

	// En macOS, los screenshots se guardan en el escritorio por defecto
	desktopPath := filepath.Join(homeDir, "Desktop")

	// Buscar archivos de screenshot
	files, err := os.ReadDir(desktopPath)
	if err != nil {
		return "", fmt.Errorf("error leyendo escritorio: %w", err)
	}

	type fileWithTime struct {
		path    string
		modTime time.Time
	}

	var screenshots []fileWithTime

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if (strings.HasPrefix(name, "Screenshot") || strings.HasPrefix(name, "Screen Shot")) &&
			strings.HasSuffix(strings.ToLower(name), ".png") {

			fullPath := filepath.Join(desktopPath, name)
			info, err := os.Stat(fullPath)
			if err != nil {
				continue
			}

			screenshots = append(screenshots, fileWithTime{
				path:    fullPath,
				modTime: info.ModTime(),
			})
		}
	}

	if len(screenshots) == 0 {
		return "", fmt.Errorf("no se encontraron screenshots en el escritorio.\nAsegúrate de tomar un screenshot con Cmd+Shift+3 o Cmd+Shift+4")
	}

	// Ordenar por tiempo de modificación (más reciente primero)
	sort.Slice(screenshots, func(i, j int) bool {
		return screenshots[i].modTime.After(screenshots[j].modTime)
	})

	// Verificar que el screenshot sea reciente (últimos 5 minutos)
	latestScreenshot := screenshots[0]
	if time.Since(latestScreenshot.modTime) > 5*time.Minute {
		return "", fmt.Errorf("el screenshot más reciente tiene más de 5 minutos.\nToma un nuevo screenshot y presiona Enter")
	}

	fmt.Printf("   Usando screenshot: %s\n", filepath.Base(latestScreenshot.path))
	fmt.Printf("   Tomado hace: %s\n", time.Since(latestScreenshot.modTime).Round(time.Second))

	return latestScreenshot.path, nil
}

// CaptureScreenMacOS es un wrapper para compatibilidad
func CaptureScreenMacOS() (image.Image, string, error) {
	screenshotPath, err := GetLatestScreenshot()
	if err != nil {
		return nil, "", err
	}

	// Abrir y leer la imagen
	file, err := os.Open(screenshotPath)
	if err != nil {
		return nil, "", fmt.Errorf("error abriendo screenshot: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, "", fmt.Errorf("error decodificando imagen: %w", err)
	}

	return img, screenshotPath, nil
}
