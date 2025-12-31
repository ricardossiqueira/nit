/* Package llm
 */
package llm

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func UnmarshalDraftOutput(data []byte) (DraftOutput, error) {
	var d DraftOutput
	err := json.Unmarshal(data, &d)
	return d, err
}

func (d *DraftOutput) Marshal() ([]byte, error) {
	return json.Marshal(d)
}

type DraftOutput struct {
	PRTitle       string `json:"pr_title"`
	PRDescription string `json:"pr_description"`
	CommitMessage string `json:"commit_message"`
}

func ParseDraftOutput(content string) (*DraftOutput, error) {
	var draft DraftOutput

	draft, err := UnmarshalDraftOutput([]byte(content))
	if err == nil {
		return &draft, nil
	}

	if jsonBlock := extractJSONBlock(content); jsonBlock != "" {
		draft, err := UnmarshalDraftOutput([]byte(jsonBlock))
		if err == nil {
			return &draft, nil
		}
	}

	return nil, fmt.Errorf("error parsing llm output: %w\ncontent: %s", err, content)
}

func extractJSONBlock(markdownText string) string {
	re := regexp.MustCompile(`(?m)^\x60\x60\x60(?:json)\n([\s\S]*?)\x60\x60\x60$`)

	if matches := re.FindStringSubmatch(strings.TrimSpace(markdownText)); len(matches) > 0 {
		jsonStr := strings.TrimSpace(matches[1])
		return jsonStr
	}

	return ""
}
