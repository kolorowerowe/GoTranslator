package linesProcessing

import "regexp"

func AddHorizontalRules(inputLines []string) (resultLines []string) {
	re := regexp.MustCompile("^(_{3,})|(\\*{3,})|(-{3,}).*")

	for _, line := range inputLines {
		if re.MatchString(line) {
			resultLines = append(resultLines, "<hr>")
		} else {
			resultLines = append(resultLines, line)
		}
	}
	return resultLines
}
