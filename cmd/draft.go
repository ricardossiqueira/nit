/*
Package cmd
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"check/internal/git"

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

		fmt.Println(diffCtx)
		// TODO: 	continue here
		return nil
	},
}

func init() {
	rootCmd.AddCommand(draftCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// draftCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// draftCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
