package internal

import (
	"strings"
)

func ConvertHTML(content ParsedContent) string {
	var sb strings.Builder
	sb.WriteString(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Your Page Title</title>
    <link rel="stylesheet" href="styles.css">
</head>
<body>
`)
	for _, element := range content.Elements {
		switch element.Type {
		case Header1:
			sb.WriteString("<h1>" + element.Content + "</h1>\n")
		case Header2:
			sb.WriteString("<h2>" + element.Content + "</h2>\n")
		case Header3:
			sb.WriteString("<h3>" + element.Content + "</h3>\n")
		case Header4:
			sb.WriteString("<h4>" + element.Content + "</h4>\n")
		case Header5:
			sb.WriteString("<h5>" + element.Content + "</h5>\n")
		case Header6:
			sb.WriteString("<h6>" + element.Content + "</h6>\n")
		case Paragraph:
			sb.WriteString("<p>" + element.Content + "</p>\n")
		case List:
			sb.WriteString("<ul>")
			// Handle list items if needed
			sb.WriteString("</ul>")
		case ListItem:
			sb.WriteString("<li>" + element.Content + "</li>")
		case Link:
			parts := strings.Split(element.Content, "|")
			if len(parts) == 2 {
				sb.WriteString("<a href=\"" + parts[1] + "\"alt=\"" + parts[0] + "\" />")
			}
		case Image:
			parts := strings.Split(element.Content, "|")
			if len(parts) == 2 {
				sb.WriteString("<img href=\"" + parts[1] + "\"alt=\"" + parts[0] + "\" />")
			}
		}

	}

	return sb.String()
}
