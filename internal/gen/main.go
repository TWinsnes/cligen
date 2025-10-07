package gen

type TemplateOptions struct {
	GolangVersion    string
	TemplateType     string
	OutputPathPrefix string
	AppName          string
	ModuleName       string
	HomebrewRepo     string
	HomebrewEnabled  bool
	HomebrewUsername string
	Features         []Feature
}

type templatePath struct {
	inputPath  string
	outputPath string
	isTemplate bool
}
