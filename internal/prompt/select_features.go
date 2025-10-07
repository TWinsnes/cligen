package prompt

import (
	"github.com/charmbracelet/huh"
)

// runSelectFeatures runs the feature selection prompt and returns the features selected by the user
func runSelectFeatures(features map[string]FeaturePrompt) (map[string]FeaturePrompt, error) {

	var selected []string

	var promptOptions []huh.Option[string]

	for _, feature := range features {
		promptOptions = append(promptOptions, huh.Option[string]{Key: feature.GetName(), Value: feature.GetName()}.Selected(feature.IsDefaultSelected()))
	}

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Features").
				Description("Select which features you would like included").
				Value(&selected).
				Options(promptOptions...),
		),
	).Run()

	if err != nil {
		return map[string]FeaturePrompt{}, err
	}

	selectedFeatures := make(map[string]FeaturePrompt)

	for _, selection := range selected {
		selectedFeatures[selection] = features[selection]
	}

	return selectedFeatures, nil
}
