/*Package llm
 */
package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"nit/internal/config"
)

type LLMResponse struct {
	Content string
}

func Generate(modelCfg config.ModelConfig, prompt string) (*LLMResponse, error) {
	// TODO: implement
	body := map[string]any{
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
		},
		"max_tokens":  modelCfg.MaxTokens,
		"temperature": modelCfg.Temperature,
		"stream":      false,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: time.Duration(modelCfg.Timeout) * time.Second}
	req, err := http.NewRequest("POST", modelCfg.Endpoint, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("llm error: status %d", resp.StatusCode)
	}

	var raw map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close resp body: %v", err)
	}

	content := extractContent(raw)

	return &LLMResponse{Content: content}, nil
}

func extractContent(raw map[string]any) string {
	// TODO: implement real parsing
	b, _ := json.Marshal(raw)
	return string(b)
}
