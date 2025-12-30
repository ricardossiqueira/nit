/*Package output
 */
package output

import (
	"fmt"
	"os"
	"path/filepath"

	"nit/internal/llm"
)

func PrintDraft(r *llm.Run) error {
	// TODO: implempent
	fmt.Println(r.Response)

	path := ".nit_draft.md"
	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	if err := os.WriteFile(abs, []byte(r.Response), 0o644); err != nil {
		return err
	}

	fmt.Println("Draft saved to ", abs)

	return nil
}
