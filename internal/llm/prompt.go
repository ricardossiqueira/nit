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
	fmt.Fprintf(&buf, "You are a senior software engineer with high proeficiency in %s\n", cfg.Review.Language)
	fmt.Fprintf(&buf, "Generate a PR title and description in %s.\n", lang)
	fmt.Fprintf(&buf, "Follow the exact Markdown pattern provided below:\n\n")

	fmt.Fprintf(&buf, "# Suggested title\n")
	fmt.Fprintf(&buf, "%s\n\n", cfg.PRStyle.TitlePattern)
	fmt.Fprintf(&buf, "---\n\n")

	fmt.Fprintf(&buf, "## Context\n")
	fmt.Fprintf(&buf, "- ...\n\n")

	fmt.Fprintf(&buf, "## Changes\n")
	fmt.Fprintf(&buf, "- ...\n\n")

	fmt.Fprintf(&buf, "## Impact")
	fmt.Fprintf(&buf, "- ...\n\n")

	fmt.Fprintf(&buf, "## Tests\n\n")
	fmt.Fprintf(&buf, "- ...\n\n")

	fmt.Fprintf(&buf, "## Coverage checklist\n")
	for _, item := range cfg.PRStyle.CoverageChecklist {
		fmt.Fprintf(&buf, "- [ ] %s\n", item)
	}

	fmt.Fprintf(&buf, "\n\nChanges summary (for context):\n%s\n\n", diff.Summary)

	fmt.Fprintf(&buf, "Full diff (truncated if necessary):\n```%s```\n", diff.RawDiff)

	return buf.String(), nil
}
