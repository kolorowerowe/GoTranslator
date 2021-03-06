package linesProcessing

import (
	"../../utils"
	"regexp"
)

func AddHeaders(inputLines []string) (resultLines []string) {

	for _, line := range inputLines {
		match6, _ := regexp.MatchString(`^######.*`, line)
		match5, _ := regexp.MatchString(`^#####.*`, line)
		match4, _ := regexp.MatchString(`^####.*`, line)
		match3, _ := regexp.MatchString(`^###.*`, line)
		match2, _ := regexp.MatchString(`^##.*`, line)
		match1, _ := regexp.MatchString(`^#.*`, line)

		if match6 {
			resultLines = append(resultLines, "<h6>"+utils.CutStringFromBeginning(line, 6)+"</h6>")
		} else if match5 {
			resultLines = append(resultLines, "<h5>"+utils.CutStringFromBeginning(line, 5)+"</h5>")
		} else if match4 {
			resultLines = append(resultLines, "<h4>"+utils.CutStringFromBeginning(line, 4)+"</h4>")
		} else if match3 {
			resultLines = append(resultLines, "<h3>"+utils.CutStringFromBeginning(line, 3)+"</h3>")
		} else if match2 {
			resultLines = append(resultLines, "<h2>"+utils.CutStringFromBeginning(line, 2)+"</h2>")
		} else if match1 {
			resultLines = append(resultLines, "<h1>"+utils.CutStringFromBeginning(line, 1)+"</h1>")
		} else {
			resultLines = append(resultLines, line)
		}
	}
	return resultLines
}
