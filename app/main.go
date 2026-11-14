package main

import (
	"fmt"

	httpapi "github.com/charm1092/YT_video_search/internal/http"
)

func main()  {
	httpHandler := httpapi.NewHTTPHandler()
	httpServer := httpapi.NewHTTPServer(httpHandler)

	if err := httpServer.StarServer(); err != nil {
		fmt.Println("не удалось запустить http-сервер: ", err)
		return
	}
}