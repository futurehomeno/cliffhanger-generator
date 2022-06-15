package generator

import (
	"embed"
	"fmt"
	"os"
	"path"
)

type FileSet map[string][]*File

type File struct {
	Path        string
	Content     []byte
	Permissions os.FileMode
}

func (f *File) Generate(cfg Config, templates embed.FS) error {
	err := f.prepareContent(cfg, templates)
	if err != nil {
		return err
	}

	f.Path, err = renderString("path", f.Path, cfg)
	if err != nil {
		return err
	}

	if f.Permissions == 0 {
		f.Permissions = os.ModePerm
	}

	err = os.MkdirAll(path.Join(cfg.RootPath(), path.Dir(f.Path)), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory at %s: %w", path.Dir(f.Path), err)
	}

	err = os.WriteFile(path.Join(cfg.RootPath(), f.Path), f.Content, f.Permissions)
	if err != nil {
		return fmt.Errorf("failed to write file at %s: %w", f.Path, err)
	}

	return nil
}

func (f *File) prepareContent(cfg Config, templates embed.FS) error {
	var err error

	if len(f.Content) == 0 {
		templatePath, err := renderString("content_template_path", f.Path, cfg.TemplateConfig())
		if err != nil {
			return fmt.Errorf("failed to render content template path for file %s: %w", f.Path, err)
		}

		templateFile, err := templates.ReadFile(path.Join("templates", templatePath+".tmpl"))
		if err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to read content template for file %s: %w", templatePath, err)
		}

		f.Content = templateFile
	}

	if len(f.Content) == 0 {
		return nil
	}

	f.Content, err = renderBytes("content", f.Content, cfg)
	if err != nil {
		return fmt.Errorf("failed to render the content for file %s: %w", f.Path, err)
	}

	return nil
}
