/*
Package cmd
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"check/internal/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage check configuration",
}

// configCmd represents the config command
var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a default .check.yaml config file",

	RunE: func(cmd *cobra.Command, args []string) error {
		cfgPath := ".check.yaml"

		if _, err := os.Stat(cfgPath); err == nil {
			return fmt.Errorf("config file %s already exists", cfgPath)
		}

		cfg := config.DefaultConfig()
		if err := cfg.Save(cfgPath); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		fmt.Println("Created default config at", cfgPath)
		return nil
	},
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show effective configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		settings := viper.AllSettings()
		for k, v := range settings {
			fmt.Printf("%s: %#v\n", k, v)
		}
		return nil
	},
}

var configEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open config file in default editor",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgFile := viper.ConfigFileUsed()
		if cfgFile == "" {
			cfgFile = ".check.yaml"
			if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
				cfg := config.DefaultConfig()
				if err := cfg.Save(cfgFile); err != nil {
					return fmt.Errorf("failed to create default config: %w", err)
				}
			}
		}

		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "nano"
		}

		absPath, err := filepath.Abs(cfgFile)
		if err != nil {
			return err
		}

		cmdEditor := exec.Command(editor, absPath)
		cmdEditor.Stdin = os.Stdin
		cmdEditor.Stdout = os.Stdout
		cmdEditor.Stderr = os.Stderr

		return cmdEditor.Run()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configEditCmd)
}
