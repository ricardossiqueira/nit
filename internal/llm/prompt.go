/* Package llm
 */
package llm

import (
	"bytes"
	"fmt"

	"nit/internal/config"
	"nit/internal/git"
)

type DraftPrompt struct {
	System string
	User   string
}

func BuildDraftPrompt(cfg *config.Config, diff *git.DiffContext, langOverride string) (*DraftPrompt, error) {
	lang := cfg.PRStyle.Language
	if langOverride != "" {
		lang = langOverride
	}

	titlePattern := cfg.PRStyle.Title.Pattern
	if titlePattern == "" {
		titlePattern = "[{type}] {scope}: {summary}"
	}

	titleMax := cfg.PRStyle.Title.MaxLength
	if titleMax == 0 {
		titleMax = 72
	}

	authoringLang := cfg.Review.Language
	if authoringLang == "" {
		authoringLang = lang
	}

	var systemPrompt bytes.Buffer
	var userPrompt bytes.Buffer

	fmt.Fprintf(&systemPrompt, `You are a senior software engineer proficient in %s.

ANALYZE the git diff below and RETURN ONLY VALID JSON in this EXACT format:

{
  "pr_title": "[concise PR title in %s, max %d chars]",
  "pr_description": "[detailed description in %s following the Markdown template below]", 
  "commit_message": "[conventional commit message in %s: type: short message]"
}

IMPORTANT: 
- Output ONLY JSON, no explanations or extra text
- pr_title: maximum %d characters
- pr_description: use the EXACT Markdown template below
- commit_message: conventional commits format (feat:, fix:, etc.)

Use this EXACT Markdown template for pr_description:

`, authoringLang, lang, titleMax, lang, lang, titleMax)

	fmt.Fprintf(&systemPrompt, "# %s\n\n", titlePattern)

	if len(cfg.PRStyle.DescriptionSections) == 0 {
		fmt.Fprintf(&systemPrompt, "## Context\n- ...\n\n## Changes\n- ...\n\n## Impact\n- ...\n\n## Tests\n- ...\n\n")
	} else {
		for _, section := range cfg.PRStyle.DescriptionSections {
			fmt.Fprintf(&systemPrompt, "## %s\n- ...\n\n", section.Name)
		}
	}

	fmt.Fprintf(&systemPrompt, "## Coverage checklist\n")

	for _, item := range cfg.PRStyle.CoverageChecklist {
		fmt.Fprintf(&systemPrompt, `- [ ] %s\n`, item)
	}

	fmt.Fprintf(
		&userPrompt,
		"**Changes summary:**\n%s\n\n**Full diff:**\ndiff\n%s",
		diff.Summary,
		diff.RawDiff,
	)

	return &DraftPrompt{
		System: systemPrompt.String(),
		User:   userPrompt.String(),
	}, nil
}
