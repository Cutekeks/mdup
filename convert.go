package main

import (
	"fmt"
	"mdup/internal"
	"os"
)

func convert(inputFile string, outputFile string) error {
	// Check if input is provided
	if inputFile == "" {
		return fmt.Errorf("input file is required")
	}
	// Read Markdown file

	content, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	parsedContent := internal.ParseMD(string(content))
	htmlOutput := internal.ConvertHTML(parsedContent)

	if outputFile != "" {
		if err := os.WriteFile(outputFile, []byte(htmlOutput), 0644); err != nil {
			return fmt.Errorf("error writing output file: %v", err)
		}
		fmt.Printf("converted HTML written to %s\n", outputFile)
	} else {
		fmt.Println(htmlOutput)
	}
	return nil
}
