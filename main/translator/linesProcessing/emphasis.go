package linesProcessing

import (
	"../../utils"
	"regexp"
)

func AddEmphasis(inputLines []string) (resultLines []string) {
	regexStrong := regexp.MustCompile(`(__.+?__)|(\*\*.+?\*\*)`)
	regexItalic := regexp.MustCompile(`(_.+?_)|(\*.+?\*)`)
	regexStroke := regexp.MustCompile(`~~.+?~~`)

	for _, line := range inputLines {
		convertedStrongs := regexStrong.ReplaceAllStringFunc(line, func(text string) string {
			return `<strong>` + utils.CutStringFromBoth(text, 2, 2) + `</strong>`
		})

		convertedItalics := regexItalic.ReplaceAllStringFunc(convertedStrongs, func(text string) string {
			return `<em>` + utils.CutStringFromBoth(text, 1, 1) + `</em>`
		})

		convertedStrokes := regexStroke.ReplaceAllStringFunc(convertedItalics, func(text string) string {
			return `<del>` + utils.CutStringFromBoth(text, 2, 2) + `</del>`
		})

		resultLines = append(resultLines, convertedStrokes)
	}
	return resultLines
}
