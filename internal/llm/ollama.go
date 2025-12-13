/*Package llm
 */
package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const TIMEOUT = 5

func UnmarshalTags(data []byte) (Tags, error) {
	var t Tags
	err := json.Unmarshal(data, &t)
	return t, err
}

func (t *Tags) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

type Tags struct {
	Models []Model `json:"models"`
}

type Model struct {
	Name       string    `json:"name"`
	Model      string    `json:"model"`
	ModifiedAt time.Time `json:"modified_at"`
	Size       int64     `json:"size"`
	Digest     string    `json:"digest"`
	Details    Details   `json:"details"`
}

type Details struct {
	ParentModel       string   `json:"parent_model"`
	Format            string   `json:"format"`
	Family            string   `json:"family"`
	Families          []string `json:"families"`
	ParameterSize     string   `json:"parameter_size"`
	QuantizationLevel string   `json:"quantization_level"`
}

func (t *Tags) PrintList() {
	var buf bytes.Buffer
	fmt.Fprint(&buf, "Available models:\n")
	for _, tag := range t.Models {
		fmt.Fprintf(&buf, " - %s\n", tag.Name)
	}
	fmt.Print(&buf)
}

func Ping() error {
	client := &http.Client{Timeout: time.Duration(TIMEOUT) * time.Second}
	_, err := client.Get("http://localhost:11434/api/version")
	if err != nil {
		return fmt.Errorf("could not connect Ollama server: %w", err)
	}
	return nil
}

func GetTags() (*Tags, error) {
	client := &http.Client{Timeout: time.Duration(TIMEOUT) * time.Second}
	resp, err := client.Get("http://localhost:11434/api/tags")
	if err != nil {
		return nil, fmt.Errorf("error fetching Ollama models: %w", err)
	}

	var tags *Tags
	if err := json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		return nil, err
	}

	return tags, nil
}
