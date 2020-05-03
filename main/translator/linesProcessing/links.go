package linesProcessing

import (
	"regexp"
)

//TODO 03-05-2020: add links with title, add references style link
func AddLinks(inputLines []string) (resultLines []string) {
	re := regexp.MustCompile(`\[([^\[\]()]+)]\(([\w\d.\/:?=#]+)\)`) //match as small part of code as possible

	for _, line := range inputLines {
		converted := re.ReplaceAllString(line, "<a href=\"$2\">$1</a>")
		resultLines = append(resultLines, converted)
	}
	return resultLines
}
