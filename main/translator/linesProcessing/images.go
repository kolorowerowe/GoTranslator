package linesProcessing

import (
	"regexp"
)

//TODO: first images before links
func AddImages(inputLines []string, referencesMap map[string]ReferencesInfo) (resultLines []string) {
	regexImageSimple := regexp.MustCompile(`!\[([^\[\]()]+)]\(([\w\d.\/\-:?=#]+)\)`)
	regexImageWithTitle := regexp.MustCompile(`!\[([^\[\]()]+)]\(([\w\d.\/\-:?=#]+)\s"([\w\d\s.\/\-:?'=#]+)"\)`)
	regexImageReference := regexp.MustCompile(`!\[([^\[\]()]+)]\[([\w\s\d.\/\-\\:?=#]+)\]`)

	for _, line := range inputLines {
		converted := regexImageSimple.ReplaceAllString(line, "<img alt=\"$1\" src=\"$2\" />")

		converted = regexImageWithTitle.ReplaceAllString(converted, "<img alt=\"$1\" src=\"$2\" title=\"$3\" />")

		if regexImageReference.MatchString(converted) {
			params := regexImageReference.FindStringSubmatch(converted)
			linkText := params[1]
			linkReference := params[2]
			referencesInfo := referencesMap[linkReference]
			converted = regexImageReference.ReplaceAllString(converted, `<img alt="`+linkText+`" src="`+referencesInfo.link+`" title=`+referencesInfo.title+` />`)
		}

		resultLines = append(resultLines, converted)
	}
	return resultLines
}
