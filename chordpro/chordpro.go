package chordpro

type Formatter interface {
	Format(tokenStream []byte) []byte
}

type Processor struct {
	Formatter Formatter
}

func (p *Processor) Process(rawSong []byte) []byte {
	// lex it
	return p.Formatter.Format(rawSong)
}
