package utils

import (
	"regexp"
	"strings"
)

func MarkdownToHtmlName(inputName string) string {

	outputName := strings.Replace(inputName, ".MD", ".html", 1)
	outputName = strings.Replace(outputName, ".md", ".html", 1)

	regexExtractAfterLastBackslash := regexp.MustCompile(`(.*\/)*(.*)`)

	outputName = regexExtractAfterLastBackslash.ReplaceAllString(outputName, "$2")

	return outputName
}

func CutStringFromBeginning(text string, fromBeginnning int) string {
	return CutStringFromBoth(text, fromBeginnning, 0)
}

func CutStringFromBoth(text string, fromBeginning int, fromEnd int) string {
	return text[fromBeginning : len(text)-fromEnd]
}
