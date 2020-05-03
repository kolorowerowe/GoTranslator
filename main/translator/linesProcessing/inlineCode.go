package linesProcessing

import "regexp"

func AddInlineCode(inputLines []string) (resultLines []string) {
	re := regexp.MustCompile("`(.+?)`") //match as small part of code as possible

	for _, line := range inputLines {
		converted := re.ReplaceAllString(line, "<code>$1</code>")
		resultLines = append(resultLines, converted)
	}
	return resultLines
}
