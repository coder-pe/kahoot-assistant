package main

import (
	"fmt"
	"net/smtp"
	"time"
)

// SendEmail envía un correo con la pregunta y respuesta
func SendEmail(question, answer string, config *Config) error {
	if !config.Email.Enabled || !config.Output.Email {
		return nil // Email deshabilitado
	}

	// Preparar el mensaje
	subject := "Kahoot Assistant - Pregunta y Respuesta"
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	body := fmt.Sprintf(`Kahoot Assistant - Resultado
========================

Fecha: %s

PREGUNTA:
%s

RESPUESTA:
%s

========================
Este correo fue generado automáticamente por Kahoot Assistant.
`, timestamp, question, answer)

	message := []byte(fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"Content-Type: text/plain; charset=UTF-8\r\n"+
			"\r\n"+
			"%s",
		config.Email.From,
		config.Email.To,
		subject,
		body,
	))

	// Configurar autenticación
	auth := smtp.PlainAuth("", config.Email.From, config.Email.Password, config.Email.SMTPHost)

	// Enviar el correo
	addr := fmt.Sprintf("%s:%d", config.Email.SMTPHost, config.Email.SMTPPort)
	err := smtp.SendMail(addr, auth, config.Email.From, []string{config.Email.To}, message)
	if err != nil {
		return fmt.Errorf("error enviando correo: %w", err)
	}

	return nil
}
