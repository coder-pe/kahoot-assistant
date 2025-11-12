package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	fmt.Println("===========================================")
	fmt.Println("   Kahoot Assistant - Iniciando...")
	fmt.Println("===========================================\n")

	// Cargar configuración
	config, err := LoadConfig("config.yaml")
	if err != nil {
		fmt.Printf("Error cargando configuración: %v\n", err)
		fmt.Println("\nAsegúrate de tener un archivo config.yaml válido.")
		os.Exit(1)
	}

	fmt.Println("Configuración cargada correctamente")
	fmt.Printf("- Salida en consola: %v\n", config.Output.Console)
	fmt.Printf("- Envío de email: %v\n", config.Output.Email && config.Email.Enabled)
	fmt.Printf("- Guardar en log: %v\n", config.Output.LogFile)
	fmt.Printf("- Tiempo máximo de procesamiento: %d segundos\n\n", config.Timeouts.MaxProcessingTime)

	// Configurar manejo de señales para salida limpia
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Canal para manejar la captura
	captureChan := make(chan bool)

	// Goroutine para manejar Ctrl+C
	go func() {
		<-sigChan
		fmt.Println("\n\nSaliendo del programa...")
		os.Exit(0)
	}()

	// Bucle principal - esperar por Print Screen
	fmt.Println("Programa listo. Presiona Print Screen para capturar y analizar.")
	fmt.Println("Presiona Ctrl+C para salir.\n")

	for {
		// Esperar a que se presione Print Screen
		err := WaitForPrintScreen()
		if err != nil {
			fmt.Printf("Error esperando tecla: %v\n", err)
			continue
		}

		// Procesar la captura en una goroutine separada
		go func() {
			processCapture(config, captureChan)
		}()

		// Esperar a que termine el procesamiento o timeout
		select {
		case <-captureChan:
			fmt.Println("\nListo para la siguiente captura.")
			fmt.Println("Presiona Print Screen nuevamente...\n")
		case <-time.After(time.Duration(config.Timeouts.MaxProcessingTime) * time.Second):
			fmt.Printf("\n¡TIMEOUT! El procesamiento tomó más de %d segundos.\n", config.Timeouts.MaxProcessingTime)
			fmt.Println("Presiona Print Screen para intentar nuevamente...\n")
		}
	}
}

func processCapture(config *Config, done chan bool) {
	defer func() {
		done <- true
	}()

	startTime := time.Now()
	fmt.Println("\n--- Iniciando procesamiento ---")

	// 1. Capturar pantalla
	fmt.Println("1. Capturando pantalla...")
	img, err := CaptureScreen()
	if err != nil {
		fmt.Printf("❌ Error capturando pantalla: %v\n", err)
		return
	}

	// Guardar captura temporal
	tempDir := os.TempDir()
	screenshotPath := filepath.Join(tempDir, fmt.Sprintf("kahoot_screenshot_%d.png", time.Now().Unix()))
	err = SaveScreenshot(img, screenshotPath)
	if err != nil {
		fmt.Printf("❌ Error guardando captura: %v\n", err)
		return
	}
	defer os.Remove(screenshotPath) // Limpiar archivo temporal al finalizar

	fmt.Printf("   Captura guardada en: %s\n", screenshotPath)

	// 2. Extraer texto con OCR
	fmt.Println("2. Extrayendo texto con OCR...")
	question, err := ExtractTextFromImage(screenshotPath, config)
	if err != nil {
		fmt.Printf("❌ Error extrayendo texto: %v\n", err)
		return
	}

	fmt.Println("   Texto extraído exitosamente")
	if config.Output.Console {
		fmt.Printf("\n   PREGUNTA DETECTADA:\n   %s\n\n", question)
	}

	// 3. Consultar a Gemini
	fmt.Println("3. Consultando a Gemini AI...")
	answer, err := AskGemini(question, config)
	if err != nil {
		fmt.Printf("❌ Error consultando Gemini: %v\n", err)
		return
	}

	fmt.Println("   Respuesta recibida de Gemini")

	// 4. Mostrar respuesta en consola si está habilitado
	if config.Output.Console {
		fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                        RESPUESTA                               ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════╝")
		fmt.Println(answer)
		fmt.Println("═══════════════════════════════════════════════════════════════")
	}

	// 5. Enviar correo si está habilitado
	if config.Output.Email && config.Email.Enabled {
		fmt.Println("\n4. Enviando correo...")
		err = SendEmail(question, answer, config)
		if err != nil {
			fmt.Printf("⚠️  Advertencia: Error enviando correo: %v\n", err)
		} else {
			fmt.Printf("   ✓ Correo enviado a: %s\n", config.Email.To)
		}
	}

	// 6. Guardar en log si está habilitado
	if config.Output.LogFile {
		fmt.Println("5. Guardando en log...")
		err = LogQuestionAnswer(question, answer, config)
		if err != nil {
			fmt.Printf("⚠️  Advertencia: Error guardando log: %v\n", err)
		} else {
			fmt.Printf("   ✓ Guardado en: %s\n", config.LogFilePath)
		}
	}

	elapsed := time.Since(startTime)
	fmt.Printf("\n✓ Procesamiento completado en %.2f segundos\n", elapsed.Seconds())

	if elapsed.Seconds() > float64(config.Timeouts.MaxProcessingTime) {
		fmt.Println("⚠️  ADVERTENCIA: Se excedió el tiempo objetivo de procesamiento!")
	}
}
