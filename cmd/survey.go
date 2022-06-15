package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"

	"github.com/futurehomeno/cliffhanger-generator/configuration"
)

func Questions() []*survey.Question { //nolint:funlen
	return []*survey.Question{
		{
			Name: "domain",
			Prompt: &survey.Select{
				Message: "Select domain of the service",
				Options: []string{
					configuration.DomainEdge,
					configuration.DomainCore,
				},
				Default: configuration.DomainEdge,
				Help: fmt.Sprintf(
					"Select '%s' if your application is an optional playground application, and '%s' if it is an internal service.",
					configuration.DomainEdge, configuration.DomainCore,
				),
			},
		},
		{
			Name: "type",
			Prompt: &survey.Select{
				Message: "Select type of the service",
				Options: []string{
					configuration.TypeApp,
					configuration.TypeAdapter,
				},
				Default: configuration.TypeApp,
				Help: fmt.Sprintf(
					"Select '%s' if your application acts as an adapter and adds additional devices to the system, or '%s' otherwise.",
					configuration.TypeAdapter, configuration.TypeApp,
				),
			},
		},
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "Provide the name of the service",
				Help:    "The name of the service must be unique alphanumeric phrase that cosists of no more than three words.",
			},
			Transform: survey.Title,
			Validate:  validateString(configuration.ValidateName),
		},
		{
			Name: "description",
			Prompt: &survey.Input{
				Message: "Provide short description of the service",
			},
			Validate: validateString(configuration.ValidateDescription),
		},
		{
			Name: "path",
			Prompt: &survey.Input{
				Message: "Provide the path where the repository should be created",
				Default: "./",
			},
			Validate: validateString(configuration.ValidatePath),
		},
		{
			Name: "includeComments",
			Prompt: &survey.Confirm{
				Message: "Do you want to include help comments to the generated code?",
				Default: true,
			},
		},
		{
			Name: "includeTesting",
			Prompt: &survey.Confirm{
				Message: "Do you want to include testing suite to the generated code?",
				Default: false,
			},
		},
	}
}

func validateString(validator func(value string) error) survey.Validator {
	return func(value interface{}) error {
		stringValue, ok := value.(string)
		if !ok {
			return fmt.Errorf("value is not a string")
		}

		return validator(stringValue)
	}
}
