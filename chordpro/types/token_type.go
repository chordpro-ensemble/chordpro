package types

type TokenType int

const (
	EOF TokenType = iota
	Illegal
	Lyric // anything that is not covered by other reserved tokens
	Chord // string inside square brackets
	MetaDirective
	LineBreak // "\n"
	Space
)

var TokenTypes = []string{
	EOF:           "EOF",
	Illegal:       "ILLEGAL",
	Lyric:         "LYRIC",
	Chord:         "CHORD",
	MetaDirective: "METADIRECTIVE",
	LineBreak:     "LINEBREAK",
	Space:         "SPACE",
}

func (t TokenType) String() string {
	return TokenTypes[t]
}
