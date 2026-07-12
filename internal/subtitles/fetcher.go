package subtitles

import (
	"os/exec"
)

func Fetcher(url, outputPath string) (string, error) {
	
	cmd := exec.Command("yt-dlp",
		"--write-auto-subs",
		"--sub-langs", "ru",
		"--sub-format", "json3",
		"--skip-download",
		"-o", outputPath,
		url,
	)

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return outputPath + ".ru.json3", nil
}
