package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	convertCmd := flag.NewFlagSet("convert", flag.ExitOnError)

	inputFile := convertCmd.String("input", "", "Path to the Markdown file to convert")
	outputFile := convertCmd.String("output", "", "Path to save the converted HTML file")
	//check if command is provided
	if len(os.Args) < 2 {
		fmt.Println("expected 'convert' command")
		os.Exit(1)
	}
	//Parse the command
	switch os.Args[1] {
	case "convert":
		err := convertCmd.Parse(os.Args[2:])
		if err != nil {
			return
		}
	default:
		fmt.Println("Unknown command. Use 'convert'.")
		os.Exit(1)
	}
	if err := convert(*inputFile, *outputFile); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
