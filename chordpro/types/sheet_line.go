package types

type SheetLineType int

const (
	LineLyric SheetLineType = iota
	LineEnvironmentDirective
)

type SheetLine struct {
	Type                 SheetLineType
	Lyrics               []Token
	EnvironmentDirective DirectiveData
}
