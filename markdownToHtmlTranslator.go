package main

func translate(content string) string {
	content = wrapInHtml(content)
	return content
}

func wrapInHtml(content string) string {
	title := "Documentation"
	return `<html lang="en">
<head>
<meta charset="utf-8">
<title>` + title + `</title>
</head>
<body>` + content + `</body>
</html>`
}
