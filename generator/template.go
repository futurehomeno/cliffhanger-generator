package generator

import (
	"bytes"
	"fmt"
	"text/template"
)

func renderString(name, content string, data interface{}) (string, error) {
	b, err := renderBytes(name, []byte(content), data)

	return string(b), err
}

func renderBytes(name string, content []byte, data interface{}) ([]byte, error) {
	tmpl, err := template.New(name).Parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("failed to parse the %s template: %w", name, err)
	}

	b := bytes.NewBuffer(nil)

	err = tmpl.Execute(b, data)
	if err != nil {
		return nil, fmt.Errorf("failed to execute the %s template: %w", name, err)
	}

	return b.Bytes(), nil
}
