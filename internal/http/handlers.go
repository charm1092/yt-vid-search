package http

import (
	"encoding/json"
	"net/http"

	"github.com/charm1092/YT_video_search/internal/search"
)
	

type HTTPHandler struct {

}

func NewHTTPHandler() *HTTPHandler {
	return &HTTPHandler{}
}

func (h *HTTPHandler) HandlerSearch(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	videoID := params.Get("video_id")
	query := params.Get("q")
	videoURL := "https://www.youtube.com/watch?v=" + videoID

	if query == "" {
		return
	}

	segmentsPath, err := search.PrepareSubtitles(videoURL, videoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	matches, err := search.FindWordInFile(segmentsPath, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
}