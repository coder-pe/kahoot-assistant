package ai

import (
	"context"
	"fmt"
	"time"

	"kahoot-assistant/internal/config"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// AskGemini envía una pregunta a Gemini API y retorna la respuesta
func AskGemini(question string, cfg *config.Config) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeouts.GeminiTimeout)*time.Second)
	defer cancel()

	client, err := genai.NewClient(ctx, option.WithAPIKey(cfg.GeminiAPIKey))
	if err != nil {
		return "", fmt.Errorf("error creando cliente Gemini: %w", err)
	}
	defer client.Close()

	// Usar el modelo Gemini Pro
	model := client.GenerativeModel("gemini-pro")

	// Configurar el prompt para responder preguntas de Kahoot
	prompt := fmt.Sprintf(`Eres un asistente que responde preguntas de Kahoot en inglés.

Pregunta extraída de la pantalla:
%s

Por favor, proporciona la respuesta correcta de manera clara y concisa. Si la pregunta tiene opciones múltiples visibles, indica cuál es la correcta. Si necesitas explicar brevemente tu respuesta, hazlo.`, question)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("error generando respuesta: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no se recibió respuesta de Gemini")
	}

	// Extraer el texto de la respuesta
	var answer string
	for _, part := range resp.Candidates[0].Content.Parts {
		answer += fmt.Sprintf("%v", part)
	}

	return answer, nil
}
