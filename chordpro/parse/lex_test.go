package parse

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/chris-skud/chordpro2/chordpro/types"
)

func TestLexer_Lex(t *testing.T) {
	type fields struct {
		pos    types.Position
		reader *bufio.Reader
	}
	tests := []struct {
		name   string
		fields fields
		want   types.Position
		want1  types.TokenType
		want2  string
	}{
		{name: "test", fields: fields{reader: bufio.NewReader(strings.NewReader("[D]This is a [C+] good song\n[C(maj7)]It is")), pos: types.Position{Line: 0, Column: 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				pos:    tt.fields.pos,
				reader: tt.fields.reader,
			}
			var out string
			for {
				pos, tok, lit := l.Lex()
				if tok == types.EOF {
					break
				}

				// out = out + str
				fmt.Printf("%d:%d %s %s\n", pos.Line, pos.Column, tok, lit)
			}

			fmt.Println(out)
			// got, got1, got2 := l.Lex()
			// require.Equal(t, Position{line: 0, column: 1}, got) // parse.Position{line: 0, column: 1}
			// require.Equal(t, Token(3), got1)
			// require.Equal(t, "[", got2)
		})
	}
}
