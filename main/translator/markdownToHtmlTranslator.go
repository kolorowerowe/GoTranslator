package translator

import (
	"bufio"
	"regexp"
	"strings"
)

func Translate(content string, fileName string) string {

	lines, _ := stringToLines(content)

	lines = addHeaders(lines)

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

func addHeaders(inputLines []string) (resultLines []string) {

	for _, line := range inputLines {
		match6, _ := regexp.MatchString(`^######.*`, line)
		match5, _ := regexp.MatchString(`^#####.*`, line)
		match4, _ := regexp.MatchString(`^####.*`, line)
		match3, _ := regexp.MatchString(`^###.*`, line)
		match2, _ := regexp.MatchString(`^##.*`, line)
		match1, _ := regexp.MatchString(`^#.*`, line)

		if match6 {
			resultLines = append(resultLines, "<h6>"+line[6:]+"</h6>")
		} else if match5 {
			resultLines = append(resultLines, "<h5>"+line[5:]+"</h5>")
		} else if match4 {
			resultLines = append(resultLines, "<h4>"+line[4:]+"</h4>")
		} else if match3 {
			resultLines = append(resultLines, "<h3>"+line[3:]+"</h3>")
		} else if match2 {
			resultLines = append(resultLines, "<h2>"+line[2:]+"</h2>")
		} else if match1 {
			resultLines = append(resultLines, "<h1>"+line[1:]+"</h1>")
		} else {
			resultLines = append(resultLines, line)
		}
	}
	return resultLines
}
