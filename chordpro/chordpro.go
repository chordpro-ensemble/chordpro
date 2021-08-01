package chordpro

type Formatter interface {
	Format() []byte
}

type Processor struct {
	Formatter Formatter
}

func (p *Processor) Parse() {
}
