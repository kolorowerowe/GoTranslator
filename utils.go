package main

import "strings"

func markdownToHtmlName(inputName string) string {

	outputName := strings.Replace(inputName, ".MD", ".html", 1)
	outputName = strings.Replace(outputName, ".md", ".html", 1)

	return outputName
}
