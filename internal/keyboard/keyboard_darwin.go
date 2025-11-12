//go:build darwin
// +build darwin

package keyboard

import (
	"bufio"
	"fmt"
	"os"
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

// IsKeyPressed no está soportado en macOS sin permisos especiales
func IsKeyPressed(vKey int) bool {
	return false
}

// RegisterPrintScreenHotkey no está soportado en macOS sin permisos especiales
func RegisterPrintScreenHotkey(callback func()) error {
	return fmt.Errorf("hotkeys globales no soportados en macOS sin permisos de accesibilidad")
}
