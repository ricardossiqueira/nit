/*Package llm
 */
package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"nit/internal/config"
)

func UnmarshalResponse(data []byte) (Response, error) {
	var r Response
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Response) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Response struct {
	CreatedAt          time.Time `json:"created_at"`
	Done               bool      `json:"done"`
	DoneReason         string    `json:"done_reason"`
	EvalCount          int64     `json:"eval_count"`
	EvalDuration       int64     `json:"eval_duration"`
	LoadDuration       int64     `json:"load_duration"`
	Message            Message   `json:"message"`
	Model              string    `json:"model"`
	PromptEvalCount    int64     `json:"prompt_eval_count"`
	PromptEvalDuration int64     `json:"prompt_eval_duration"`
	TotalDuration      int64     `json:"total_duration"`
}

type Message struct {
	Content  string `json:"content"`
	Role     string `json:"role"`
	Thinking string `json:"thinking"`
}

func Generate(ctx context.Context, store RunStore, modelCfg config.ModelConfig, prompt string) (*Run, error) {
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
		// TODO: implement stream handler
		"stream": false,
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

	b, _ := json.Marshal(raw)
	llmResponse, err := UnmarshalResponse(b)
	if err != nil {
		return nil, fmt.Errorf("failed to parse llm response %w", err)
	}

	run := &Run{
		Model:    modelCfg.ModelName,
		Endpoint: modelCfg.Endpoint,
		//TODO: fix
		System:     "Você é um assistente especializado em revisão de código.",
		Prompt:     prompt,
		Response:   llmResponse.Message.Content,
		DurationMS: llmResponse.TotalDuration,
		CreatedAt:  time.Now(),
	}

	if store != nil {
		_ = store.SaveRun(ctx, run)
	}

	return run, nil

}
