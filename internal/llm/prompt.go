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

	// INSTRUÇÕES EM INGLÊS (melhor compreensão do LLM)
	fmt.Fprintf(&buf, `You are a senior software engineer proficient in %s.

ANALYZE the git diff below and RETURN ONLY VALID JSON in this EXACT format:

{
  "pr_title": "[concise PR title in %s, max 72 chars]",
  "pr_description": "[detailed description in %s following the Markdown template below]", 
  "commit_message": "[conventional commit message in %s: type: short message]"
}

IMPORTANT: 
- Output ONLY JSON, no explanations or extra text
- pr_title: maximum 72 characters
- pr_description: use the EXACT Markdown template below
- commit_message: conventional commits format (feat:, fix:, etc.)

Use this EXACT Markdown template for pr_description:

`, cfg.Review.Language, lang, lang, lang)

	// TEMPLATE Markdown (permanece igual)
	fmt.Fprintf(&buf, `# %s

## Context
- ...

## Changes
- ...

## Impact
- ...

## Tests
- ...

## Coverage checklist
`, cfg.PRStyle.TitlePattern)

	for _, item := range cfg.PRStyle.CoverageChecklist {
		fmt.Fprintf(&buf, `- [ ] %s\n`, item)
	}

	fmt.Fprintf(&buf, `

**Changes summary:**
%s

**Full diff:**
diff
%s`, diff.Summary, diff.RawDiff)

	return buf.String(), nil
}
