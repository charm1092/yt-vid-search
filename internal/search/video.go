package search

import (
	"os"
	"path/filepath"

	"github.com/charm1092/YT_video_search/internal/subtitles"
	"github.com/charm1092/YT_video_search/internal/cache"
)

func PrepareSubtitles(videoURL, videoID string) (string, error) {
	cacheDir := filepath.Join("data", "video", videoID)
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return "", err
	}

	jsonBasePath := filepath.Join(cacheDir, "subtitles")
	jsonPath := jsonBasePath + ".ru.json3"
	segmentsPath := filepath.Join(cacheDir, "subs.txt")

	if _, err := os.Stat(segmentsPath); os.IsNotExist(err) {
		if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
			jsonPath, err = subtitles.Fetcher(videoURL, jsonBasePath)
			if err != nil {
				return "", err
			}
		} else if err != nil {
			return "", err
		}

		segments, err := subtitles.ParseFile(jsonPath)
		if err != nil {
			return "", err
		}

		if err := subtitles.SaveSegments(segmentsPath, segments); err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}

	if err := cache.Touch(cacheDir); err != nil {
    	return "", err
	}
	if err := cache.Cleanup("data/video", 5); err != nil {
		return "", err
	}

	return segmentsPath, nil
}