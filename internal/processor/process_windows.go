//go:build windows
// +build windows

package processor

import (
	"fmt"
	"os"
	"path/filepath"
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

		// Continuar con el procesamiento de screenshot
		processScreenshot(screenshotPath, cfg, startTime)
	} else if triggerType == keyboard.TriggerClipboard {
		// 1. Procesar desde clipboard
		fmt.Println("1. Procesando desde portapapeles...")
		processClipboard(cfg, startTime)
	}
}
