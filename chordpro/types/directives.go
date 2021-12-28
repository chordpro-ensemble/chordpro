package types

type DirectiveType int

const (
	StartOfChorus DirectiveType = iota
	EndOfChorus
)

// top level directives like title
type Directive struct {
	Type DirectiveType
}

// inline directives like start-of-chorus
type InlineDirective struct {
	Type DirectiveType
}
