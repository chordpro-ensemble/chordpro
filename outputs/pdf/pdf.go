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
func (p *Processor) Process(tokenTypes []types.TokenType, w io.Writer) error {
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
	// for _, block := range contentBlocks {
	// 	pdf.Cell(100, 10, block.Content)
	// 	pdf.Ln(-1)
	// }

	// write the file
	_ = pdf.OutputFileAndClose("simple.pdf")

	return nil
}
