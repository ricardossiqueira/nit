/* Package llm
 */
package llm

import (
	"bytes"
	"fmt"

	"nit/internal/config"
	"nit/internal/git"
)

func BuildDraftPrompt(cfg *config.Config, diff *git.DiffContext, langOverride string) (string, error) {
	lang := cfg.PRStyle.Language
	if langOverride != "" {
		lang = langOverride
	}

	var buf bytes.Buffer
	// TODO: read and parse from a markdown file
	fmt.Fprintf(&buf, "read from markdown")
	fmt.Fprintf(&buf, "%s", lang)

	return buf.String(), nil
}
