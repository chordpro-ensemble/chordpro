package pdf

import "github.com/jung-kurt/gofpdf"

type Formatter struct{}

func (f *Formatter) Format(tokenStream []byte) []byte {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, string(tokenStream))
	_ = pdf.OutputFileAndClose("simple.pdf")
	return []byte("")
}
