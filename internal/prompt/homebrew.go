package prompt

import (
	"fmt"
	"log/slog"
	"regexp"

	"github.com/charmbracelet/huh"
	"github.com/twinsnes/cligen/internal/config"
	"github.com/twinsnes/cligen/internal/gen"
)

type HomebrewConfig struct {
	name            string
	path            string
	hasTemplateDir  bool
	defaultSelected bool
	enabled         bool
	repo            string
	githubUsername  string
}

func NewHomebrewConfig(conf *config.Config) *HomebrewConfig {
	return &HomebrewConfig{
		name:            "Homebrew",
		path:            "",
		hasTemplateDir:  false,
		defaultSelected: true,
		enabled:         conf.HomebrewConfig.Enabled,
		repo:            conf.HomebrewConfig.Repo,
		githubUsername:  conf.HomebrewConfig.GithubUsername,
	}
}

func (h *HomebrewConfig) GetName() string {
	return h.name
}

func (h *HomebrewConfig) GetPath() string {
	return h.path
}

func (h *HomebrewConfig) HasTemplateDir() bool {
	return h.hasTemplateDir
}

func (h *HomebrewConfig) IsDefaultSelected() bool {
	return h.defaultSelected
}

// RunPrompt runs the homebrew config prompt and updates the config based on the user input
func (h *HomebrewConfig) RunPrompt() error {
	var repo string
	var username string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Homebrew repo").
				Description("The name of the homebrew repo.").
				Placeholder(h.repo).
				Validate(validateHomebrewRepo).
				Value(&repo),
			huh.NewInput().
				Title("Homebrew username").
				Description("The username of the homebrew repo owner.").
				Placeholder(h.githubUsername).
				Validate(validateGithubUsername).
				Value(&username),
		),
	)

	err := form.Run()
	if err != nil {
		return err
	}

	if repo != "" {
		h.repo = repo
	}
	if username != "" {
		h.githubUsername = username
	}

	h.enabled = true

	return nil
}

func (h *HomebrewConfig) UpdateTemplateOptions(options *gen.TemplateOptions) error {
	options.HomebrewEnabled = h.enabled
	options.HomebrewRepo = h.repo
	options.HomebrewUsername = h.githubUsername

	return nil
}

func validateHomebrewRepo(repo string) error {
	matched, err := regexp.Match("^[a-zA-Z0-9_-]*$", []byte(repo))

	if err != nil {
		slog.Debug("error matching repo",
			slog.Any("error", err),
		)

		return fmt.Errorf("internal error when validating homebrew repo, enable debug logging for more details")
	}

	if !matched {
		return fmt.Errorf("homebrew repo must only contain letters, numbers, hyphens or underscores")
	}
	return nil
}

func validateGithubUsername(username string) error {
	matched, err := regexp.Match("^[a-zA-Z0-9_-]*$", []byte(username))
	if err != nil {
		slog.Debug("error matching github username",
			slog.Any("error", err),
		)
		return fmt.Errorf("internal error when validating github username, enable debug logging for more details")
	}
	if !matched {
		return fmt.Errorf("github username must only contain letters, numbers, hyphens or underscores")
	}
	return nil
}
