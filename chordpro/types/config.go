package types

type Config struct {
	LyricsOnly bool `json:"lyricsOnly"`
	Transpose  struct {
		Offset int `json:"offset"`
	} `json:"transpose"`
	Instrument struct {
		Type        string `json:"type"`
		Description string `json:"description"`
	} `json:"instrument"`
	Tuning []string `json:"tuning"`
}
