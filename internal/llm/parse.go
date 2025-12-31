/*Package llm
 */
package llm

import (
	"fmt"
	"regexp"
	"strings"
)

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
