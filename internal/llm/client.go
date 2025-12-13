/*Package llm
 */
package llm

import (
	"bytes"
	"encoding/json"
	"net/http"

	"nit/internal/config"
)

type LLMResponse struct {
	Content string
}

func Generate(modelCfg config.ModelConfig, prompt string) (*LLMResponse, error) {
	// TODO: implement
	body := map[string]interface{}{
		"model": modelCfg.ModelName,
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "Você é um assistente especializado em revisão de código.",
			},
			{
				"role":    "user",
				"content": prompt,
			},
			"max_tokens":  modelCfg.MaxTokens,
			"temperature": modelCfg.Temperature,
		},
	}

	data, err := json.Marshal(body)
	req, err := http.NewRequest("POST", modelCfg.Endpoint, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// TODO: Continue here

	return &LLMResponse{}, nil
}
