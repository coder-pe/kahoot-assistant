package notification

import (
	"fmt"
	"net/smtp"
	"time"

	"kahoot-assistant/internal/config"
)

// SendEmail envía un correo con la pregunta y respuesta
func SendEmail(question, answer string, cfg *config.Config) error {
	if !cfg.Email.Enabled || !cfg.Output.Email {
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
		cfg.Email.From,
		cfg.Email.To,
		subject,
		body,
	))

	// Configurar autenticación
	auth := smtp.PlainAuth("", cfg.Email.From, cfg.Email.Password, cfg.Email.SMTPHost)

	// Enviar el correo
	addr := fmt.Sprintf("%s:%d", cfg.Email.SMTPHost, cfg.Email.SMTPPort)
	err := smtp.SendMail(addr, auth, cfg.Email.From, []string{cfg.Email.To}, message)
	if err != nil {
		return fmt.Errorf("error enviando correo: %w", err)
	}

	return nil
}
