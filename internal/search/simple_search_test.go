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

	got, err := FindWordV1(path, "world")
	if err != nil {
		t.Fatal(err)
	}

	if got != "yes, 00:00" {
		t.Fatalf("want yes, got %q", got)
	}
}

func TestFindWordV1NotFound(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "subs.txt")

	data := "100\t200\thello world\n300\t400\tanother line\n"

	err := os.WriteFile(path, []byte(data), 0644)
	if err != nil {
		t.Fatal(err)
	}

	got, err := FindWordV1(path, "missing")
	if err != nil {
		t.Fatal(err)
	}

	if got != "no" {
		t.Fatalf("want no, got %q", got)
	}
}
