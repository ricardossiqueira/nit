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
)

func PrintDraft(resp *llm.Run, format OutputFormat) error {
	fmt.Println(resp.Response)
	switch format {
	case FormatPretty:
		fmt.Printf("🆕 **PR Title:** %s\n\n📝 **Description:**\n%s\n\n💬 **Commit:** %s\n",
			resp.Response.PRTitle, resp.Response.PRDescription, resp.Response.CommitMessage)
	case FormatJSON:
		json.NewEncoder(os.Stdout).Encode(resp.Response)
	case FormatCommit:
		fmt.Println(resp.Response.CommitMessage)
	case FormatPRTitle:
		fmt.Println(resp.Response.PRTitle)
	case FormatPRBody:
		fmt.Println(resp.Response.PRDescription)
	}
	return nil
}
