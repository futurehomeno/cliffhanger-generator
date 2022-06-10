package configuration

import (
	"embed"
)

//go:embed templates/* templates/**/*
var TemplateSet embed.FS
