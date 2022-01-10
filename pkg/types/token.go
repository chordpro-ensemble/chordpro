package types

// should be able to just use the position relative to it's containing lyric token
type ChordToken struct {
	RelativePos Offset
	Literal     string
}

type Token struct {
	Pos     Position
	Typ     TokenType
	Literal string
	Chords  []ChordToken
}
