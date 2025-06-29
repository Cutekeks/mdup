package internal

import "strings"

type ElementType int

const (
	Header1 ElementType = iota
	Header2
	Header3
	Header4
	Header5
	Header6
	Paragraph
	List
	ListItem
	Link
	Image
)

type ParsedContent struct {
	Elements []Element
}
type Element struct {
	Type    ElementType
	Content string
}

func ParseMD(content string) ParsedContent {
	var elements []Element
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "###### ") {
			elements = append(elements, Element{Type: Header6, Content: strings.TrimPrefix(line, "###### ")})
		} else if strings.HasPrefix(line, "##### ") {
			elements = append(elements, Element{Type: Header5, Content: strings.TrimPrefix(line, "##### ")})
		} else if strings.HasPrefix(line, "#### ") {
			elements = append(elements, Element{Type: Header4, Content: strings.TrimPrefix(line, "#### ")})
		} else if strings.HasPrefix(line, "### ") {
			elements = append(elements, Element{Type: Header3, Content: strings.TrimPrefix(line, "### ")})
		} else if strings.HasPrefix(line, "## ") {
			elements = append(elements, Element{Type: Header2, Content: strings.TrimPrefix(line, "## ")})
		} else if strings.HasPrefix(line, "# ") {
			elements = append(elements, Element{Type: Header1, Content: strings.TrimPrefix(line, "# ")})
		} else if line != "" {
			elements = append(elements, Element{Type: Paragraph, Content: line})
		}
	}
	return ParsedContent{Elements: elements}
}
