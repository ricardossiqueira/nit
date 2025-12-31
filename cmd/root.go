/*
	Package cmd

Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"nit/internal/config"
	"nit/internal/db"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	configLoaded *config.Config
	store        *db.Store
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nit",
	Short: "AI-powered PR assistant",
	Long: `nit helps you create better pull requests with AI-generated
titles, descriptions and code reviews before pushing to github

Commands:
	draft     Generate PR title & description from git diff
	review    Generate code review comments
	pr        Create GitHub PR with generated content
	config    Manage configuration

Examples:
	nit draft
	nit draft --base develop
	nit review --mode detailed`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is nit.yaml, $HOME/nit.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME")
		viper.SetConfigName("nit")

	}

	viper.AutomaticEnv()
	// TODO: log loaded file?
	_ = viper.ReadInConfig()

	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Printf("Error unmarshalling config: %v\n", err)
		os.Exit(1)
		// random comment
	}
	configLoaded = &cfg
}

func GetConfig() *config.Config {
	return configLoaded
}

func SetRunStore(s *db.Store) {
	store = s
}

func GetRunStore() *db.Store {
	return store
}
