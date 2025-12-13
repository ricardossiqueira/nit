/*Package llm
 */
package llm

import (
	"encoding/json"
	"time"
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
