//go:build darwin
// +build darwin

package processor

import (
	"fmt"
	"time"

	"kahoot-assistant/internal/config"
	"kahoot-assistant/internal/keyboard"
	"kahoot-assistant/internal/screenshot"
)

// ProcessCapture procesa una captura según el tipo de trigger (screenshot o clipboard)
func ProcessCapture(cfg *config.Config, triggerType keyboard.TriggerType, done chan bool) {
	defer func() {
		done <- true
	}()

	startTime := time.Now()
	fmt.Println("\n--- Iniciando procesamiento ---")

	if triggerType == keyboard.TriggerScreenshot {
		// 1. Buscar el screenshot más reciente
		fmt.Println("1. Buscando screenshot más reciente...")
		screenshotPath, err := screenshot.GetLatestScreenshot()
		if err != nil {
			fmt.Printf("❌ Error buscando screenshot: %v\n", err)
			return
		}

		// Continuar con el procesamiento de screenshot
		processScreenshot(screenshotPath, cfg, startTime)
	} else if triggerType == keyboard.TriggerClipboard {
		// 1. Procesar desde clipboard
		fmt.Println("1. Procesando desde portapapeles...")
		processClipboard(cfg, startTime)
	}
}
