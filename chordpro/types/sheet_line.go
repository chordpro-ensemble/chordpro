package types

type SheetLineType int

const (
	LyricChord SheetLineType = iota
	LineEnvironmentDirective
)

type SheetLine struct {
	Type          SheetLineType
	LyricChordSet LyricChordSet
}

type LyricChordSet struct {
	Chords []Token
	Lyrics []Token
}

/*
[
    {
        "rowset": {
            "chords": [{}],
            "lyrics": [{}],
        }
    }
]
*/
