package linesProcessing

import (
	"../../utils"
	"regexp"
)

type InListType string

const (
	NONE      InListType = "None"
	ORDERED              = "Ordered"
	UNORDERED            = "UNORDERED"
)

func AddLists(inputLines []string) (resultLines []string) {

	inList := NONE

	for _, line := range inputLines {
		lineToAppend := line
		switch inList {
		case NONE:
			if isOrderedListElement(line) {
				resultLines = append(resultLines, "<ol>")
				lineToAppend = convertToListElement(lineToAppend)
				inList = ORDERED
			} else if isUnorderedListElement(line) {
				resultLines = append(resultLines, "<ul>")
				lineToAppend = convertToListElement(lineToAppend)
				inList = UNORDERED
			}
		case ORDERED:
			if isOrderedListElement(line) {
				lineToAppend = convertToListElement(lineToAppend)
			} else if isUnorderedListElement(lineToAppend) {
				resultLines = append(resultLines, "</ol>")
				resultLines = append(resultLines, "<ul>")
				lineToAppend = convertToListElement(lineToAppend)
				inList = UNORDERED
			} else {
				resultLines = append(resultLines, "</ol>")
				inList = NONE
			}
		case UNORDERED:
			if isOrderedListElement(line) {
				resultLines = append(resultLines, "</ul>")
				resultLines = append(resultLines, "<ol>")
				lineToAppend = convertToListElement(lineToAppend)
				inList = ORDERED
			} else if isUnorderedListElement(lineToAppend) {
				lineToAppend = convertToListElement(lineToAppend)
			} else {
				resultLines = append(resultLines, "</ul>")
				inList = NONE
			}
		}

		resultLines = append(resultLines, lineToAppend)
	}

	return resultLines
}

func isUnorderedListElement(line string) bool {
	matchHeader, _ := regexp.MatchString(`^[\*\-\+] .*`, line)
	return matchHeader
}

func convertToListElement(line string) string {
	return "<li>" + utils.CutStringFromBeginning(line, 2) + "</li>"
}

func isOrderedListElement(line string) bool {
	matchHeader, _ := regexp.MatchString(`^\d\..*`, line)
	return matchHeader
}
