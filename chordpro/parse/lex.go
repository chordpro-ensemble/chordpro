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
		pos:    types.Position{Line: 0, Column: 0},
		reader: bufio.NewReader(reader),
	}
}

// Lex scans the input for the next token. It returns the position of the token,
// the token's type, and the literal value.
func (l *Lexer) Lex() (types.Position, types.TokenType, string) {
	// keep looping until we return a token
	// if it's a chord being procssed, we need to correlate it to
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
		case '{':
			startPos := l.pos
			directiveType, directive := l.lexDirective()
			return startPos, directiveType, directive
		default:
			if unicode.IsSpace(r) {
				return l.pos, types.Space, " "
			} else if l.isLyric(r) {
				// backup and let lexLyric rescan the beginning of the lyric
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
		if l.isLyric(r) {
			lyr = lyr + string(r)
		} else {
			// scanned something not in the lyric
			l.backup()
			return lyr
		}
	}
}

func (l *Lexer) isLyric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsPunct(r) && string(r) != "["
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

// will need to look at spec and detemine how many top-level directives
// there are vs. inline and figure out a way to pass the accodingly to the
// output formatter. something like {soc} may be really tricky to deal with
// in the outut formatter althoug i guess you just apply a style until
// the closing bracket is found

// a directive may just have a special rule in that it's open and close are whole lines
func (l *Lexer) lexDirective() (types.TokenType, string) {
	var drctv string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// reached end of directive without a closing bracket, return what we have
				return types.DirectiveTokenType(drctv), drctv
			}
		}

		l.pos.Column++
		// could look for a linebreak as directives should always be on their own line
		if r != '}' {
			drctv = drctv + string(r)
			continue
		}

		return types.DirectiveTokenType(drctv), drctv
	}
}
