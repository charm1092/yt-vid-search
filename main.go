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
	var urlRaw string
	fmt.Println("введите ссылку видео: ")
	fmt.Scan(&urlRaw)
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

	fmt.Println("файлы готовы к работе")

	var s string
	fmt.Scan(&s)

	answer, err := search.FindWordInFile(segmentsPath, s)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for i, v := range answer {
		fmt.Printf("%d) %s - %s -- %s\n", i+1, search.ConvertMsToNormalTime(v.Start), search.ConvertMsToNormalTime(v.End), v.Text)
	}

	fmt.Println("выберете нужный номер фрагмента: ")

	var choice int
	fmt.Scan(&choice)

	if choice < 1 || choice > len(answer) || len(answer) == 0 {
		return
	} 
	selected := answer[choice-1]

	fmt.Println("результат поиска: ")
	result, err := search.FindContextBySelectedV1(selected, segmentsPath)
	if err != nil {
		fmt.Println("error: ошибка поиска контекста, ", err)
		return
	}

	for _, s := range result {
		fmt.Println(s)
	}
	
}
