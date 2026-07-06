package subtitles

import (
	"os/exec"
)

func Fetcher(url string) (string, error) {

	cmd := exec.Command("yt-dlp",
		"--write-auto-subs",
		"--sub-langs", "ru",
		"--sub-format", "json3",
		"--skip-download",
		url,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
