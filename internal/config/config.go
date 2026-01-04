/*
Package config
*/
package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	GitHub  GitHubConfig  `mapstructure:"github" yaml:"github"`
	Model   ModelConfig   `mapstructure:"model" yaml:"model"`
	PRStyle PRStyleConfig `mapstructure:"pr_style" yaml:"pr_style"`
	Review  ReviewConfig  `mapstructure:"review" yaml:"review"`
	Prompt  PromptConfig  `mapstructure:"prompt" yaml:"prompt"`
}

type GitHubConfig struct {
	DefaultBaseBranch string `mapstructure:"default_base_branch" yaml:"default_base_branch"`
	UseGHCLI          bool   `mapstructure:"use_gh_cli" yaml:"use_gh_cli"`
}

type ModelConfig struct {
	Provider       string  `mapstructure:"provider" yaml:"provider"`
	ModelName      string  `mapstructure:"model_name" yaml:"model_name"`
	Endpoint       string  `mapstructure:"endpoint" yaml:"endpoint"`
	MaxTokens      int     `mapstructure:"max_tokens" yaml:"max_tokens"`
	Temperature    float64 `mapstructure:"temperature" yaml:"temperature"`
	TimeoutSeconds int     `mapstructure:"timeout_seconds" yaml:"timeout_seconds"`
}

type PRStyleConfig struct {
	Language            string      `mapstructure:"language" yaml:"language"`
	Title               TitleConfig `mapstructure:"title" yaml:"title"`
	DescriptionSections []Section   `mapstructure:"description_sections" yaml:"description_sections"`
	CoverageChecklist   []string    `mapstructure:"coverage_checklist" yaml:"coverage_checklist"`
}

type ReviewConfig struct {
	Focus           []string          `mapstructure:"focus" yaml:"focus"`
	Language        string            `mapstructure:"language" yaml:"language"`
	PythonVersion   string            `mapstructure:"python_version" yaml:"python_version"`
	StyleGuide      string            `mapstructure:"style_guide" yaml:"style_guide"`
	MaxDiffLines    int               `mapstructure:"max_diff_lines" yaml:"max_diff_lines"`
	SeverityMapping map[string]string `mapstructure:"severity_mapping" yaml:"severity_mapping"`
}

type PromptConfig struct {
	SystemInstructions string   `mapstructure:"system_instructions" yaml:"system_instructions"`
	ExtraRules         []string `mapstructure:"extra_rules" yaml:"extra_rules"`
}

type TitleConfig struct {
	Pattern      string   `mapstructure:"pattern" yaml:"pattern"`
	AllowedTypes []string `mapstructure:"allowed_types" yaml:"allowed_types"`
	MaxLength    int      `mapstructure:"max_length" yaml:"max_length"`
}

type Section struct {
	Name     string `mapstructure:"name" yaml:"name"`
	Required bool   `mapstructure:"required" yaml:"required"`
}

func DefaultConfig() *Config {
	return &Config{
		GitHub: GitHubConfig{
			DefaultBaseBranch: "master",
			UseGHCLI:          true,
		},
		Model: ModelConfig{
			Provider:       "ollama",
			ModelName:      "deepseek-coder-v2",
			Endpoint:       "http://localhost:11434/v1/chat",
			MaxTokens:      2048,
			Temperature:    0.2,
			TimeoutSeconds: 60,
		},
		PRStyle: PRStyleConfig{
			Language: "pt-BR",
			Title: TitleConfig{
				Pattern:      "[{type}] {scope}: {summary}",
				AllowedTypes: []string{"feat", "fix", "chore", "refactor", "docs"},
				MaxLength:    72,
			},
			DescriptionSections: []Section{
				{Name: "Context", Required: true},
				{Name: "Changes", Required: true},
				{Name: "Impact", Required: false},
				{Name: "Tests", Required: true},
			},
			CoverageChecklist: []string{
				"Explains the reason for the change",
				"Lists the main changes",
				"Mentions affected modules",
				"Indicates impact or risks",
				"Describes how to test",
			},
		},
		Review: ReviewConfig{
			Focus:         []string{"bugs", "legibility", "style"},
			Language:      "en",
			PythonVersion: "3.11",
			StyleGuide:    "pep8",
			MaxDiffLines:  800,
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
