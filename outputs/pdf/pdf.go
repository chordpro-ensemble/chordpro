package pdf

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/chris-skud/chordpro2/chordpro/types"
	"github.com/phpdave11/gofpdf"
)

type Processor struct{}

// Processor is the PDF specific processing of the token stream (lexed song)
// whose bytes are written to the passed writer. It does need knowledge
// of the lexer grammar
func (p *Processor) Process(sheetLines []types.SheetLine, w io.Writer) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	title := "Beautiful Song"

	// file metadata
	pdf.SetAuthor("Chris Skudlarczyk", false)
	pdf.SetTitle(title, true)

	// header
	pdf.SetFont("Courier", "", 12)
	wd := pdf.GetStringWidth(title) + 6
	pdf.SetX((210 - wd) / 2)
	pdf.Cell(wd, 9, title)
	pdf.Ln(10)

	b, _ := json.MarshalIndent(sheetLines, "", "  ")
	fmt.Println(string(b))

	for _, line := range sheetLines {
		switch line.Type {
		case types.LyricChord:

			// Process the chord line if there are chords
			// maybe I should just write an empty line here
			// to be filled in the next line.
			if len(line.LyricChordSet.Chords) != 0 {
				pdf.SetX(5)
				// var chords string
				// for _, chordToken := range line.LyricChordSet.Chords {
				// 	pad := spaces(chordToken.Pos.Column - len(chords))
				// 	chords += pad + chordToken.Literal
				// }
				// pdf.CellFormat(200, 6, chords, "0", 0, "", false, 0, "")
				pdf.Ln(-1)
			}

			// Process the lyrics line
			pdf.SetX(5)
			// var lyrics string
			for _, lyricToken := range line.LyricChordSet.Lyrics {

				// ok, this is a hack around what may be a larger
				// issue with the pdf package not properly supporting utf-8 chars
				lyricToken.Literal = strings.ReplaceAll(lyricToken.Literal, "’", "'")

				// add current literal to lyrics string
				// fmt.Println(lyricToken.Literal)
				// lyrics += lyricToken.Literal

				thisCol := lyricToken.Pos.Column
				var chord string
				for _, chordToken := range line.LyricChordSet.Chords {
					fmt.Println(chordToken.Literal)

					// need to offset the column to account for length of chord literal

					if chordToken.Pos.Column == thisCol {
						chord = chordToken.Literal
						break
					}
				}

				fmt.Println(chord)

				preX, preY := pdf.GetXY()
				// p, _ := pdf.GetFontSize()
				// w := float64(len(lyricToken.Literal)) * p
				w := pdf.GetStringWidth(lyricToken.Literal)

				pdf.CellFormat(w, 6, lyricToken.Literal, "0", 0, "L", false, 0, "")
				if chord != "" {
					postLyricX, postLyricY := pdf.GetXY()
					pdf.SetXY(preX, preY-10)
					cw := pdf.GetStringWidth(chord)
					pdf.CellFormat(cw, 6, chord, "0", 0, "", false, 0, "")

					pdf.SetXY(postLyricX, postLyricY)
				}

			}
			// pdf.CellFormat(200, 6, lyrics, "0", 0, "", false, 0, "")
			pdf.Ln(10)

		//case types.Directive:
		default:
		}
	}

	// how the perl project does it
	// https://github.com/ChordPro/chordpro/blob/24ee42236c093063c03cda9f7545e9b9fd117c1a/lib/App/Music/ChordPro/Output/PDF.pm#L1292

	return pdf.Output(w)
}

func spaces(count int) string {
	return strings.Repeat(" ", count)
}
