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

func GenerateCommit(cfg *config.Config, store *db.Store, lang string, baseBranch string) error {
	if store == nil {
		return fmt.Errorf("store not initialized")
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

	prompt, err := llm.BuildCommitPrompt(cfg, diffCtx, lang)
	if err != nil {
		return fmt.Errorf("failed to build commit prompt: %w", err)
	}

	resp, err := llm.Generate(context.Background(), store, cfg.Model, prompt, currentBranchCtx.Name, "commit")
	if err != nil {
		return fmt.Errorf("llm generation failed: %w", err)
	}

	return output.PrintCommit(resp.Response)
}

func LastCommit(store *db.Store) error {
	if store == nil {
		return fmt.Errorf("store not initialized")
	}
	resp, err := store.GetLastByType(context.Background(), "commit")
	if err != nil {
		return err
	}
	return output.PrintCommit(resp)
}

func GeneratePR(cfg *config.Config, store *db.Store, lang string, baseBranch string) error {
	if store == nil {
		return fmt.Errorf("store not initialized")
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

	titlePrompt, err := llm.BuildPRTitlePrompt(cfg, diffCtx, lang)
	if err != nil {
		return fmt.Errorf("failed to build pr title prompt: %w", err)
	}
	titleResp, err := llm.Generate(context.Background(), store, cfg.Model, titlePrompt, currentBranchCtx.Name, "title")
	if err != nil {
		return fmt.Errorf("llm title generation failed: %w", err)
	}

	bodyPrompt, err := llm.BuildPRBodyPrompt(cfg, diffCtx, lang)
	if err != nil {
		return fmt.Errorf("failed to build pr body prompt: %w", err)
	}
	bodyResp, err := llm.Generate(context.Background(), store, cfg.Model, bodyPrompt, currentBranchCtx.Name, "body")
	if err != nil {
		return fmt.Errorf("llm body generation failed: %w", err)
	}

	return output.PrintPR(titleResp.Response, bodyResp.Response)
}

func LastPR(store *db.Store) error {
	if store == nil {
		return fmt.Errorf("store not initialized")
	}

	title, err := store.GetLastByType(context.Background(), "title")
	if err != nil {
		return err
	}

	body, err := store.GetLastByType(context.Background(), "body")
	if err != nil {
		return err
	}

	return output.PrintPR(title, body)
}

func LastTitle(store *db.Store) error {
	if store == nil {
		return fmt.Errorf("store not initialized")
	}
	title, err := store.GetLastByType(context.Background(), "title")
	if err != nil {
		return err
	}
	return output.PrintTitle(title)
}

func LastBody(store *db.Store) error {
	if store == nil {
		return fmt.Errorf("store not initialized")
	}
	body, err := store.GetLastByType(context.Background(), "body")
	if err != nil {
		return err
	}
	return output.PrintBody(body)
}
