/*Package draft
 */
package draft

import (
	"context"
	"fmt"
	"nit/internal/config"
	"nit/internal/db"
	"nit/internal/git"
	"nit/internal/llm"
	"nit/internal/output"
)

func Generate(cfg *config.Config, store *db.Store, runID string, useLast bool, outputFormat string, lang string, baseBranch string) error {
	if runID != "" {
		draft, err := store.GetDraftByID(context.TODO(), runID)
		if err != nil {
			return err
		}
		return output.PrintDraft(draft, output.OutputFormat(outputFormat))
	}

	if useLast {
		draft, err := store.GetLastDraft(context.TODO())
		if err != nil {
			return err
		}
		return output.PrintDraft(draft, output.OutputFormat(outputFormat))
	}

	if cfg == nil {
		return fmt.Errorf("config not loaded")
	}

	if baseBranch == "" {
		baseBranch = cfg.GitHub.DefaultBaseBranch
	}

	diffCtx, err := git.ParseDiff(baseBranch, cfg.Review.MaxDiffLines)
	if err != nil {
		return fmt.Errorf("failed to get diff: %w", err)
	}

	currentBranchCtx, err := git.GetBranch()
	if err != nil {
		return fmt.Errorf("failed to get current branch: %w", err)
	}

	prompt, err := llm.BuildDraftPrompt(cfg, diffCtx, lang)
	if err != nil {
		return fmt.Errorf("failed to build prompt: %w", err)
	}

	resp, err := llm.Generate(context.TODO(), store, cfg.Model, prompt, currentBranchCtx.Name)
	if err != nil {
		return fmt.Errorf("llm generation failed: %w", err)
	}

	if err := store.SaveRun(context.TODO(), resp); err != nil {
		return fmt.Errorf("failed saving response to the db: %w", err)
	}

	output.PrintDraft(&resp.Response, output.OutputFormat(outputFormat))

	return nil
}
