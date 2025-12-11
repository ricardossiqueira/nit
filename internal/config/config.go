/*
Package config
*/
package config

import (
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	GitHub  GitHubConfig  `yaml:"github"`
	Model   ModelConfig   `yaml:"model"`
	PRStyle PRStyleConfig `yaml:"pr_style"`
	Review  ReviewConfig  `yaml:"review"`
	Prompt  PromptConfig  `yaml:"prompt"`
}

type GitHubConfig struct {
	DefaultBaseBranch string `yaml:"default_base_branch"`
	UseGHCLI          bool   `yaml:"use_gh_cli"`
}

type ModelConfig struct {
	Provider    string  `yaml:"provider"`
	ModelName   string  `yaml:"model_name"`
	Endpoint    string  `yaml:"endpoint"`
	MaxTokens   int     `yaml:"max_tokens"`
	Temperature float64 `yaml:"temperature"`
	Timeout     int     `yaml:"timeout_in_seconds"`
}

type PRStyleConfig struct {
	Language           string    `yaml:"language"`
	TitlePattern       string    `yaml:"title_pattern"`
	AllowedTypes       []string  `yaml:"allowed_types"`
	DescriptionSection []Section `yaml:"description_section"`
	CoverageChecklist  []string  `yaml:"coverage_checklist"`
}

type ReviewConfig struct {
	Focus           []string          `yaml:"focus"`
	PythonVersion   string            `yaml:"python_version"`
	StyleGuide      string            `yaml:"style_guide"`
	MaxDiffLines    int               `yaml:"max_diff_lines"`
	SeverityMapping map[string]string `yaml:"severity_mapping"`
}

type PromptConfig struct {
	SystemInstructions string   `yaml:"system_instructions"`
	ExtraRules         []string `yaml:"extra_rules"`
}

type Section struct {
	Name     string `yaml:"name"`
	Required bool   `yaml:"required"`
}

func DefaultConfig() *Config {
	return &Config{
		GitHub: GitHubConfig{
			DefaultBaseBranch: "master",
			UseGHCLI:          true,
		},
		Model: ModelConfig{
			Provider:    "ollama",
			ModelName:   "deepseek-coder",
			Endpoint:    "http://localhost:11434/v1/chat",
			MaxTokens:   2048,
			Temperature: 0.2,
			Timeout:     60,
		},
		PRStyle: PRStyleConfig{
			Language:     "pt-BR",
			TitlePattern: "[{type}] {scope}: {summary}",
			AllowedTypes: []string{"feat", "fix", "chore", "refactor", "docs"},
		},
	}
}

func (c *Config) Save(filename string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0o644)
}
