package types

// should be able to just use the position relative to it's containing lyric token
type Chord2 struct {
	RelativePos Offset
	Literal     string
}

type Token struct {
	Pos     Position
	Typ     TokenType
	Literal string
}

type Token2 struct {
	Pos     Position
	Typ     TokenType
	Literal string
	Chords  []Chord2
}
