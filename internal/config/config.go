/*
Package config
*/
package config

import (
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	GitHub  GitHubConfig  `mapstruct:"github"`
	Model   ModelConfig   `mapstruct:"model"`
	PRStyle PRStyleConfig `mapstruct:"pr_style"`
	Review  ReviewConfig  `mapstruct:"review"`
	Prompt  PromptConfig  `mapstruct:"prompt"`
}

type GitHubConfig struct {
	DefaultBaseBranch string `mapstruct:"default_base_branch"`
	UseGHCLI          bool   `mapstruct:"use_gh_cli"`
}

type ModelConfig struct {
	Provider    string  `mapstruct:"provider"`
	ModelName   string  `mapstruct:"model_name"`
	Endpoint    string  `mapstruct:"endpoint"`
	MaxTokens   int     `mapstruct:"max_tokens"`
	Temperature float64 `mapstruct:"temperature"`
	Timeout     int     `mapstruct:"timeout_in_seconds"`
}

type PRStyleConfig struct {
	Language           string    `mapstruct:"language"`
	TitlePattern       string    `mapstruct:"title_pattern"`
	AllowedTypes       []string  `mapstruct:"allowed_types"`
	DescriptionSection []Section `mapstruct:"description_section"`
	CoverageChecklist  []string  `mapstruct:"coverage_check_list"`
}

type ReviewConfig struct {
	Focus           []string          `mapstruct:"focus"`
	Language        string            `mapstruct:"language"`
	StyleGuide      string            `mapstruct:"style_guide"`
	MaxDiffLines    int               `mapstruct:"max_diff_lines"`
	SeverityMapping map[string]string `mapstruct:"severity_mapping"`
}

type PromptConfig struct {
	SystemInstructions string   `mapstruct:"system_instructions"`
	ExtraRules         []string `mapstruct:"extra_rules"`
}

type Section struct {
	Name     string `mapstruct:"name"`
	Required bool   `mapstruct:"required"`
}

func DefaultConfig() *Config {
	return &Config{
		GitHub: GitHubConfig{
			DefaultBaseBranch: "master",
			UseGHCLI:          true,
		},
		Model: ModelConfig{
			Provider:    "ollama",
			ModelName:   "deepseek-r1:1.5b",
			Endpoint:    "http://localhost:11434/api/chat",
			MaxTokens:   2048,
			Temperature: 0.2,
			Timeout:     60,
		},
		PRStyle: PRStyleConfig{
			Language:     "en-US",
			TitlePattern: "[{type}] {scope}: {summary}",
			AllowedTypes: []string{"feat", "fix", "chore", "refactor", "docs"},
		},
		Review: ReviewConfig{
			Language: "golang",
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
