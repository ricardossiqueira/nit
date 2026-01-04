/*
Package cmd
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	configapp "nit/internal/app/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var showModelOnly bool

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage nit configuration",
}

// configCmd represents the config command
var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a default nit.yaml config file",

	RunE: func(cmd *cobra.Command, args []string) error {
		return configapp.Init("nit.yaml")
	},
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show effective configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		return configapp.Show(viper.GetViper(), GetConfig(), showModelOnly)
	},
}

var configEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open config file in default editor",
	RunE: func(cmd *cobra.Command, args []string) error {
		return configapp.Edit(viper.GetViper())
	},
}

func init() {
	configShowCmd.Flags().BoolVarP(&showModelOnly, "model", "m", false, "Return current model from the config file")

	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configEditCmd)

}
