//go:build darwin
// +build darwin

package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TriggerType indica qué tipo de captura realizar
type TriggerType string

const (
	TriggerScreenshot TriggerType = "screenshot"
	TriggerClipboard  TriggerType = "clipboard"
)

// WaitForPrintScreen espera la entrada del usuario en macOS
// En macOS, debido a restricciones de seguridad, no podemos detectar teclas globales
// sin permisos especiales. En su lugar, el usuario debe:
// 1. Tomar screenshot con Cmd+Shift+4 o Cmd+Shift+3
// 2. Presionar Enter en este programa para procesar
func WaitForPrintScreen() error {
	fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
	fmt.Println("║             MODO macOS - Instrucciones                         ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Debido a las restricciones de seguridad de macOS:")
	fmt.Println()
	fmt.Println("1. Toma un screenshot de la pregunta de Kahoot:")
	fmt.Println("   • Cmd+Shift+3 = Pantalla completa")
	fmt.Println("   • Cmd+Shift+4 = Selección de área")
	fmt.Println("   • Cmd+Shift+5 = Herramienta de captura")
	fmt.Println()
	fmt.Println("2. Presiona ENTER aquí para procesar la última captura")
	fmt.Println()
	fmt.Println("Presiona Ctrl+C para salir")
	fmt.Println()
	fmt.Print("Esperando... Presiona ENTER cuando hayas tomado el screenshot: ")

	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("error leyendo entrada: %w", err)
	}

	fmt.Println("\n✓ Procesando última captura de pantalla...")
	return nil
}

// WaitForTrigger espera la entrada del usuario y permite elegir entre screenshot o clipboard
func WaitForTrigger() (TriggerType, error) {
	fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
	fmt.Println("║             MODO macOS - Instrucciones                         ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Opciones disponibles:")
	fmt.Println()
	fmt.Println("1. SCREENSHOT - Capturar pantalla")
	fmt.Println("   • Toma un screenshot con Cmd+Shift+4 (área) o Cmd+Shift+3 (pantalla)")
	fmt.Println("   • Escribe 's' o 'screenshot' y presiona ENTER")
	fmt.Println()
	fmt.Println("2. CLIPBOARD - Copiar texto/imagen con Cmd+C")
	fmt.Println("   • Copia el texto o imagen de la pregunta (Cmd+C)")
	fmt.Println("   • Escribe 'c' o 'clipboard' y presiona ENTER")
	fmt.Println()
	fmt.Println("Presiona Ctrl+C para salir del programa")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("¿Qué tipo de captura quieres usar? (s/c): ")

		input, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("error leyendo entrada: %w", err)
		}

		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "s", "screenshot":
			fmt.Println("\n✓ Procesando última captura de pantalla...")
			return TriggerScreenshot, nil
		case "c", "clipboard":
			fmt.Println("\n✓ Leyendo contenido del portapapeles...")
			return TriggerClipboard, nil
		default:
			fmt.Println("❌ Opción inválida. Por favor escribe 's' para screenshot o 'c' para clipboard.")
		}
	}
}

// IsKeyPressed no está soportado en macOS sin permisos especiales
func IsKeyPressed(vKey int) bool {
	return false
}

// RegisterPrintScreenHotkey no está soportado en macOS sin permisos especiales
func RegisterPrintScreenHotkey(callback func()) error {
	return fmt.Errorf("hotkeys globales no soportados en macOS sin permisos de accesibilidad")
}
