package gen

import (
	"cligen/templates"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"text/template"
)

type TemplateOptions struct {
	GolangVersion    string
	TemplateType     string
	OutputPathPrefix string
	AppName          string
}

type templatePath struct {
	inputPath  string
	outputPath string
}

func GetTemplates() ([]string, error) {

	root := "types"

	var folders []string

	dirs, err := fs.ReadDir(templates.FolderFS, root)

	for _, dir := range dirs {
		if dir.IsDir() {
			folders = append(folders, dir.Name())
		}
	}

	return folders, err
}

func RenderTemplate(options TemplateOptions) error {

	templatePaths, err := getTypeTemplatePaths(options.TemplateType)
	if err != nil {
		return err
	}

	slog.Info("Type Template paths", slog.Any("templatePaths", templatePaths))

	commonTemplatePaths, err := getCommonTemplatePaths()
	if err != nil {
		return err
	}

	slog.Info("Common Template paths", slog.Any("templatePaths", commonTemplatePaths))

	templatePaths = append(commonTemplatePaths, templatePaths...)

	slog.Info("All Template paths", slog.Any("templatePaths", templatePaths))

	for _, templatePath := range templatePaths {
		outputPath := filepath.Join(options.OutputPathPrefix, templatePath.outputPath)

		slog.Info("Template path", slog.Any("templatePath", templatePath))
		b, err := fs.ReadFile(templates.FolderFS, templatePath.inputPath)
		if err != nil {
			return err
		}

		t, err := template.New(outputPath).Parse(string(b))
		if err != nil {
			return err
		}

		if err := os.MkdirAll(filepath.Dir(outputPath), 0770); err != nil {
			return err
		}
		f, err := os.Create(outputPath)

		if err != nil {
			return err
		}
		defer f.Close()

		err = t.Execute(f, options)
		if err != nil {
			return err
		}
		slog.Info("Template rendered", slog.Any("templatePath", outputPath))
	}

	return nil
}

func getTemplatePaths(pathPrefix string) ([]templatePath, error) {
	var templatePaths []templatePath

	err := fs.WalkDir(templates.FolderFS, pathPrefix, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		// remove pathPrefix from path
		outputPath := path[len(pathPrefix)+1:]

		if len(outputPath) > 5 && outputPath[len(outputPath)-5:] == ".tmpl" {
			outputPath = outputPath[:len(outputPath)-5]
		}

		templatePaths = append(templatePaths, templatePath{inputPath: path, outputPath: outputPath})
		return nil
	})

	if err != nil {
		return []templatePath{}, err
	}

	return templatePaths, nil
}

func getCommonTemplatePaths() ([]templatePath, error) {
	pathPrefix := "common"

	return getTemplatePaths(pathPrefix)
}

func getTypeTemplatePaths(templateType string) ([]templatePath, error) {
	pathPrefix := "types/" + templateType

	return getTemplatePaths(pathPrefix)
}
