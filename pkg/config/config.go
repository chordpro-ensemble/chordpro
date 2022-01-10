package config

type Config struct {
	LyricsOnly bool `json:"lyrics-only"`
	Transpose  struct {
		Offset int `json:"offset"`
	} `json:"transpose"`
	Include    []string `json:"include"`
	Instrument struct {
		Type        string `json:"type"`
		Description string `json:"description"`
	} `json:"instrument"`
	Tuning []string `json:"tuning"`
	// Pdf    struct {
	// 	Info struct {
	// 		Title    string `json:"title"`
	// 		Author   string `json:"author"`
	// 		Subject  string `json:"subject"`
	// 		Keywords string `json:"keywords"`
	// 	} `json:"info"`
	// 	Papersize string `json:"papersize"`
	// 	Theme     struct {
	// 		Foreground string `json:"foreground"`
	// 		Background string `json:"background"`
	// 	} `json:"theme"`
	// 	Columnspace   int  `json:"columnspace"`
	// 	Margintop     int  `json:"margintop"`
	// 	Marginbottom  int  `json:"marginbottom"`
	// 	Marginleft    int  `json:"marginleft"`
	// 	Marginright   int  `json:"marginright"`
	// 	Headspace     int  `json:"headspace"`
	// 	Footspace     int  `json:"footspace"`
	// 	HeadFirstOnly bool `json:"head-first-only"`
	// 	Spacing       struct {
	// 		Title  float64 `json:"title"`
	// 		Lyrics float64 `json:"lyrics"`
	// 		Chords float64 `json:"chords"`
	// 		Grid   float64 `json:"grid"`
	// 		Tab    int     `json:"tab"`
	// 		Toc    float64 `json:"toc"`
	// 		Empty  int     `json:"empty"`
	// 	} `json:"spacing"`
	// 	Chorus struct {
	// 		Indent int `json:"indent"`
	// 		Bar    struct {
	// 			Offset int    `json:"offset"`
	// 			Width  int    `json:"width"`
	// 			Color  string `json:"color"`
	// 		} `json:"bar"`
	// 		Tag    string `json:"tag"`
	// 		Recall struct {
	// 			Choruslike bool   `json:"choruslike"`
	// 			Tag        string `json:"tag"`
	// 			Type       string `json:"type"`
	// 			Quote      bool   `json:"quote"`
	// 		} `json:"recall"`
	// 	} `json:"chorus"`
	// 	Labels struct {
	// 		Width   string `json:"width"`
	// 		Align   string `json:"align"`
	// 		Comment string `json:"comment"`
	// 	} `json:"labels"`
	// 	Chordscolumn          int    `json:"chordscolumn"`
	// 	Capoheading           string `json:"capoheading"`
	// 	TitlesDirectiveIgnore bool   `json:"titles-directive-ignore"`
	// 	Diagrams              struct {
	// 		Show      string  `json:"show"`
	// 		Width     int     `json:"width"`
	// 		Height    int     `json:"height"`
	// 		Vcells    int     `json:"vcells"`
	// 		Linewidth float64 `json:"linewidth"`
	// 		Hspace    float64 `json:"hspace"`
	// 		Vspace    int     `json:"vspace"`
	// 	} `json:"diagrams"`
	// 	Kbdiagrams struct {
	// 		Show      string  `json:"show"`
	// 		Width     int     `json:"width"`
	// 		Height    int     `json:"height"`
	// 		Keys      int     `json:"keys"`
	// 		Base      string  `json:"base"`
	// 		Linewidth float64 `json:"linewidth"`
	// 		Pressed   string  `json:"pressed"`
	// 		Hspace    float64 `json:"hspace"`
	// 		Vspace    float64 `json:"vspace"`
	// 	} `json:"kbdiagrams"`
	// 	EvenOddPages   int `json:"even-odd-pages"`
	// 	PagealignSongs int `json:"pagealign-songs"`
	// 	Formats        struct {
	// 		Default struct {
	// 			Title    []string `json:"title"`
	// 			Subtitle []string `json:"subtitle"`
	// 			Footer   []string `json:"footer"`
	// 		} `json:"default"`
	// 		Title struct {
	// 			Title    []string `json:"title"`
	// 			Subtitle []string `json:"subtitle"`
	// 			Footer   []string `json:"footer"`
	// 		} `json:"title"`
	// 		First struct {
	// 			Footer []string `json:"footer"`
	// 		} `json:"first"`
	// 	} `json:"formats"`
	// 	SplitMarker []string      `json:"split-marker"`
	// 	Fontdir     []interface{} `json:"fontdir"`
	// 	Fontconfig  struct {
	// 		Times struct {
	// 			string     `json:""`
	// 			Bold       string `json:"bold"`
	// 			Italic     string `json:"italic"`
	// 			Bolditalic string `json:"bolditalic"`
	// 		} `json:"times"`
	// 		Helvetica struct {
	// 			string      `json:""`
	// 			Bold        string `json:"bold"`
	// 			Oblique     string `json:"oblique"`
	// 			Boldoblique string `json:"boldoblique"`
	// 		} `json:"helvetica"`
	// 		Courier struct {
	// 			string     `json:""`
	// 			Bold       string `json:"bold"`
	// 			Italic     string `json:"italic"`
	// 			Bolditalic string `json:"bolditalic"`
	// 		} `json:"courier"`
	// 		Dingbats struct {
	// 			string `json:""`
	// 		} `json:"dingbats"`
	// 	} `json:"fontconfig"`
	// 	Fonts struct {
	// 		Title struct {
	// 			Name string `json:"name"`
	// 			Size int    `json:"size"`
	// 		} `json:"title"`
	// 		Text struct {
	// 			Name string `json:"name"`
	// 			Size int    `json:"size"`
	// 		} `json:"text"`
	// 		Chord struct {
	// 			Name string `json:"name"`
	// 			Size int    `json:"size"`
	// 		} `json:"chord"`
	// 		Chordfingers struct {
	// 			Name        string `json:"name"`
	// 			Size        int    `json:"size"`
	// 			Numbercolor string `json:"numbercolor"`
	// 		} `json:"chordfingers"`
	// 		Comment struct {
	// 			Name       string `json:"name"`
	// 			Size       int    `json:"size"`
	// 			Background string `json:"background"`
	// 		} `json:"comment"`
	// 		CommentItalic struct {
	// 			Name string `json:"name"`
	// 			Size int    `json:"size"`
	// 		} `json:"comment_italic"`
	// 		CommentBox struct {
	// 			Name  string `json:"name"`
	// 			Size  int    `json:"size"`
	// 			Frame int    `json:"frame"`
	// 		} `json:"comment_box"`
	// 		Tab struct {
	// 			Name string `json:"name"`
	// 			Size int    `json:"size"`
	// 		} `json:"tab"`
	// 		Toc struct {
	// 			Name string `json:"name"`
	// 			Size int    `json:"size"`
	// 		} `json:"toc"`
	// 		Grid struct {
	// 			Name string `json:"name"`
	// 			Size int    `json:"size"`
	// 		} `json:"grid"`
	// 	} `json:"fonts"`
	// 	Outlines []struct {
	// 		Fields   []string `json:"fields"`
	// 		Label    string   `json:"label"`
	// 		Line     string   `json:"line"`
	// 		Collapse bool     `json:"collapse"`
	// 		Letter   int      `json:"letter"`
	// 		Fold     bool     `json:"fold"`
	// 	} `json:"outlines"`
	// 	Showlayout bool `json:"showlayout"`
	// 	Csv        struct {
	// 		Songsonly bool `json:"songsonly"`
	// 	} `json:"csv"`
	// } `json:"pdf"`
	// Chordpro struct {
	// 	Chorus struct {
	// 		Recall struct {
	// 			Tag   string `json:"tag"`
	// 			Type  string `json:"type"`
	// 			Quote bool   `json:"quote"`
	// 		} `json:"recall"`
	// 	} `json:"chorus"`
	// } `json:"chordpro"`
	// HTML struct {
	// 	Styles struct {
	// 		Display string `json:"display"`
	// 		Print   string `json:"print"`
	// 	} `json:"styles"`
	// } `json:"html"`
}
