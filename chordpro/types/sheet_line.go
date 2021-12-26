package types

type SheetLineType int

const (
	LyricChord SheetLineType = iota
)

type SheetLine struct {
	Type          SheetLineType
	LyricChordSet LyricChordSet
}

type LyricChordSet struct {
	Chords []Token2
	Lyrics []Token2
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
