package chordpro

type Token int

const (
	EOF = iota
	IDENT
	CHORDOPEN
	CHORDCLOSE
)

var tokens = []string{
	EOF:        "EOF",
	IDENT:      "IDENT",
	CHORDOPEN:  "[",
	CHORDCLOSE: "]",
}

func (t Token) String() string {
	return tokens[t]
}

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
