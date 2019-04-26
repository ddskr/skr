package utils

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func TestUpSearchFile(t *testing.T) {
	reg := regexp.MustCompile(`^mpd.fifo$`)
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	if upSearchFile(reg, dir) == "/tmp" {
		t.Log("pass")
		return
	}
	t.Error(`Please make sure that the "mpd.fifo" is in the "/tmp"`)
}

func ExampleupSearchFile(t *testing.T) {
	reg := regexp.MustCompile(`^mpd.fifo$`)
	dir, _ := os.Getwd()
	println(upSearchFile(reg, dir))

	// Output:
	//
	// /tmp
}

func ExampleSearchModFile() {
	file := SearchModFile()
	println(file)

	// Output:
	//
	//
}
