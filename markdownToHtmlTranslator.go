package main

func translate(content string, fileName string) string {
	content = wrapInHtml(content, fileName)
	return content
}

func wrapInHtml(content string, fileName string) string {
	return `<html lang="en"><head><meta charset="utf-8"><title>` + fileName + `</title></head><body>` + content + `</body></html>`
}
