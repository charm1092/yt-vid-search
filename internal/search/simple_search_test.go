package search

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConvertMsToNormalTime(t *testing.T)  {
	got := ConvertMsToNormalTime(129022)
	want := "02:09"

	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestFindWordV1(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "subs.txt")

	data := "100\t200\thello world\n300\t400\tanother line\n"

	err := os.WriteFile(path, []byte(data), 0644)
	if err != nil {
		t.Fatal(err)
	}

	got, start, err := FindWordV1(path, "паренёк")
	if err != nil {
		t.Fatal(err)
	}

	if got != "yes" {
		t.Fatalf("want yes, got %q", got)
	}

	if start != 100 {
		t.Fatalf("want start 100, got %d", start)
	}
}