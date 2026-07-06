package subtitles

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type SubtitleSegment struct {
	StartMs int
	EndMs   int
	Text    string
}

// элемент массива segs
type JSONSeg struct {
	Text string `json:"utf8"`
}

// элемент массива events
type JSONEvent struct {
	StartFrom   int       `json:"tStartMs"`
	DurationSeg int       `json:"dDurationMs"`
	Segs        []JSONSeg `json:"segs"`
}

// массив events
type JSONFile struct {
	Events []JSONEvent `json:"events"`
}

func ParseFile(path string) ([]SubtitleSegment, error) {
	// открыть файл
	file, err := os.Open(path)
	if err != nil {
		return []SubtitleSegment{}, err
	}
	defer file.Close()

	var data JSONFile
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return []SubtitleSegment{}, err
	}

	// пройтись по data.Events
	// склеить Segs.Text
	segments := make([]SubtitleSegment, 0, len(data.Events))

	for _, event := range data.Events {
		text := ""

		for _, seg := range event.Segs {
			text += seg.Text
		}

		segments = append(segments, SubtitleSegment{
			StartMs: event.StartFrom,
			Text:    text,
		})
	}

	return segments, nil
}

func SaveSegments(path_to_save string, segments []SubtitleSegment) error {
	file, err := os.Create(path_to_save)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, segment := range segments {
		_, err := fmt.Fprintf(writer, "%d\t%s\n", segment.StartMs, segment.Text)

		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
