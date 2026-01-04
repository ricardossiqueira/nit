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

func BuildCommitPrompt(cfg *config.Config, diff *git.DiffContext, langOverride string) (*DraftPrompt, error) {
	lang := cfg.PRStyle.Language
	if langOverride != "" {
		lang = langOverride
	}

	var systemPrompt bytes.Buffer
	var userPrompt bytes.Buffer

	fmt.Fprintf(&systemPrompt, "Você é um engenheiro de software experiente. Gere apenas uma mensagem de commit no formato conventional commits (ex: feat: algo curto) em %s. Seja conciso.\n", lang)

	fmt.Fprintf(
		&userPrompt,
		"Resumo das mudanças:\n%s\n\nDiff completo:\n%s",
		diff.Summary,
		diff.RawDiff,
	)

	return &DraftPrompt{
		System: systemPrompt.String(),
		User:   userPrompt.String(),
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

	var systemPrompt bytes.Buffer
	var userPrompt bytes.Buffer

	fmt.Fprintf(&systemPrompt, "Você é um engenheiro de software experiente. Escreva apenas um título de PR curto em %s, máximo de %d caracteres, seguindo o padrão %q. Não inclua mais nada além do título.\n", lang, titleMax, cfg.PRStyle.Title.Pattern)
	fmt.Fprintf(&userPrompt, "Resumo das mudanças:\n%s\n\nDiff completo:\n%s", diff.Summary, diff.RawDiff)

	return &DraftPrompt{
		System: systemPrompt.String(),
		User:   userPrompt.String(),
	}, nil
}

func BuildPRBodyPrompt(cfg *config.Config, diff *git.DiffContext, langOverride string) (*DraftPrompt, error) {
	lang := cfg.PRStyle.Language
	if langOverride != "" {
		lang = langOverride
	}

	var systemPrompt bytes.Buffer
	var userPrompt bytes.Buffer

	fmt.Fprintf(&systemPrompt, "Você é um engenheiro de software experiente. Escreva apenas a descrição detalhada do PR em %s usando markdown.\n", lang)

	if len(cfg.PRStyle.DescriptionSections) == 0 {
		fmt.Fprintf(&systemPrompt, "Use este template:\n## Contexto\n- ...\n\n## Mudanças\n- ...\n\n## Impacto\n- ...\n\n## Testes\n- ...\n")
	} else {
		fmt.Fprintf(&systemPrompt, "Use este template com tópicos:\n")
		for _, section := range cfg.PRStyle.DescriptionSections {
			fmt.Fprintf(&systemPrompt, "## %s\n- ...\n\n", section.Name)
		}
	}

	if len(cfg.PRStyle.CoverageChecklist) > 0 {
		fmt.Fprintf(&systemPrompt, "Checklist de cobertura:\n")
		for _, item := range cfg.PRStyle.CoverageChecklist {
			fmt.Fprintf(&systemPrompt, "- [ ] %s\n", item)
		}
	}

	fmt.Fprintf(&userPrompt, "Resumo das mudanças:\n%s\n\nDiff completo:\n%s", diff.Summary, diff.RawDiff)

	return &DraftPrompt{
		System: systemPrompt.String(),
		User:   userPrompt.String(),
	}, nil
}
