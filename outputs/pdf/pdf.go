package pdf

import (
	"github.com/jung-kurt/gofpdf"
)

type Formatter struct{}

func (f *Formatter) Format(tokenStream []byte) []byte {
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
	pdf.Cell(100, 10, "furst")
	pdf.Ln(-1)
	pdf.Cell(100, 10, "second asdf")

	// write the file
	_ = pdf.OutputFileAndClose("simple.pdf")
	return []byte("")
}
