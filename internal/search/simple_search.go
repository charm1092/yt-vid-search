package search

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type FoundSegment struct {
	Start int
	End   int
	Text  string
}

func NewFoundSegment(start, end int, text string) *FoundSegment {
	return &FoundSegment{
		Start: start,
		End: end,
		Text: text,
	}
}

func ConvertMsToNormalTime(msTime int) string {
	t := time.Duration(msTime) * time.Millisecond

	hours := int(t.Hours())
	minutes := int(t.Minutes()) % 60
	seconds := int(t.Seconds()) % 60

	if hours > 0 {
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	}

	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

func FindWordInFile(path string, word string) ([]FoundSegment, error) {
	word = strings.ToLower(word)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	slice_of_starts := make([]FoundSegment, 0, 100)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		textLow := strings.ToLower(parts[2])

		

		if strings.Contains(textLow, word) {
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, err
			}

			end, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, err
			}

			segment := FoundSegment{
				Start: start,
				End: end,
				Text: parts[2],
			}
			
			slice_of_starts = append(slice_of_starts, segment)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(slice_of_starts) == 0 {
		fmt.Println("Ничего не найдено")
		return nil, nil
	}
	return slice_of_starts, nil
}

func FindContextBySelectedV1(selectedSegment FoundSegment, path string) ([]FoundSegment, error) {
	allSegments := make([]FoundSegment, 0, 100)
	segmentContext := make([]FoundSegment, 5)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")

		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}

		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		temp := FoundSegment{
			Start: start,
			End: end,
			Text: parts[2],
		}
		
		allSegments = append(allSegments, temp)
	}

	for i := 0; i < len(allSegments); i++ {
		if allSegments[i].Start == selectedSegment.Start {
			for j := 0; j < len(segmentContext); j++ {
				segmentContext[j] = allSegments[i-2+j]
			}
		}
	}

	return segmentContext, nil
}
