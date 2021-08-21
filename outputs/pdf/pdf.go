package pdf

import (
	"io"

	"github.com/chris-skud/chordpro2/chordpro/types"
	"github.com/jung-kurt/gofpdf"
)

type Processor struct{}

// Processor is the PDF specific processing of the token stream (lexed song)
// whose bytes are written to the passed writer. It does need knowledge
// of the lexer grammar
func (p *Processor) Process(tokens []types.Token, w io.Writer) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 14)

	title := "Beautiful Song"

	// file metadata
	pdf.SetAuthor("Chris Skudlarczyk", false)
	pdf.SetTitle(title, true)

	// header
	pdf.SetFont("Arial", "B", 15)
	wd := pdf.GetStringWidth(title) + 6
	pdf.SetX((210 - wd) / 2)
	pdf.Cell(wd, 9, title)
	pdf.Ln(10)

	// body
	lineCount := tokens[len(tokens)-1].Pos.Line
	rows := make([]types.ContentBlock, lineCount+1)
	for _, token := range tokens {
		rowPos := token.Pos.Line - 1 // adjust for 0 based slice
		switch token.Typ {
		case types.Lyric:
			rows[rowPos].Content += token.Literal
		case types.Chord:
			rows[rowPos].Content += token.Literal
		}
	}

	for _, row := range rows {
		pdf.Cell(100, 10, row.Content)
		pdf.Ln(-1)
	}

	return pdf.Output(w)
}
