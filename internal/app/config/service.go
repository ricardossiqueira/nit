package configapp

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"nit/internal/config"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// Init creates a default config file at cfgPath if it does not already exist.
func Init(cfgPath string) error {
	if _, err := os.Stat(cfgPath); err == nil {
		return fmt.Errorf("config file %s already exists", cfgPath)
	}

	cfg := config.DefaultConfig()
	if err := cfg.Save(cfgPath); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	fmt.Println("Created default config at", cfgPath)
	return nil
}

// Show prints either the model name or the full configuration using the provided viper instance.
func Show(v *viper.Viper, cfg *config.Config, showModelOnly bool) error {
	if showModelOnly {
		if cfg == nil {
			return fmt.Errorf("config not loaded")
		}
		fmt.Print(cfg.Model.ModelName)
		return nil
	}

	allSettings := v.AllSettings()
	yamlBytes, err := yaml.Marshal(allSettings)
	if err != nil {
		return err
	}
	fmt.Println(string(yamlBytes))
	return nil
}

// Edit opens the config file in the default editor, creating it if needed.
func Edit(v *viper.Viper) error {
	cfgFile := v.ConfigFileUsed()
	if cfgFile == "" {
		cfgFile = "nit.yaml"
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
}
