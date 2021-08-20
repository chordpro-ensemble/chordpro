package types

type ContentBlockType int

const (
	ChordBlock ContentBlockType = iota
	LyricBlock
)

type ContentBlock struct {
	Content string
	Type    ContentBlockType
}
