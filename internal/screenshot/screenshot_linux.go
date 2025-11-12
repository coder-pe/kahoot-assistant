//go:build linux
// +build linux

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
	_ "image/jpeg"
)

// GetLatestScreenshot encuentra el screenshot más reciente en Linux
func GetLatestScreenshot() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error obteniendo directorio home: %w", err)
	}

	// Ubicaciones comunes de screenshots en diferentes entornos de escritorio Linux
	screenshotDirs := []string{
		filepath.Join(homeDir, "Pictures", "Screenshots"),
		filepath.Join(homeDir, "Pictures"),
		filepath.Join(homeDir, ".local", "share", "screenshots"),
		filepath.Join(homeDir, "Desktop"),
		homeDir,
	}

	type fileWithTime struct {
		path    string
		modTime time.Time
	}

	var screenshots []fileWithTime

	for _, dir := range screenshotDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			continue
		}

		files, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			name := file.Name()
			nameLower := strings.ToLower(name)

			isScreenshot := false
			if strings.Contains(nameLower, "screenshot") ||
				strings.Contains(nameLower, "screen shot") ||
				strings.HasPrefix(nameLower, "shot") ||
				strings.Contains(nameLower, "capture") {
				isScreenshot = true
			}

			ext := filepath.Ext(nameLower)
			if isScreenshot && (ext == ".png" || ext == ".jpg" || ext == ".jpeg") {
				fullPath := filepath.Join(dir, name)
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
	}

	if len(screenshots) == 0 {
		return "", fmt.Errorf("no se encontraron screenshots recientes")
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
	fmt.Printf("   Ubicación: %s\n", filepath.Dir(latestScreenshot.path))
	fmt.Printf("   Tomado hace: %s\n", time.Since(latestScreenshot.modTime).Round(time.Second))

	return latestScreenshot.path, nil
}

// CaptureScreenLinux es un wrapper para compatibilidad
func CaptureScreenLinux() (image.Image, string, error) {
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
