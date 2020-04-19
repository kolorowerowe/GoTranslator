package linesProcessing

import "regexp"

func AddEmphasis(inputLines []string) (resultLines []string) {
	regexStrong := regexp.MustCompile(`(__.+?__)|(\*\*.+?\*\*)`)
	regexItalic := regexp.MustCompile(`(_.+?_)|(\*.+?\*)`)
	regexStroke := regexp.MustCompile(`~~.+?~~`)

	for _, line := range inputLines {
		convertedStrongs := regexStrong.ReplaceAllStringFunc(line, func(n string) string {
			return `<strong>` + n[2:len(n)-2] + `</strong>`
		})

		convertedItalics := regexItalic.ReplaceAllStringFunc(convertedStrongs, func(n string) string {
			return `<em>` + n[1:len(n)-1] + `</em>`
		})

		convertedStrokes := regexStroke.ReplaceAllStringFunc(convertedItalics, func(n string) string {
			return `<del>` + n[2:len(n)-2] + `</del>`
		})

		resultLines = append(resultLines, convertedStrokes)
	}
	return resultLines
}
