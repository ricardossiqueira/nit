/*
Package cmd
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"nit/internal/llm"

	"github.com/spf13/cobra"
)

// llmCmd represents the llm command
var llmCmd = &cobra.Command{
	Use:   "llm",
	Short: "Ollama connection health check.",
}

var llmStatus = &cobra.Command{
	Use:   "list",
	Short: "List Ollama available models.",
	RunE: func(cmd *cobra.Command, args []string) error {
		llmTags, err := llm.GetTags()
		if err != nil {
			return err
		}

		llmTags.PrintList()
		return nil
	},
}

var llmPing = &cobra.Command{
	Use:   "ping",
	Short: "Check Ollama connection.",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := llm.Ping()
		if err != nil {
			return err
		}
		fmt.Println("OK.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(llmCmd)
	llmCmd.AddCommand(llmStatus)
	llmCmd.AddCommand(llmPing)
}
