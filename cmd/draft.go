/*
Package cmd
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"nit/internal/git"
	"nit/internal/llm"
	"nit/internal/output"

	"github.com/spf13/cobra"
)

var (
	baseBranch string
	lang       string
)

// draftCmd represents the draft command
var draftCmd = &cobra.Command{
	Use:   "draft",
	Short: "Generate PR title & description from git diff",

	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := GetConfig()

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

		prompt, err := llm.BuildDraftPrompt(cfg, diffCtx, lang)
		if err != nil {
			return fmt.Errorf("failed to build prompt: %w", err)
		}

		resp, err := llm.Generate(cfg.Model, prompt)
		if err != nil {
			return fmt.Errorf("llm generation failed: %w", err)
		}

		if err := output.PrintDraft(resp); err != nil {
			return fmt.Errorf("failed to render draft: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(draftCmd)

	draftCmd.Flags().StringVar(&baseBranch, "base", "", "base branch for diff (overrides config)")
	draftCmd.Flags().StringVar(&lang, "lang", "", "force language for description (e.g. pt, en)")
}
