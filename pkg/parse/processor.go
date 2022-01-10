package parse

import (
	"io"

	"github.com/chordpro-ensemble/pkg/types"
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

	l := NewLexer(reader)
	var tokens []types.Token
	// var offset int
	// var currLine int
	var activeChord *types.ChordToken
	for {
		pos, typ, lit := l.Lex()
		if typ == types.EOF {
			break
		}

		// set an active chord to later associate it to a lyric
		if typ == types.Chord {
			activeChord = &types.ChordToken{
				RelativePos: 0,
				Literal:     lit,
			}
			// chords are added to lyric tokens and not
			// tokens themselves
			continue
		}

		if typ == types.MetaDirective {
			metaDirectives.Set(lit)
			// meta directives are not added as tokens
			continue
		}

		// it's a lyric so add it to the slice of tokens
		token := types.Token{
			Pos: types.Position{
				Line:   pos.Line,
				Column: pos.Column,
			},
			Typ: typ, Literal: lit,
		}
		if activeChord != nil {
			token.Chords = append(
				token.Chords,
				*activeChord,
			)
			activeChord = nil
		}

		tokens = append(tokens, token)
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
		// 	sheetLines[linePos].Type = types.LineLyric
		// 	sheetLines[linePos].LineLyricSet.Chords = append(
		// 		sheetLines[linePos].LineLyricSet.Chords,
		// 		token,
		// 	)
		case types.Lyric, types.Space:
			sheetLines[linePos].Type = types.LineLyric
			sheetLines[linePos].LyricTokens = append(
				sheetLines[linePos].LyricTokens,
				token,
			)
		case types.EnvironmentDirective:
			sheetLines[linePos].Type = types.LineEnvironmentDirective
			sheetLines[linePos].EnvironmentDirective = types.DirectiveData{
				Name: token.Literal,
				// Label: "",
			}
		}
	}

	check(
		p.outputProcessor.Process(metaDirectives, sheetLines, writer),
	)

	return nil
}
