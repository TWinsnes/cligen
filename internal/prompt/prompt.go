package prompt

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"

	"github.com/charmbracelet/huh"
	"github.com/twinsnes/cligen/internal/config"
	"github.com/twinsnes/cligen/internal/gen"
	"github.com/twinsnes/cligen/internal/git"
)

type UserPrompt struct {
	conf              *config.Config
	AvailableFeatures map[string]FeaturePrompt
}

type FeaturePrompt interface {
	RunPrompt() error
	IsDefaultSelected() bool

	gen.Feature
}

func NewUserPrompt(conf *config.Config) *UserPrompt {

	homebrewFeature := NewHomebrewFeature(conf)
	docsFeature := NewDocsFeature(conf)

	return &UserPrompt{
		conf: conf,
		AvailableFeatures: map[string]FeaturePrompt{
			homebrewFeature.GetName(): homebrewFeature,
			docsFeature.GetName():     docsFeature,
		},
	}
}

func (p *UserPrompt) Run() (gen.TemplateOptions, error) {

	options, err := getBaseSettings(p.conf)

	if err != nil {
		return gen.TemplateOptions{}, err
	}

	selectedFeatures, err := runSelectFeatures(p.AvailableFeatures)

	if err != nil {
		return gen.TemplateOptions{}, err
	}

	var features []gen.Feature
	for _, feature := range selectedFeatures {
		err := feature.RunPrompt()
		if err != nil {
			return gen.TemplateOptions{}, err
		}

		features = append(features, feature)
	}

	options.Features = features

	return options, nil
}

func getBaseSettings(conf *config.Config) (gen.TemplateOptions, error) {
	var appName string
	var golangVersion string
	var goModuleName string
	var templateType string

	defaultAppName := getDefaultAppName()
	defaultModule := getDefaultModule()

	templateTypeOptions, err := getTemplateTypeOptions()
	if err != nil {
		return gen.TemplateOptions{}, err
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("App name").
				Description("The name of the CLI app.").
				Placeholder(defaultAppName).
				Validate(validateAppName).
				Value(&appName),
			huh.NewInput().
				Title("Go Module name").
				Description("The name of the Go module.").
				Placeholder(defaultModule).
				Value(&goModuleName),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Golang version").
				Description("The version of Go to use.").
				Options(
					huh.NewOption("1.25", "1.25"),
					huh.NewOption("1.24", "1.24"),
					huh.NewOption("1.23", "1.23"),
				).
				Value(&golangVersion),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Template").
				Description("The template to generate the app from.").
				Options(templateTypeOptions...).
				Value(&templateType),
		),
	)

	err = form.Run()
	if err != nil {
		return gen.TemplateOptions{}, err
	}

	if appName == "" {
		appName = defaultAppName
	}

	if goModuleName == "" {
		goModuleName = defaultModule
	}

	return gen.TemplateOptions{
		AppName:       appName,
		GolangVersion: golangVersion,
		ModuleName:    goModuleName,
		TemplateType:  templateType,
		DryRun:        false,
	}, nil
}

func validateAppName(appName string) error {
	matched, err := regexp.Match("^[a-zA-Z0-9_-]*$", []byte(appName))

	if err != nil {
		slog.Debug("error matching app name",
			slog.Any("error", err),
		)

		return fmt.Errorf("internal error when validating app name, enable debug logging for more details")
	}

	if !matched {
		return fmt.Errorf("app name must only contain letters, numbers, hyphens or underscores")
	}
	return nil
}

func getDefaultAppName() string {
	cwd, err := os.Getwd()
	if err == nil {
		return filepath.Base(cwd)
	}
	return "cligen"
}

func getDefaultModule() string {
	m := git.ModuleFromGit()

	if m == "" {
		m = "github.com/twinsnes/cligen"
	}

	return m
}

func getTemplateTypeOptions() ([]huh.Option[string], error) {
	templates, err := gen.ListTemplates()
	if err != nil {
		return []huh.Option[string]{}, err
	}
	var options []huh.Option[string]
	for _, template := range templates {
		options = append(options, huh.Option[string]{Key: template, Value: template})
	}
	return options, nil
}
