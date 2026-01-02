/*
Package cmd
*/
package cmd

import (
	"github.com/spf13/cobra"
	"nit/internal/app/draft"
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
		return draft.Generate(
			cfg,
			store,
			runID,
			useLast,
			outputFormat,
			lang,
			baseBranch,
		)
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
