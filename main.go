package main

import (
	"fmt"
	"net/url"
	"github.com/charm1092/YT_video_search/internal/search"
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

	segmentsPath, err := search.PrepareSubtitles(urlRaw, videoID)
	if err != nil {
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
	result, err := search.FindContextBySelected(selected, segmentsPath)
	if err != nil {
		fmt.Println("error: ошибка поиска контекста, ", err)
		return
	}

	fmt.Println(result)
}