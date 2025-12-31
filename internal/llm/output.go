/* Package llm
 */
package llm

import "encoding/json"

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
