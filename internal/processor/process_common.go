package processor

import (
	"fmt"
	"time"

	"kahoot-assistant/internal/ai"
	"kahoot-assistant/internal/clipboard"
	"kahoot-assistant/internal/config"
	"kahoot-assistant/internal/logger"
	"kahoot-assistant/internal/notification"
	"kahoot-assistant/internal/ocr"
)

// processScreenshot procesa un screenshot dado su path (común para todos los OS)
func processScreenshot(screenshotPath string, cfg *config.Config, startTime time.Time) {
	// 2. Extraer texto con OCR
	fmt.Println("2. Extrayendo texto con OCR...")
	question, err := ocr.ExtractTextFromImage(screenshotPath, cfg)
	if err != nil {
		fmt.Printf("❌ Error extrayendo texto: %v\n", err)
		return
	}

	fmt.Println("   Texto extraído exitosamente")
	if cfg.Output.Console {
		fmt.Printf("\n   PREGUNTA DETECTADA:\n   %s\n\n", question)
	}

	// 3. Consultar a Gemini
	fmt.Println("3. Consultando a Gemini AI...")
	answer, err := ai.AskGemini(question, cfg)
	if err != nil {
		fmt.Printf("❌ Error consultando Gemini: %v\n", err)
		return
	}

	fmt.Println("   Respuesta recibida de Gemini")

	// 4. Mostrar respuesta en consola si está habilitado
	if cfg.Output.Console {
		fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                        RESPUESTA                               ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════╝")
		fmt.Println(answer)
		fmt.Println("═══════════════════════════════════════════════════════════════")
	}

	// 5. Enviar correo si está habilitado
	if cfg.Output.Email && cfg.Email.Enabled {
		fmt.Println("\n4. Enviando correo...")
		err = notification.SendEmail(question, answer, cfg)
		if err != nil {
			fmt.Printf("⚠️  Advertencia: Error enviando correo: %v\n", err)
		} else {
			fmt.Printf("   ✓ Correo enviado a: %s\n", cfg.Email.To)
		}
	}

	// 6. Guardar en log si está habilitado
	if cfg.Output.LogFile {
		fmt.Println("5. Guardando en log...")
		err = logger.LogQuestionAnswer(question, answer, cfg)
		if err != nil {
			fmt.Printf("⚠️  Advertencia: Error guardando log: %v\n", err)
		} else {
			fmt.Printf("   ✓ Guardado en: %s\n", cfg.LogFilePath)
		}
	}

	elapsed := time.Since(startTime)
	fmt.Printf("\n✓ Procesamiento completado en %.2f segundos\n", elapsed.Seconds())

	if elapsed.Seconds() > float64(cfg.Timeouts.MaxProcessingTime) {
		fmt.Println("⚠️  ADVERTENCIA: Se excedió el tiempo objetivo de procesamiento!")
	}
}

// processClipboard procesa el contenido del portapapeles
func processClipboard(cfg *config.Config, startTime time.Time) {
	// 2. Obtener contenido del clipboard
	fmt.Println("2. Leyendo contenido del portapapeles...")
	content, err := clipboard.GetContent()
	if err != nil {
		fmt.Printf("❌ Error leyendo portapapeles: %v\n", err)
		return
	}
	defer clipboard.Cleanup(content)

	var question string

	if content.Type == "text" {
		// Es texto directo
		fmt.Println("   ✓ Texto detectado en el portapapeles")
		question = content.Text

		if cfg.Output.Console {
			fmt.Printf("\n   TEXTO COPIADO:\n   %s\n\n", question)
		}
	} else if content.Type == "image" {
		// Es imagen, aplicar OCR
		fmt.Println("   ✓ Imagen detectada en el portapapeles")
		fmt.Println("   Extrayendo texto con OCR...")

		question, err = ocr.ExtractTextFromImage(content.ImagePath, cfg)
		if err != nil {
			fmt.Printf("❌ Error extrayendo texto de la imagen: %v\n", err)
			return
		}

		fmt.Println("   Texto extraído exitosamente")
		if cfg.Output.Console {
			fmt.Printf("\n   PREGUNTA DETECTADA:\n   %s\n\n", question)
		}
	}

	// 3. Consultar a Gemini
	fmt.Println("3. Consultando a Gemini AI...")
	answer, err := ai.AskGemini(question, cfg)
	if err != nil {
		fmt.Printf("❌ Error consultando Gemini: %v\n", err)
		return
	}

	fmt.Println("   Respuesta recibida de Gemini")

	// 4. Mostrar respuesta en consola si está habilitado
	if cfg.Output.Console {
		fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                        RESPUESTA                               ║")
		fmt.Println("╚════════════════════════════════════════════════════════════════╝")
		fmt.Println(answer)
		fmt.Println("═══════════════════════════════════════════════════════════════")
	}

	// 5. Enviar correo si está habilitado
	if cfg.Output.Email && cfg.Email.Enabled {
		fmt.Println("\n4. Enviando correo...")
		err = notification.SendEmail(question, answer, cfg)
		if err != nil {
			fmt.Printf("⚠️  Advertencia: Error enviando correo: %v\n", err)
		} else {
			fmt.Printf("   ✓ Correo enviado a: %s\n", cfg.Email.To)
		}
	}

	// 6. Guardar en log si está habilitado
	if cfg.Output.LogFile {
		fmt.Println("5. Guardando en log...")
		err = logger.LogQuestionAnswer(question, answer, cfg)
		if err != nil {
			fmt.Printf("⚠️  Advertencia: Error guardando log: %v\n", err)
		} else {
			fmt.Printf("   ✓ Guardado en: %s\n", cfg.LogFilePath)
		}
	}

	elapsed := time.Since(startTime)
	fmt.Printf("\n✓ Procesamiento completado en %.2f segundos\n", elapsed.Seconds())

	if elapsed.Seconds() > float64(cfg.Timeouts.MaxProcessingTime) {
		fmt.Println("⚠️  ADVERTENCIA: Se excedió el tiempo objetivo de procesamiento!")
	}
}
