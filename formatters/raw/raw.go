package raw

import (
	"fmt"
	"io"

	"github.com/chris-skud/chordpro2/chordpro/types"
)

type Processor struct{}

// A
func (p *Processor) Process(metaDirectives types.MetaDirectives, sheetLines []types.SheetLine, w io.Writer) error {
	w.Write([]byte(fmt.Sprintf("%v", metaDirectives)))
	for _, line := range sheetLines {
		switch line.Type {
		case types.LineEnvironmentDirective:
			fmt.Fprintf(w, "Environment Directive: %s, %s", line.EnvironmentDirective.Name, line.EnvironmentDirective.Label)
		case types.LineLyric:
			fmt.Fprintf(w, "Lyric Line: %+v", line.LyricTokens)
		default:
			fmt.Fprint(w, "unknown")
		}
	}

	return nil
}
