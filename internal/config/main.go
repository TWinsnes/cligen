package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Config represents the configuration of cligen.
type Config struct {
	HomebrewConfig HomebrewConfig `yaml:"homebrew"`
}

// HomebrewConfig represents the configuration of the Homebrew tap.
type HomebrewConfig struct {
	Enabled        bool   `yaml:"enabled"`
	Repo           string `yaml:"repo"`
	GithubUsername string `yaml:"github_username"`
}

const (
	Filename = ".cligen.yaml"
	Dir      = "~/.config"
)

// NewConfig creates a new Config instance with default values
func NewConfig() *Config {
	return &Config{
		HomebrewConfig: HomebrewConfig{
			Enabled:        true,
			Repo:           "",
			GithubUsername: "",
		},
	}
}

// LoadConfig loads the config file from the default location and returns a Config instance.
// If the config file doesn't exist, it returns a Config instance with default values.
// If the config file exists but cannot be read, it returns an error.
// If the config file exists but cannot be parsed, it returns an error.
func LoadConfig() (*Config, error) {

	fullPath, err := GetConfigPath()

	if err != nil {
		return nil, fmt.Errorf("failed to get config path: %w", err)
	}

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return NewConfig(), nil
	}

	// Read the config file
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", fullPath, err)
	}

	// Parse YAML content
	config := NewConfig()
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("failed to parse config file %s: %w", fullPath, err)
	}

	return config, nil
}

func (c *Config) SaveConfig() error {

	fullPath, err := GetConfigPath()
	if err != nil {
		return fmt.Errorf("failed to get config path: %w", err)
	}

	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	err = os.WriteFile(fullPath, data, 0644)

	return err
}

func GetConfigPath() (string, error) {
	configDir := Dir

	// Expand the home directory if the path starts with ~
	if strings.HasPrefix(configDir, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get user home directory: %w", err)
		}
		configDir = strings.Replace(configDir, "~", homeDir, 1)
	}

	return filepath.Join(configDir, Filename), nil
}
