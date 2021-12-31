package types

import (
	"fmt"
	"strings"
)

type DirectiveType int

const (
	Meta DirectiveType = iota
	Environment
)

type Directive int

const (
	StartOfChorus Directive = iota
	EndOfChorus
)

func DirectiveTypeFromString() string {
	return ""
}

type DirectiveData struct {
	Name  string
	Label string
}

type MetaDirectives map[string]DirectiveData

func (m MetaDirectives) Title() string {
	if title, ok := m["title"]; ok {
		return title.Label
	}
	return ""
}

func (m MetaDirectives) Set(literal string) {
	directiveData := directiveParts(literal)
	switch directiveData.Name {
	case "title", "t":
		m["title"] = DirectiveData{Name: "title", Label: directiveData.Label}
	}
}

func directiveParts(literal string) DirectiveData {
	literal = strings.TrimLeft(literal, "{")
	parts := strings.Split(literal, ":")
	if len(parts) == 1 {
		// it's potentially a non-label type of directive like {soc}
		return DirectiveData{Name: parts[0]}
	}
	if len(parts) != 2 {
		fmt.Println("error: unsupported meta directive: " + literal)
		return DirectiveData{}
	}

	name := strings.TrimSpace(parts[0])
	label := strings.TrimSpace(parts[1])

	return DirectiveData{Name: name, Label: label}
}

func DirectiveTokenType(literal string) TokenType {
	directiveData := directiveParts(literal)
	switch directiveData.Name {
	// this could ultimately be faster with a big regex
	case "title", "t":
		return MetaDirective
	case "start_of_chorus", "soc":
		return EnvironmentDirective
	case "end_of_chorus", "eoc":
		return EnvironmentDirective
	default:
		return Illegal
	}

}
