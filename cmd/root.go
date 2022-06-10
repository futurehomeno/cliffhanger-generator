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

	fmt.Println("Welcome to the Cliffhanger generator!")

	cfg := &configuration.Config{}

	err = survey.Ask(Questions(), cfg)
	if err != nil {
		fail(fmt.Errorf("failed to gather information: %w", err))
	}

	g := generator.New(configuration.TemplateSet, configuration.FileSet, configuration.CommandSet)

	err = g.Generate(cfg)
	if err != nil {
		fail(fmt.Errorf("failed to generate project: %w", err))
	}

	fmt.Println("Done!")
}

func fail(err error) {
	fmt.Printf("%s\n", err.Error())
	os.Exit(1)
}
