package main

import (
	"fmt"
	"os"
	"time"
)

// LogQuestionAnswer guarda la pregunta y respuesta en el archivo de log
func LogQuestionAnswer(question, answer string, config *Config) error {
	if !config.Output.LogFile {
		return nil // Logging deshabilitado
	}

	// Abrir o crear el archivo de log
	file, err := os.OpenFile(config.LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error abriendo archivo de log: %w", err)
	}
	defer file.Close()

	// Preparar el contenido del log
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf(`
================================================================================
[%s]
PREGUNTA:
%s

RESPUESTA:
%s
================================================================================

`, timestamp, question, answer)

	// Escribir en el archivo
	_, err = file.WriteString(logEntry)
	if err != nil {
		return fmt.Errorf("error escribiendo en log: %w", err)
	}

	return nil
}
