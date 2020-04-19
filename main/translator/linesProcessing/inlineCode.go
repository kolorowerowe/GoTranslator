package linesProcessing

import "regexp"

func AddInlineCode(inputLines []string) (resultLines []string) {
	re := regexp.MustCompile("`.+?`") //match as small part of code as possible

	for _, line := range inputLines {
		converted := re.ReplaceAllStringFunc(line, func(n string) string {
			return `<code>` + n[1:len(n)-1] + `</code>`
		})
		resultLines = append(resultLines, converted)
	}
	return resultLines
}
