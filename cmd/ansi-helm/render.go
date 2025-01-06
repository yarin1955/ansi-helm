package main

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"
)

// RenderTemplates renders templates in the specified templatesPath using the values file.
func RenderTemplates(templatesPath string, valuesPath string) (string, error) {
	// Read the values.yaml file
	valuesData, err := os.ReadFile(valuesPath)
	if err != nil {
		return "", err
	}

	// Prepare a buffer to store rendered output
	var renderedOutput bytes.Buffer

	// Walk through all template files in the templates directory
	err = filepath.Walk(templatesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Process only YAML files
		if !info.IsDir() && filepath.Ext(path) == ".yaml" {
			templateContent, readErr := os.ReadFile(path)
			if readErr != nil {
				return readErr
			}

			// Parse and execute the template
			tmpl, parseErr := template.New(filepath.Base(path)).Parse(string(templateContent))
			if parseErr != nil {
				return parseErr
			}

			execErr := tmpl.Execute(&renderedOutput, map[string]interface{}{
				"Values": string(valuesData),
			})
			if execErr != nil {
				return execErr
			}

			// Add a separator for multiple templates
			renderedOutput.WriteString("\n---\n")
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	return renderedOutput.String(), nil
}
