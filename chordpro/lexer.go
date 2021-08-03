package chordpro

import (
	"bufio"
	"io"
	"log"
)

type Token int

const (
	EOF = iota
	IDENT
	CHORDOPEN
	CHORDCLOSE
)

var tokens = []string{
	EOF:        "EOF",
	IDENT:      "IDENT",
	CHORDOPEN:  "[",
	CHORDCLOSE: "]",
}

func (t Token) String() string {
	return tokens[t]
}

type Position struct {
	line   int
	column int
}

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{line: 1, column: 0},
		reader: bufio.NewReader(reader),
	}
}

// Lex scans the input for the next token. It returns the position of the token,
// the token's type, and the literal value.
func (l *Lexer) Lex() (Position, Token, string) {
	// keep looping until we return a token
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}

			// at this point there isn't much we can do
			log.Fatal(err)
		}

		// update the column to the position of the newly read in rune
		l.pos.column++

		switch r {
		case '\n':
			l.resetPosition()
		case '[':
			return l.pos, CHORDOPEN, "["
		case ']':
			return l.pos, CHORDCLOSE, "]"
		default:
			continue
			// if unicode.IsSpace(r) {
			// 	continue // nothing to do here, just move on
			// } else if unicode.IsDigit(r) {
			// 	// backup and let lexInt rescan the beginning of the int
			// 	startPos := l.pos
			// 	l.backup()
			// 	lit := l.lexInt()
			// 	return startPos, INT, lit
			// } else if unicode.IsLetter(r) {
			// 	// backup and let lexIdent rescan the beginning of the ident
			// 	startPos := l.pos
			// 	l.backup()
			// 	lit := l.lexIdent()
			// 	return startPos, IDENT, lit
			// } else {
			// 	return l.pos, ILLEGAL, string(r)
			// }
		}
	}
}

func (l *Lexer) resetPosition() {
	l.pos.line++
	l.pos.column = 0
}
