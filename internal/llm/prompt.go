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
	fmt.Fprintf(&buf, "Você é um engenheiro de software sênior especializado em Golang.")
	fmt.Fprintf(&buf, "Gere um título e uma descricão de Pull request em %s.\n", lang)
	fmt.Fprintf(&buf, "Siga exatamente o formato Markdown abaixo:\n\n")

	fmt.Fprintf(&buf, "# Título sugerido\n")
	fmt.Fprintf(&buf, "%s\n\n", cfg.PRStyle.TitlePattern)
	fmt.Fprintf(&buf, "---\n\n")

	fmt.Fprintf(&buf, "## Contexto\n")
	fmt.Fprintf(&buf, "- ...\n\n")

	fmt.Fprintf(&buf, "## Mudancas\n")
	fmt.Fprintf(&buf, "- ...\n\n")

	fmt.Fprintf(&buf, "## Impacto")
	fmt.Fprintf(&buf, "- ...\n\n")

	fmt.Fprintf(&buf, "## Testes\n\n")
	fmt.Fprintf(&buf, "- ...\n\n")

	fmt.Fprintf(&buf, "## Checklist de cobertura\n")
	for _, item := range cfg.PRStyle.CoverageChecklist {
		fmt.Fprintf(&buf, "- [ ] %s\n", item)
	}

	fmt.Fprintf(&buf, "\n\nResumo das mudancas (para contexto):\n%s\n\n", diff.Summary)

	fmt.Fprintf(&buf, "Diff completo (truncado se necessário):\n```%s```\n", diff.RawDiff)

	return buf.String(), nil
}
