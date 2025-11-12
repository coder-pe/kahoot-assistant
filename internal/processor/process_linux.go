//go:build linux
// +build linux

package processor

import (
	"fmt"
	"time"

	"kahoot-assistant/internal/config"
	"kahoot-assistant/internal/screenshot"
)

// ProcessCapture procesa una captura de pantalla en Linux
func ProcessCapture(cfg *config.Config, done chan bool) {
	defer func() {
		done <- true
	}()

	startTime := time.Now()
	fmt.Println("\n--- Iniciando procesamiento ---")

	// 1. Buscar el screenshot más reciente
	fmt.Println("1. Buscando screenshot más reciente...")
	screenshotPath, err := screenshot.GetLatestScreenshot()
	if err != nil {
		fmt.Printf("❌ Error buscando screenshot: %v\n", err)
		return
	}

	// Continuar con el procesamiento común
	processScreenshot(screenshotPath, cfg, startTime)
}
