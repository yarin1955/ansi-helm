package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Default paths
	defaultTemplatesPath := "./templates"
	defaultValuesPath := "./values.yaml"

	// Ensure the templates folder exists
	if _, err := os.Stat(defaultTemplatesPath); os.IsNotExist(err) {
		log.Fatalf("Templates path does not exist: %s", defaultTemplatesPath)
	}

	// Use the default values.yaml if not provided via arguments
	valuesPath := defaultValuesPath
	if len(os.Args) > 2 {
		valuesPath = os.Args[2]
	} else {
		fmt.Printf("Using default values file: %s\n", valuesPath)
	}

	// Ensure the values.yaml file exists
	if _, err := os.Stat(valuesPath); os.IsNotExist(err) {
		log.Fatalf("Values file does not exist: %s", valuesPath)
	}

	// Call the RenderTemplates function
	renderedPlaybook, err := RenderTemplates(defaultTemplatesPath, valuesPath)
	if err != nil {
		log.Fatalf("Failed to render templates: %v", err)
	}

	// Print the rendered playbook to stdout
	fmt.Println(renderedPlaybook)
}
