package generator

import (
	"embed"
	"fmt"
	"os"
)

type Generator interface {
	Generate(cfg Config) error
}

func New(templateSet embed.FS, fileSet FileSet, commandSet CommandSet) Generator {
	return &generator{
		templateSet: templateSet,
		fileSet:     fileSet,
		commandSet:  commandSet,
	}
}

type generator struct {
	cfg         *Config
	templateSet embed.FS
	fileSet     FileSet
	commandSet  CommandSet
}

func (g *generator) Generate(cfg Config) error {
	err := cfg.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate configuration: %w", err)
	}

	err = cfg.Configure()
	if err != nil {
		return fmt.Errorf("failed to complete configuration: %w", err)
	}

	err = g.prepare(cfg)
	if err != nil {
		return fmt.Errorf("failed to prepare root directory: %w", err)
	}

	err = g.generateFiles(cfg)
	if err != nil {
		return fmt.Errorf("failed to generate files: %w", err)
	}

	err = g.executeCommands(cfg)
	if err != nil {
		return fmt.Errorf("failed to execute commands: %w", err)
	}

	return nil
}

func (g *generator) prepare(cfg Config) error {
	stat, err := os.Stat(cfg.RootPath())
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to check root directory %s: %w", cfg.RootPath(), err)
	} else if stat != nil {
		return fmt.Errorf("provided path %s is not empty, delete it before continuing", cfg.RootPath())
	}

	err = os.MkdirAll(cfg.RootPath(), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create root directory: %w", err)
	}

	return nil
}

func (g *generator) generateFiles(cfg Config) error {
	var files []*File

	for _, name := range cfg.FileSets() {
		fs, _ := g.fileSet[name]
		files = append(files, fs...)
	}

	for _, f := range files {
		err := f.Generate(cfg, g.templateSet)
		if err != nil {
			return fmt.Errorf("failed to generate file %s: %w", f.Path, err)
		}
	}

	return nil
}

func (g *generator) executeCommands(cfg Config) error {
	var commands []*Command

	for _, name := range cfg.CommandSets() {
		cs, _ := g.commandSet[name]
		commands = append(commands, cs...)
	}

	for _, c := range commands {
		err := c.Execute(cfg)
		if err != nil {
			return fmt.Errorf("failed to execute command %s: %w", c.String(), err)
		}
	}

	return nil
}
