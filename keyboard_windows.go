//go:build windows
// +build windows

package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
)

const (
	VK_SNAPSHOT = 0x2C // Print Screen key
)

// GetAsyncKeyState verifica el estado de una tecla
func GetAsyncKeyState(vKey int) uint16 {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(vKey))
	return uint16(ret)
}

// WaitForPrintScreen espera hasta que se presione la tecla Print Screen
func WaitForPrintScreen() error {
	fmt.Println("Esperando que presiones la tecla Print Screen...")
	fmt.Println("Presiona Ctrl+C para salir")

	// Estado anterior de la tecla
	var previousState uint16 = 0

	for {
		// Verificar el estado de Print Screen
		currentState := GetAsyncKeyState(VK_SNAPSHOT)

		// Detectar cuando la tecla se presiona (transición de no presionada a presionada)
		// El bit más significativo indica si la tecla está presionada actualmente
		if (currentState&0x8000) != 0 && (previousState&0x8000) == 0 {
			fmt.Println("\nTecla Print Screen detectada!")
			return nil
		}

		previousState = currentState

		// Pequeña pausa para no consumir 100% CPU
		time.Sleep(50 * time.Millisecond)
	}
}

// IsKeyPressed verifica si una tecla específica está presionada
func IsKeyPressed(vKey int) bool {
	state := GetAsyncKeyState(vKey)
	return (state & uint16(0x8000)) != 0
}

// RegisterHotKey registra una tecla de acceso rápido (alternativa más eficiente)
func RegisterPrintScreenHotkey(callback func()) error {
	user32 := windows.NewLazySystemDLL("user32.dll")
	procRegisterHotKey := user32.NewProc("RegisterHotKey")
	procUnregisterHotKey := user32.NewProc("UnregisterHotKey")
	procGetMessage := user32.NewProc("GetMessageW")

	const (
		MOD_NOREPEAT = 0x4000
		WM_HOTKEY    = 0x0312
	)

	// Registrar la tecla Print Screen como hotkey
	ret, _, err := procRegisterHotKey.Call(
		0,              // NULL window handle
		1,              // Hotkey ID
		MOD_NOREPEAT,   // No repetir
		VK_SNAPSHOT,    // Print Screen
	)

	if ret == 0 {
		return fmt.Errorf("error registrando hotkey: %v", err)
	}

	defer procUnregisterHotKey.Call(0, 1)

	fmt.Println("Hotkey registrado. Presiona Print Screen para capturar.")
	fmt.Println("Presiona Ctrl+C para salir")

	// Estructura MSG para GetMessage
	type MSG struct {
		Hwnd    uintptr
		Message uint32
		WParam  uintptr
		LParam  uintptr
		Time    uint32
		Pt      struct{ X, Y int32 }
	}

	var msg MSG
	for {
		ret, _, _ := procGetMessage.Call(
			uintptr(unsafe.Pointer(&msg)),
			0,
			0,
			0,
		)

		if ret == 0 {
			break
		}

		if msg.Message == WM_HOTKEY {
			fmt.Println("\nPrint Screen presionado!")
			callback()
		}
	}

	return nil
}
