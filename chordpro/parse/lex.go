package parse

import (
	"bufio"
	"io"
	"unicode"

	"github.com/chris-skud/chordpro2/chordpro/types"
)

// var openDirectiveRegex = regexp.MustCompile(`^\{~?&?`)

type Lexer struct {
	pos    types.Position
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    types.Position{Line: 1, Column: 0},
		reader: bufio.NewReader(reader),
	}
}

// Lex scans the input for the next token. It returns the position of the token,
// the token's type, and the literal value.
func (l *Lexer) Lex() (types.Position, types.TokenType, string) {
	// keep looping until we return a token
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, types.EOF, ""
			}
			panic(err)
		}

		// update the column to the position of the newly read in rune
		l.pos.Column++

		switch r {
		case '\n':
			l.resetPosition()
			return l.pos, types.LineBreak, "\n"
		case '[':
			startPos := l.pos
			chord := l.lexChord()
			return startPos, types.Chord, chord
		// case ']':
		// 	return l.pos, types.CloseChord, ""
		default:
			if unicode.IsSpace(r) {
				return l.pos, types.Space, " "
			} else if unicode.IsLetter(r) {
				// backup and let lexLyric rescan the beginning of the ident
				startPos := l.pos
				l.backup()
				lit := l.lexLyric()
				return startPos, types.Lyric, lit
			} else {
				return l.pos, types.Illegal, string(r)
			}
		}
	}
}

func (l *Lexer) resetPosition() {
	l.pos.Line++
	l.pos.Column = 0
}

func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos.Column--
}

// lexLyric scans the input until the end of an lyric and then returns the
// literal.
func (l *Lexer) lexLyric() string {
	var lyr string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// at the end of the lyric
				return lyr
			}
		}

		l.pos.Column++
		if unicode.IsLetter(r) {
			lyr = lyr + string(r)
		} else {
			// scanned something not in the lyric
			l.backup()
			return lyr
		}
	}
}

// lexChord scans the input until the end of the chord then returns the chord literal
func (l *Lexer) lexChord() string {
	var chrd string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// reached end of chord without a closing bracket, return what we have
				return chrd
			}
		}

		l.pos.Column++
		if r != ']' {
			chrd = chrd + string(r)
			continue
		}

		return chords[chrd] // translate unsupported notations and return
	}
}
