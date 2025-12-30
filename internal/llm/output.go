/* Package llm
 */
package llm

type DraftOutput struct {
	PRTitle       string `json:"pr_title"`
	PRDescription string `json:"pr_description"`
	CommitMessage string `json:"commit_message"`
}
