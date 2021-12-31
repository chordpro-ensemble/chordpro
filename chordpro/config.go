package chordpro

type Config struct {
	Include  []string `json:"include"`
	Settings struct {
		Strict              bool   `json:"strict"`
		Lineinfo            bool   `json:"lineinfo"`
		Titles              string `json:"titles"`
		Columns             int    `json:"columns"`
		SuppressEmptyChords bool   `json:"suppress-empty-chords"`
		SuppressEmptyLyrics bool   `json:"suppress-empty-lyrics"`
		LyricsOnly          bool   `json:"lyrics-only"`
		Memorize            bool   `json:"memorize"`
		InlineChords        bool   `json:"inline-chords"`
		ChordsUnder         bool   `json:"chords-under"`
		Transpose           int    `json:"transpose"`
		Transcode           string `json:"transcode"`
		Decapo              bool   `json:"decapo"`
		Chordnames          string `json:"chordnames"`
		Notenames           bool   `json:"notenames"`
	} `json:"settings"`
	Metadata struct {
		Keys      []string `json:"keys"`
		Strict    bool     `json:"strict"`
		Separator string   `json:"separator"`
	} `json:"metadata"`
	Meta struct {
	} `json:"meta"`
	Dates struct {
		Today struct {
			Format string `json:"format"`
		} `json:"today"`
	} `json:"dates"`
	User struct {
		Name     string `json:"name"`
		Fullname string `json:"fullname"`
	} `json:"user"`
	Instrument struct {
		Type        string `json:"type"`
		Description string `json:"description"`
	} `json:"instrument"`
	Tuning []string `json:"tuning"`
	Notes  struct {
		System  string        `json:"system"`
		Sharp   []interface{} `json:"sharp"`
		Flat    []interface{} `json:"flat"`
		Movable bool          `json:"movable"`
	} `json:"notes"`
	Chords   []interface{} `json:"chords"`
	Diagrams struct {
		Auto   bool   `json:"auto"`
		Show   string `json:"show"`
		Sorted bool   `json:"sorted"`
	} `json:"diagrams"`
	Diagnostics struct {
		Format string `json:"format"`
	} `json:"diagnostics"`
	Contents []struct {
		Fields []string `json:"fields"`
		Label  string   `json:"label"`
		Line   string   `json:"line"`
		Pageno string   `json:"pageno"`
		Fold   bool     `json:"fold"`
		Omit   bool     `json:"omit"`
	} `json:"contents"`
	Toc struct {
		Title string `json:"title"`
		Line  string `json:"line"`
		Order string `json:"order"`
	} `json:"toc"`
	Delegates struct {
		Abc struct {
			Type    string `json:"type"`
			Handler string `json:"handler"`
		} `json:"abc"`
		Ly struct {
			Type    string `json:"type"`
			Handler string `json:"handler"`
		} `json:"ly"`
	} `json:"delegates"`
	Pdf struct {
		Info struct {
			Title    string `json:"title"`
			Author   string `json:"author"`
			Subject  string `json:"subject"`
			Keywords string `json:"keywords"`
		} `json:"info"`
		Papersize string `json:"papersize"`
		Theme     struct {
			Foreground string `json:"foreground"`
			Background string `json:"background"`
		} `json:"theme"`
		Columnspace   int  `json:"columnspace"`
		Margintop     int  `json:"margintop"`
		Marginbottom  int  `json:"marginbottom"`
		Marginleft    int  `json:"marginleft"`
		Marginright   int  `json:"marginright"`
		Headspace     int  `json:"headspace"`
		Footspace     int  `json:"footspace"`
		HeadFirstOnly bool `json:"head-first-only"`
		Spacing       struct {
			Title  float64 `json:"title"`
			Lyrics float64 `json:"lyrics"`
			Chords float64 `json:"chords"`
			Grid   float64 `json:"grid"`
			Tab    int     `json:"tab"`
			Toc    float64 `json:"toc"`
			Empty  int     `json:"empty"`
		} `json:"spacing"`
		Chorus struct {
			Indent int `json:"indent"`
			Bar    struct {
				Offset int    `json:"offset"`
				Width  int    `json:"width"`
				Color  string `json:"color"`
			} `json:"bar"`
			Tag    string `json:"tag"`
			Recall struct {
				Choruslike bool   `json:"choruslike"`
				Tag        string `json:"tag"`
				Type       string `json:"type"`
				Quote      bool   `json:"quote"`
			} `json:"recall"`
		} `json:"chorus"`
		Labels struct {
			Width   string `json:"width"`
			Align   string `json:"align"`
			Comment string `json:"comment"`
		} `json:"labels"`
		Chordscolumn          int    `json:"chordscolumn"`
		Capoheading           string `json:"capoheading"`
		TitlesDirectiveIgnore bool   `json:"titles-directive-ignore"`
		Diagrams              struct {
			Show      string  `json:"show"`
			Width     int     `json:"width"`
			Height    int     `json:"height"`
			Vcells    int     `json:"vcells"`
			Linewidth float64 `json:"linewidth"`
			Hspace    float64 `json:"hspace"`
			Vspace    int     `json:"vspace"`
		} `json:"diagrams"`
		Kbdiagrams struct {
			Show      string  `json:"show"`
			Width     int     `json:"width"`
			Height    int     `json:"height"`
			Keys      int     `json:"keys"`
			Base      string  `json:"base"`
			Linewidth float64 `json:"linewidth"`
			Pressed   string  `json:"pressed"`
			Hspace    float64 `json:"hspace"`
			Vspace    float64 `json:"vspace"`
		} `json:"kbdiagrams"`
		EvenOddPages   int `json:"even-odd-pages"`
		PagealignSongs int `json:"pagealign-songs"`
		Formats        struct {
			Default struct {
				Title    []string `json:"title"`
				Subtitle []string `json:"subtitle"`
				Footer   []string `json:"footer"`
			} `json:"default"`
			Title struct {
				Title    []string `json:"title"`
				Subtitle []string `json:"subtitle"`
				Footer   []string `json:"footer"`
			} `json:"title"`
			First struct {
				Footer []string `json:"footer"`
			} `json:"first"`
		} `json:"formats"`
		SplitMarker []string      `json:"split-marker"`
		Fontdir     []interface{} `json:"fontdir"`
		Fontconfig  struct {
			Times struct {
				string     `json:""`
				Bold       string `json:"bold"`
				Italic     string `json:"italic"`
				Bolditalic string `json:"bolditalic"`
			} `json:"times"`
			Helvetica struct {
				string      `json:""`
				Bold        string `json:"bold"`
				Oblique     string `json:"oblique"`
				Boldoblique string `json:"boldoblique"`
			} `json:"helvetica"`
			Courier struct {
				string     `json:""`
				Bold       string `json:"bold"`
				Italic     string `json:"italic"`
				Bolditalic string `json:"bolditalic"`
			} `json:"courier"`
			Dingbats struct {
				string `json:""`
			} `json:"dingbats"`
		} `json:"fontconfig"`
		Fonts struct {
			Title struct {
				Name string `json:"name"`
				Size int    `json:"size"`
			} `json:"title"`
			Text struct {
				Name string `json:"name"`
				Size int    `json:"size"`
			} `json:"text"`
			Chord struct {
				Name string `json:"name"`
				Size int    `json:"size"`
			} `json:"chord"`
			Chordfingers struct {
				Name        string `json:"name"`
				Size        int    `json:"size"`
				Numbercolor string `json:"numbercolor"`
			} `json:"chordfingers"`
			Comment struct {
				Name       string `json:"name"`
				Size       int    `json:"size"`
				Background string `json:"background"`
			} `json:"comment"`
			CommentItalic struct {
				Name string `json:"name"`
				Size int    `json:"size"`
			} `json:"comment_italic"`
			CommentBox struct {
				Name  string `json:"name"`
				Size  int    `json:"size"`
				Frame int    `json:"frame"`
			} `json:"comment_box"`
			Tab struct {
				Name string `json:"name"`
				Size int    `json:"size"`
			} `json:"tab"`
			Toc struct {
				Name string `json:"name"`
				Size int    `json:"size"`
			} `json:"toc"`
			Grid struct {
				Name string `json:"name"`
				Size int    `json:"size"`
			} `json:"grid"`
		} `json:"fonts"`
		Outlines []struct {
			Fields   []string `json:"fields"`
			Label    string   `json:"label"`
			Line     string   `json:"line"`
			Collapse bool     `json:"collapse"`
			Letter   int      `json:"letter"`
			Fold     bool     `json:"fold"`
		} `json:"outlines"`
		Showlayout bool `json:"showlayout"`
		Csv        struct {
			Songsonly bool `json:"songsonly"`
		} `json:"csv"`
	} `json:"pdf"`
	Chordpro struct {
		Chorus struct {
			Recall struct {
				Tag   string `json:"tag"`
				Type  string `json:"type"`
				Quote bool   `json:"quote"`
			} `json:"recall"`
		} `json:"chorus"`
	} `json:"chordpro"`
	HTML struct {
		Styles struct {
			Display string `json:"display"`
			Print   string `json:"print"`
		} `json:"styles"`
	} `json:"html"`
	A2Crd struct {
		InferTitles bool   `json:"infer-titles"`
		Classifier  string `json:"classifier"`
		Tabstop     int    `json:"tabstop"`
	} `json:"a2crd"`
	Parser struct {
		Preprocess struct {
			All       []interface{} `json:"all"`
			Directive []interface{} `json:"directive"`
			Songline  []interface{} `json:"songline"`
		} `json:"preprocess"`
	} `json:"parser"`
	Debug struct {
		Config  int `json:"config"`
		Fonts   int `json:"fonts"`
		Images  int `json:"images"`
		Layout  int `json:"layout"`
		Meta    int `json:"meta"`
		Mma     int `json:"mma"`
		Spacing int `json:"spacing"`
		Song    int `json:"song"`
		Abc     int `json:"abc"`
		Ly      int `json:"ly"`
	} `json:"debug"`
}
