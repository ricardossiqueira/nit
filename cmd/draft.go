/*
Package cmd
*/
package cmd

import (
	"context"
	"fmt"

	"nit/internal/git"
	"nit/internal/llm"
	"nit/internal/output"

	"github.com/spf13/cobra"
)

var (
	baseBranch   string
	lang         string
	outputFormat string
	runID        string
	useLast      bool
)

// draftCmd represents the draft command
var draftCmd = &cobra.Command{
	Use:   "draft",
	Short: "Generate PR title & description from git diff",

	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := GetConfig()
		store := GetRunStore()

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
	},
}

func init() {
	draftCmd.Flags().StringVarP(&outputFormat, "format", "f", "pretty",
		"output format: pretty, json, commit, pr-title, pr-body, pr")
	draftCmd.Flags().StringVarP(&runID, "run-id", "r", "", "reuse draft from specific run ID")
	draftCmd.Flags().BoolVarP(&useLast, "last", "l", false, "reuse most recent draft")

	rootCmd.AddCommand(draftCmd)
	draftCmd.Flags().StringVar(&baseBranch, "base", "", "base branch for diff (overrides config)")
	draftCmd.Flags().StringVar(&lang, "lang", "", "force language for description (e.g. pt, en)")
}
