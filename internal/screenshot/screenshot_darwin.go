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

	_ "image/jpeg" // Soporte para JPEG
	_ "image/png"  // Soporte para PNG
)

// GetLatestScreenshot encuentra el screenshot más reciente en el escritorio de macOS
func GetLatestScreenshot() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error obteniendo directorio home: %w", err)
	}

	// En macOS, los screenshots se guardan en el escritorio por defecto
	desktopPath := filepath.Join(homeDir, "Desktop")

	fmt.Println("Buscando screenshots en el escritorio...", desktopPath)

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
		nameLower := strings.ToLower(name)

		// Buscar archivos de imagen (png, jpg, jpeg, heic)
		// Acepta: Screenshot, Screen Shot, Captura, captura, o cualquier .png/.jpg/.heic
		isImageFile := strings.HasSuffix(nameLower, ".png") ||
			strings.HasSuffix(nameLower, ".jpg") ||
			strings.HasSuffix(nameLower, ".jpeg") ||
			strings.HasSuffix(nameLower, ".heic")

		if !isImageFile {
			continue
		}

		// Priorizar archivos con nombres relacionados a screenshots
		isScreenshotLike := strings.Contains(nameLower, "screenshot") ||
			strings.Contains(nameLower, "screen shot") ||
			strings.Contains(nameLower, "captura") ||
			strings.Contains(nameLower, "capture") ||
			strings.Contains(nameLower, "shot")

		// Si es una imagen, agregarla a la lista
		fullPath := filepath.Join(desktopPath, name)
		info, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		// Solo considerar archivos modificados en los últimos 10 minutos
		if time.Since(info.ModTime()) > 10*time.Minute {
			continue
		}

		screenshots = append(screenshots, fileWithTime{
			path:    fullPath,
			modTime: info.ModTime(),
		})

		// Informar si encontramos un archivo
		if isScreenshotLike {
			fmt.Printf("   Encontrado: %s (hace %s)\n", name, time.Since(info.ModTime()).Round(time.Second))
		}
	}

	if len(screenshots) == 0 {
		return "", fmt.Errorf("no se encontraron imágenes recientes en el escritorio.\n\nBusqué archivos .png, .jpg, .jpeg, .heic modificados en los últimos 10 minutos.\n\nAsegúrate de:\n  1. Tomar screenshot con Cmd+Shift+3 o Cmd+Shift+4\n  2. O convertir el .heic a .png\n  3. Presionar Enter dentro de 10 minutos")
	}

	// Ordenar por tiempo de modificación (más reciente primero)
	sort.Slice(screenshots, func(i, j int) bool {
		return screenshots[i].modTime.After(screenshots[j].modTime)
	})

	// El más reciente ya está primero gracias al sort
	latestScreenshot := screenshots[0]

	fmt.Printf("   ✓ Usando imagen: %s\n", filepath.Base(latestScreenshot.path))
	fmt.Printf("   ✓ Modificada hace: %s\n", time.Since(latestScreenshot.modTime).Round(time.Second))
	fmt.Printf("   ✓ Ubicación: %s\n", latestScreenshot.path)

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
