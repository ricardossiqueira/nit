/*Package output
 */
package output

import (
	"encoding/json"
	"fmt"
	"os"

	"nit/internal/llm"
)

type OutputFormat string

const (
	FormatPretty  OutputFormat = "pretty"
	FormatJSON    OutputFormat = "json"
	FormatCommit  OutputFormat = "commit"
	FormatPRTitle OutputFormat = "pr-title"
	FormatPRBody  OutputFormat = "pr-body"
	FormatPR      OutputFormat = "pr"
)

func PrintDraft(resp *llm.DraftOutput, format OutputFormat) error {
	switch format {
	case FormatPretty:
		fmt.Printf("🆕 **PR Title:** %s\n\n📝 **Description:**\n%s\n\n💬 **Commit:** %s\n",
			resp.PRTitle, resp.PRDescription, resp.CommitMessage)
	case FormatJSON:
		json.NewEncoder(os.Stdout).Encode(resp)
	case FormatCommit:
		fmt.Println(resp.CommitMessage)
	case FormatPRTitle:
		fmt.Println(resp.PRTitle)
	case FormatPRBody:
		fmt.Println(resp.PRDescription)
	case FormatPR:
		fmt.Println(resp.PRTitle)
		fmt.Println()
		fmt.Println(resp.PRDescription)
	}
	return nil
}
