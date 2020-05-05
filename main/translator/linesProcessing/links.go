package linesProcessing

import (
	"regexp"
)

//TODO: first images before links
func AddLinks(inputLines []string, referencesMap map[string]ReferencesInfo) (resultLines []string) {
	regexLinkClassic := regexp.MustCompile(`\[([^\[\]()]+)]\(([\w\d.\/:?=#]+)\)`)
	regexLinkWithTitle := regexp.MustCompile(`\[([^\[\]()]+)]\(([\w\d.\/\-:?=#]+)\s"([\w\d\s.\/\-:?'=#]+)"\)`)
	regexLinkReference := regexp.MustCompile(`\[([^\[\]()]+)]\[([\w\s\d.\/\-\\:?=#]+)\]`)

	for _, line := range inputLines {
		converted := regexLinkClassic.ReplaceAllString(line, "<a href=\"$2\">$1</a>")

		converted = regexLinkWithTitle.ReplaceAllString(converted, "<a href=\"$2\" title=\"$3\">$1</a>")

		if regexLinkReference.MatchString(converted) {
			params := regexLinkReference.FindStringSubmatch(converted)
			linkText := params[1]
			linkReference := params[2]
			referencesInfo := referencesMap[linkReference]
			converted = regexLinkReference.ReplaceAllString(converted, `<a href="`+referencesInfo.link+`" title="`+referencesInfo.title+`">`+linkText+`</a>`)
		}

		resultLines = append(resultLines, converted)
	}
	return resultLines
}
