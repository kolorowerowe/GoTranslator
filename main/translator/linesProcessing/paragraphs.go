package linesProcessing

import "regexp"

func AddParagraphs(inputLines []string) (resultLines []string) {
	isInParagraph := false

	for _, line := range inputLines {

		if isInParagraph {
			if isHeaderLine(line) {
				resultLines = append(resultLines, "</p>")
				isInParagraph = false
			} else if isEmptyLine(line) {
				resultLines = append(resultLines, "</p>")
				isInParagraph = false
			}
		} else {
			if !isEmptyLine(line) && !isHeaderLine(line) {
				resultLines = append(resultLines, "<p>")
				isInParagraph = true
			}
		}

		resultLines = append(resultLines, line)
	}

	if isInParagraph {
		resultLines = append(resultLines, "</p>\n")
	}

	return resultLines
}

func isHeaderLine(line string) bool {
	matchHeader, _ := regexp.MatchString(`^#.*`, line)
	return matchHeader
}

func isEmptyLine(line string) bool {
	return line == ""
}
