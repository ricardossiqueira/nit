/*
	Package cmd

Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"check/internal/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	configLoaded *config.Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "check",
	Short: "AI-powered PR assistant",
	Long: `check helps you create better pull requests with AI-generated
titles, descriptions and code reviews before pushing to github

Commands:
	draft     Generate PR title & description from git diff
	review    Generate code review comments
	pr        Create GitHub PR with generated content
	config    Manage configuration

Examples:
	check draft
	check draft --base develop
	check review --mode detailed`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .check.yml, $HOME/.check.yml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME")
		viper.SetConfigName(".check")

	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}

	var cfg config.Config
	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Error unmarshalling config: %v\n", err)
		os.Exit(1)

	}
	configLoaded = &cfg
}

func GetConfig() *config.Config {
	return configLoaded
}
