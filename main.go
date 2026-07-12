package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/charm1092/YT_video_search/internal/search"
	"github.com/charm1092/YT_video_search/internal/subtitles"
)

func main() {
	urlRaw := "https://www.youtube.com/watch?v=f2kvTd54uMw"
	parsedURL, err := url.Parse(urlRaw)
	if err != nil {
		fmt.Println("некорректная ссылка")
		return
	}

	videoID := parsedURL.Query().Get("v")
	if videoID == "" {
		fmt.Println("в ссылке нет ID видео")
		return
	}

	cacheDir := filepath.Join("data", "video", videoID)
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		fmt.Println("error:", err)
		return
	}

	jsonBasePath := filepath.Join(cacheDir, "subtitles")
	jsonPath := jsonBasePath + ".ru.json3"
	segmentsPath := filepath.Join(cacheDir, "subs.txt")

	if _, err := os.Stat(segmentsPath); os.IsNotExist(err) {
		if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
			jsonPath, err = subtitles.Fetcher(urlRaw, jsonBasePath)
			if err != nil {
				fmt.Println("error:", err)
				return
			}
		} else if err != nil {
			fmt.Println("error:", err)
			return
		}

		segments, err := subtitles.ParseFile(jsonPath)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		if err := subtitles.SaveSegments(segmentsPath, segments); err != nil {
			fmt.Println("error:", err)
			return
		}
	} else if err != nil {
		fmt.Println("error:", err)
		return
	}

	var s string
	fmt.Scan(&s)

	answer, err := search.FindWordV2(segmentsPath, s)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(answer)
}
