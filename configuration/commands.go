package configuration

import (
	"github.com/futurehomeno/cliffhanger-generator/generator"
)

var CommandSet = map[string][]*generator.Command{
	"common": {
		{
			Cmd:  "go",
			Args: []string{"mod", "init", "github.com/futurehomeno/{{.RepositoryName}}"},
			Dir:  "{{.SourcePath}}",
		},
		{
			Cmd:  "go",
			Args: []string{"get", "github.com/futurehomeno/cliffhanger@latest"},
			Dir:  "{{.SourcePath}}",
		},
		{
			Cmd:  "go",
			Args: []string{"get", "github.com/futurehomeno/fimpgo@latest"},
			Dir:  "{{.SourcePath}}",
		},
		{
			Cmd:  "go",
			Args: []string{"get", "github.com/sirupsen/logrus@latest"},
			Dir:  "{{.SourcePath}}",
		},
		{
			Cmd:  "go",
			Args: []string{"get", "-u", "golang.org/x/sys"},
			Dir:  "{{.SourcePath}}",
		},
		{
			Cmd:  "go",
			Args: []string{"mod", "tidy"},
			Dir:  "{{.SourcePath}}",
		},
	},
}
