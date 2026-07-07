package main

import (
	"fmt"
	"github.com/charm1092/YT_video_search/internal/subtitles"
	"github.com/charm1092/YT_video_search/internal/search"
)

func main()  {
	url := "https://www.youtube.com/watch?v=f2kvTd54uMw"

	path_json, err := subtitles.Fetcher(url)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	
	segments, err := subtitles.ParseFile(path_json)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	
	if err := subtitles.SaveSegments("subs.txt", segments); err != nil {
		fmt.Println("error: ", err)
		return
	}

	var s string
	fmt.Scan(&s)

	answer, err := search.FindWordV1("subs.txt", s)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println(answer)
}