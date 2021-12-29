package chordpro

import (
	"io"

	"github.com/chris-skud/chordpro2/chordpro/parse"
	"github.com/chris-skud/chordpro2/chordpro/types"
)

type OutputProcessor interface {
	Process(metaDirectives types.MetaDirectives, sheetLines []types.SheetLine, w io.Writer) error
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
	metaDirectives := types.MetaDirectives{}

	l := parse.NewLexer(reader)
	var tokens []types.Token2
	// var offset int
	// var currLine int
	var activeChord *types.Chord2
	for {
		pos, tok, lit := l.Lex()
		if tok == types.EOF {
			break
		}

		// set an active chord to later associate it to a lyric
		if tok == types.Chord {
			activeChord = &types.Chord2{
				RelativePos: 0,
				Literal:     lit,
			}
			// chords are added to lyric tokens and not
			// tokens themselves
			continue
		}

		if tok == types.MetaDirective {
			metaDirectives.Set(lit)
			// meta directives are not added as tokens
			continue
		}

		// it's a lyric so add it to the slice of tokens
		token2 := types.Token2{
			Pos: types.Position{
				Line:   pos.Line,
				Column: pos.Column,
			},
			Typ: tok, Literal: lit,
		}
		if activeChord != nil {
			token2.Chords = append(
				token2.Chords,
				*activeChord,
			)
			activeChord = nil
		}

		tokens = append(tokens, token2)
	}

	// convert tokens slice into typed rows of token slices,
	// including the creation of lyric-chord sets (a given lyric)
	// usually (but not always) has a row above it with chords

	// Seems possible to apply a pipeline-like pattern to this processing, but it's not yeat clear
	// whether the linear input-as-last-output really fits as any processing may need the original token list vs. a previously
	// mutated sheet list. Stuff like formatting directives ({soc}...{eoc)) are a wrinkle.

	lineCount := tokens[len(tokens)-1].Pos.Line + 1
	sheetLines := make([]types.SheetLine, lineCount)
	for _, token := range tokens {
		linePos := token.Pos.Line
		switch token.Typ {
		// case types.Chord:
		// 	sheetLines[linePos].Type = types.LyricChord
		// 	sheetLines[linePos].LyricChordSet.Chords = append(
		// 		sheetLines[linePos].LyricChordSet.Chords,
		// 		token,
		// 	)
		case types.Lyric, types.Space:
			sheetLines[linePos].Type = types.LyricChord
			sheetLines[linePos].LyricChordSet.Lyrics = append(
				sheetLines[linePos].LyricChordSet.Lyrics,
				token,
			)
		}
	}

	check(
		p.outputProcessor.Process(metaDirectives, sheetLines, writer),
	)

	return nil
}
