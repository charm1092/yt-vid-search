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

func FindWordV1(path string, word string) (string, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")

		if strings.Contains(parts[2], word) {
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				return "", 0, err
			}

			return "yes", start, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", 0, err
	}

	return "no", 0, nil

}