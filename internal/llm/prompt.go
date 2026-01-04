/* Package llm
 */
package llm

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"nit/internal/config"
	"nit/internal/git"
)

type DraftPrompt struct {
	System string
	User   string
}

type promptTemplateData struct {
	Language          string
	TitleMax          int
	TitlePattern      string
	Sections          []config.Section
	CoverageChecklist []string
	Summary           string
	RawDiff           string
}

func BuildCommitPrompt(cfg *config.Config, diff *git.DiffContext, langOverride string) (*DraftPrompt, error) {
	lang := cfg.PRStyle.Language
	if langOverride != "" {
		lang = langOverride
	}

	data := promptTemplateData{
		Language: lang,
		Summary:  diff.Summary,
		RawDiff:  diff.RawDiff,
	}

	prompt, err := renderPromptTemplate(filepath.Join("prompts", "commit.md"), data)
	if err != nil {
		return nil, fmt.Errorf("failed to render commit prompt: %w", err)
	}

	return &DraftPrompt{
		System: prompt,
		User:   "",
	}, nil
}

func BuildPRTitlePrompt(cfg *config.Config, diff *git.DiffContext, langOverride string) (*DraftPrompt, error) {
	lang := cfg.PRStyle.Language
	if langOverride != "" {
		lang = langOverride
	}
	titleMax := cfg.PRStyle.Title.MaxLength
	if titleMax == 0 {
		titleMax = 72
	}

	data := promptTemplateData{
		Language:     lang,
		TitleMax:     titleMax,
		TitlePattern: cfg.PRStyle.Title.Pattern,
		Summary:      diff.Summary,
		RawDiff:      diff.RawDiff,
	}

	prompt, err := renderPromptTemplate(filepath.Join("prompts", "title.md"), data)
	if err != nil {
		return nil, fmt.Errorf("failed to render pr title prompt: %w", err)
	}

	return &DraftPrompt{
		System: prompt,
		User:   "",
	}, nil
}

func BuildPRBodyPrompt(cfg *config.Config, diff *git.DiffContext, langOverride string) (*DraftPrompt, error) {
	lang := cfg.PRStyle.Language
	if langOverride != "" {
		lang = langOverride
	}

	data := promptTemplateData{
		Language:          lang,
		Sections:          cfg.PRStyle.DescriptionSections,
		CoverageChecklist: cfg.PRStyle.CoverageChecklist,
		Summary:           diff.Summary,
		RawDiff:           diff.RawDiff,
	}

	prompt, err := renderPromptTemplate(filepath.Join("prompts", "body.md"), data)
	if err != nil {
		return nil, fmt.Errorf("failed to render pr body prompt: %w", err)
	}

	return &DraftPrompt{
		System: prompt,
		User:   "",
	}, nil
}

func renderPromptTemplate(path string, data promptTemplateData) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read prompt template %s: %w", path, err)
	}

	tmpl, err := template.New(filepath.Base(path)).Parse(string(content))
	if err != nil {
		return "", fmt.Errorf("parse prompt template %s: %w", path, err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("execute prompt template %s: %w", path, err)
	}

	return buf.String(), nil
}
