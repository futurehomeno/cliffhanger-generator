package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"

	"github.com/futurehomeno/cliffhanger-generator/configuration"
	"github.com/futurehomeno/cliffhanger-generator/generator"
)

func Execute() {
	var err error

	fmt.Println("Welcome to the Cliffhanger generator!") //nolint:forbidigo

	err = survey.Ask(Questions(), cfg)
	if err != nil {
		fail(fmt.Errorf("failed to gather information: %w", err))
	}

	g := generator.New(configuration.TemplateSet, configuration.FileSet, configuration.CommandSet)

	err = g.Generate(cfg)
	if err != nil {
		fail(fmt.Errorf("failed to generate project: %w", err))
	}

	fmt.Println("Done!") //nolint:forbidigo
}

func fail(err error) {
	fmt.Printf("%s\n", err.Error()) //nolint:forbidigo
	os.Exit(1)
}
