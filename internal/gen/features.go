package gen

type Feature interface {
	// GetName returns the name of the feature.
	GetName() string

	// GetPath returns the path to the feature's template directory.'
	GetPath() string

	// HasTemplateDir returns true if the feature has a template directory with files to process.
	HasTemplateDir() bool

	// UpdateTemplateOptions updates the TemplateOptions with the feature's template options.'
	UpdateTemplateOptions(options *TemplateOptions) error
}
