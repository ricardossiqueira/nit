/*
Package cmd
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"nit/internal/app/draft"
)

var (
	baseBranch string
	lang       string
	useLast    string
)

// draftCmd represents the draft command
var draftCmd = &cobra.Command{
	Use:   "draft",
	Short: "Generate commit messages or PR text from git diff",
}

var draftCommitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Generate a commit message from git diff",
	RunE: func(cmd *cobra.Command, args []string) error {
		if useLast != "" {
			if useLast != "commit" && useLast != "pr" {
				return fmt.Errorf("invalid value for --last on commit: %s (use commit)", useLast)
			}
			return draft.LastCommit(GetRunStore())
		}
		return draft.GenerateCommit(GetConfig(), GetRunStore(), lang, baseBranch)
	},
}

var draftPRCmd = &cobra.Command{
	Use:   "pr",
	Short: "Generate PR title and body from git diff",
	RunE: func(cmd *cobra.Command, args []string) error {
		switch useLast {
		case "title":
			return draft.LastTitle(GetRunStore())
		case "body":
			return draft.LastBody(GetRunStore())
		case "pr", "":
			if useLast == "pr" {
				return draft.LastPR(GetRunStore())
			}
		default:
			return fmt.Errorf("invalid value for --last: %s (use commit|title|body|pr)", useLast)
		}
		return draft.GeneratePR(GetConfig(), GetRunStore(), lang, baseBranch)
	},
}

func init() {
	rootCmd.AddCommand(draftCmd)
	draftCmd.PersistentFlags().StringVar(&baseBranch, "base", "", "base branch for diff (overrides config)")
	draftCmd.PersistentFlags().StringVar(&lang, "lang", "", "force language for description (e.g. pt, en)")
	draftCmd.PersistentFlags().StringVarP(&useLast, "last", "l", "", "reuse most recent generated draft (commit|title|body|pr)")
	draftCmd.AddCommand(draftCommitCmd)
	draftCmd.AddCommand(draftPRCmd)
}
