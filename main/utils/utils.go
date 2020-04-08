package utils

import (
	"strings"
)

func MarkdownToHtmlName(inputName string) string {

	outputName := strings.Replace(inputName, ".MD", ".html", 1)
	outputName = strings.Replace(outputName, ".md", ".html", 1)

	return outputName
}
