package chordpro

import (
	"fmt"
	"io"

	"github.com/chris-skud/chordpro2/chordpro/parse"
	"github.com/chris-skud/chordpro2/chordpro/types"
)

type OutputProcessor interface {
	Process(tokenTypes []types.Token, w io.Writer) error
}

func NewProcessor(outputProcessor OutputProcessor) *Processor {
	return &Processor{
		outputProcessor: outputProcessor,
	}

}

type Processor struct {
	outputProcessor OutputProcessor
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (p *Processor) Process(reader io.Reader, writer io.Writer) error {

	l := parse.NewLexer(reader)
	var tokens []types.Token
	for {
		pos, tok, lit := l.Lex()
		if tok == types.EOF {
			break
		}

		tokens = append(tokens, types.Token{Pos: types.Position{Line: pos.Line, Column: pos.Column}, Typ: tok, Literal: lit})
	}

	check(
		p.outputProcessor.Process(tokens, writer),
	)

	// lineCount := tokens[len(tokens)-1].Pos.Line
	// blocks := make([]types.ContentBlock, lineCount+1)
	// for _, token := range tokens {
	// 	blockPos := token.Pos.Line - 1 // adjust for 0 based slice
	// 	switch token.Typ {
	// 	case types.Lyric:
	// 		blocks[blockPos].Content += token.Literal
	// 	case types.Chord:
	// 		blocks[blockPos].Content += token.Literal
	// 	}

	// }

	fmt.Printf("\n%+v\n", tokens)
	// contentBlocks := []types.ContentBlock{}

	return nil
}
