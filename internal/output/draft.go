/*Package output
 */
package output

import (
	"fmt"

	"nit/internal/llm"
)

func PrintDraft(r *llm.Run) {
	fmt.Printf("PR title: %s\n", r.Response.PRTitle)
	fmt.Printf("PR description: %s\n", r.Response.CommitMessage)
	fmt.Printf("Commit message: %s\n", r.Response.PRDescription)
}
