//go:build windows
// +build windows

package processor

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"kahoot-assistant/internal/config"
	"kahoot-assistant/internal/screenshot"
)

// ProcessCapture procesa una captura de pantalla en Windows
func ProcessCapture(cfg *config.Config, done chan bool) {
	defer func() {
		done <- true
	}()

	startTime := time.Now()
	fmt.Println("\n--- Iniciando procesamiento ---")

	// 1. Capturar pantalla
	fmt.Println("1. Capturando pantalla...")
	img, err := screenshot.CaptureScreen()
	if err != nil {
		fmt.Printf("❌ Error capturando pantalla: %v\n", err)
		return
	}

	// Guardar captura temporal
	tempDir := os.TempDir()
	screenshotPath := filepath.Join(tempDir, fmt.Sprintf("kahoot_screenshot_%d.png", time.Now().Unix()))
	err = screenshot.SaveScreenshot(img, screenshotPath)
	if err != nil {
		fmt.Printf("❌ Error guardando captura: %v\n", err)
		return
	}
	defer os.Remove(screenshotPath) // Limpiar archivo temporal al finalizar

	fmt.Printf("   Captura guardada en: %s\n", screenshotPath)

	// Continuar con el procesamiento común
	processScreenshot(screenshotPath, cfg, startTime)
}
