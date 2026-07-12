package search

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

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

func FindWordV2(path string, word string) ([]string, error) {
	word = strings.ToLower(word)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	slice_of_starts := make([]string, 0, 100)
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

			conv_time := ConvertMsToNormalTime(start)
			slice_of_starts = append(slice_of_starts, conv_time)
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