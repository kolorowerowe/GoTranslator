package translator

import (
	"./linesProcessing"
	"bufio"
	"strings"
)

func Translate(content string, fileName string) string {

	lines, _ := stringToLines(content)

	lines = linesProcessing.AddHeaders(lines)
	lines = linesProcessing.AddInlineCode(lines)

	content = linesToString(lines)
	content = wrapInHtml(content, fileName)
	return content
}

func stringToLines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func linesToString(lines []string) string {
	return strings.Join(lines, "\n")
}

func wrapInHtml(content string, fileName string) string {
	return `<html lang="en"><head><meta charset="utf-8"><title>` + fileName + `</title></head><body>` + content + `</body></html>`
}
