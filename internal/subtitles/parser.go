package subtitles

type SubtitleSegment struct {
	StartMs int
	Text    string
}

type JSONSeg struct {
	Text string `json:"utf8"`
}


type JSONEvents struct {
	StartFrom int       `json:"tStartMs"`
	Segs      []JSONSeg `json:"segs"`
}

type JSONFile struct {
	Events []JSONEvents `json:"events"`
}