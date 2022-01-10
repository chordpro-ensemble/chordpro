package parse

import (
	"bufio"
	"strings"
	"testing"

	"github.com/chordpro-ensemble/pkg/types"
	"github.com/stretchr/testify/require"
)

func TestLexer_Lex(t *testing.T) {
	type fields struct {
		pos    types.Position
		reader *bufio.Reader
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "test", fields: fields{reader: bufio.NewReader(strings.NewReader("[D]Lyric a [C+]song\n[C(maj7)]lyric")), pos: types.Position{Line: 0, Column: 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lexer{
				pos:    tt.fields.pos,
				reader: tt.fields.reader,
			}
			var tokens []types.Token
			for {
				pos, typ, lit := l.Lex()
				if typ == types.EOF {
					break
				}
				tokens = append(tokens, types.Token{Pos: pos, Typ: typ, Literal: lit})
			}
			require.Len(t, tokens, 10)
		})
	}
}
