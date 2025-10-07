package prompt

import (
	"github.com/twinsnes/cligen/internal/config"
	"github.com/twinsnes/cligen/internal/gen"
)

type DocsFeature struct {
	conf            *config.Config
	name            string
	path            string
	hasTemplateDir  bool
	defaultSelected bool
	enabled         bool
}

func NewDocsFeature(conf *config.Config) *DocsFeature {
	return &DocsFeature{
		conf:            conf,
		name:            "Docs",
		path:            "features/pages",
		hasTemplateDir:  true,
		defaultSelected: true,
		enabled:         false,
	}
}

func (d *DocsFeature) GetName() string {
	return d.name
}

func (d *DocsFeature) GetPath() string {
	return d.path
}

func (d *DocsFeature) IsEnabled() bool {
	return d.enabled
}

func (d *DocsFeature) HasTemplateDir() bool {
	return d.hasTemplateDir
}

func (d *DocsFeature) IsDefaultSelected() bool {
	return d.defaultSelected
}

func (d *DocsFeature) RunPrompt() error {
	d.enabled = true
	return nil
}

func (d *DocsFeature) UpdateTemplateOptions(options *gen.TemplateOptions) error {

	options.DocsEnabled = d.enabled
	return nil
}
