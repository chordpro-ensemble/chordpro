package types

import "strings"

type MetaDirectives struct {
	Title string
}

func (m *MetaDirectives) Set(literal string) {
	literal = strings.TrimLeft(literal, "{")
	parts := strings.Split(literal, ":")
	if len(parts) != 2 {
		panic("bad directive")
	}

	directiveName := parts[0]
	directiveValue := strings.TrimSpace(parts[1])
	switch directiveName {
	case "title", "t":
		m.Title = directiveValue
	}
}
