package linesProcessing

import (
	"regexp"
)

type ReferencesInfo struct {
	link, title string
}

func ExtractReferences(inputLines []string) (resultLines []string, referencesMap map[string]ReferencesInfo) {

	regexReferences := regexp.MustCompile(`\[([^\[\]()]+)]:\s*([^\s]+)\s*(".*")?`)
	// 1 - reference
	// 2 - link
	// 3 - title (optional)(with " ")

	referencesMap = make(map[string]ReferencesInfo)

	for _, line := range inputLines {
		if regexReferences.MatchString(line) {
			params := regexReferences.FindStringSubmatch(line)
			referencesMap[params[1]] = ReferencesInfo{link: params[2], title: params[3]}
		} else {
			resultLines = append(resultLines, line)
		}
	}

	return resultLines, referencesMap
}
