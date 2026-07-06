package main

import (
	"fmt"
	"github.com/charm1092/YT_video_search/internal/subtitles"
)

func main()  {
	url := "https://www.youtube.com/watch?v=f2kvTd54uMw"

	out, err := subtitles.Fetcher(url)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println(out)
}