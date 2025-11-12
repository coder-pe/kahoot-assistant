//go:build linux
// +build linux

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

// WaitForPrintScreen espera la entrada del usuario en Linux
// En Linux, podemos usar xdotool o similar, pero requiere X11 y permisos
// Para mantener compatibilidad, usamos el mismo enfoque que macOS
func WaitForPrintScreen() error {
	fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
	fmt.Println("║             MODO LINUX - Instrucciones                         ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Para usar este programa en Linux:")
	fmt.Println()
	fmt.Println("1. Toma un screenshot de la pregunta de Kahoot:")
	fmt.Println("   • Print Screen = Pantalla completa")
	fmt.Println("   • Alt+Print Screen = Ventana activa")
	fmt.Println("   • Shift+Print Screen = Selección de área")
	fmt.Println("   • O usa: gnome-screenshot, flameshot, spectacle, etc.")
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
	fmt.Println("║             MODO LINUX - Instrucciones                         ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Opciones disponibles:")
	fmt.Println()
	fmt.Println("1. SCREENSHOT - Capturar pantalla con Print Screen")
	fmt.Println("   • Toma un screenshot de la pregunta")
	fmt.Println("   • Escribe 's' o 'screenshot' y presiona ENTER")
	fmt.Println()
	fmt.Println("2. CLIPBOARD - Copiar texto/imagen con Ctrl+C")
	fmt.Println("   • Copia el texto o imagen de la pregunta (Ctrl+C)")
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

// IsKeyPressed no está soportado en esta implementación
func IsKeyPressed(vKey int) bool {
	return false
}

// RegisterPrintScreenHotkey no está soportado en esta implementación
func RegisterPrintScreenHotkey(callback func()) error {
	return fmt.Errorf("hotkeys globales requieren configuración adicional en Linux")
}
