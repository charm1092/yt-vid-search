package cache

import (
	"os"
	"path/filepath"
	"sort"
	"time"
)

type CacheEntry struct {
	Path     string
	LastUsed time.Time
}

func Touch(cacheDir string) error {
	now := time.Now()
	if err := os.Chtimes(cacheDir, now, now); err != nil {
		return err
	}

	return nil
}

func Cleanup(rootDir string, limit int) error {
	folders := make([]CacheEntry, 0, 5)

	entries, err := os.ReadDir(rootDir)
	if err != nil {
		return err
	}

	for _, f := range entries {
		if !f.IsDir() {
			continue
		}

		info, err := f.Info()
		if err != nil {
			return err
		}

		folders = append(folders, CacheEntry{
			Path: filepath.Join(rootDir, f.Name()),
			LastUsed: info.ModTime(),
		})
	}

	sort.Slice(folders, func(i, j int) bool {
    	return folders[i].LastUsed.Before(folders[j].LastUsed)
	})

	extra := len(folders) - limit
	if extra > 0 {
		for i, _ := range folders[:extra] {
			if err := os.RemoveAll(folders[i].Path); err != nil {
				return err
			}
		}
	}
	return nil
}
