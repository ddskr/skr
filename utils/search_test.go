package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func TestUpSearchFile(t *testing.T) {
	reg := regexp.MustCompile(`^mpd.fifo$`)
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	if upSearchFile(reg, dir) == "/tmp/mpd.fifo" {
		t.Log("pass")
		return
	}
	t.Error(`Please make sure that the "mpd.fifo" is in the "/tmp"`)
}

func Example_upSearchFile() {
	reg := regexp.MustCompile(`^mpd.fifo$`)
	//dir, _ := os.Getwd()
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(upSearchFile(reg, dir))

	// Output:
	//
	// /tmp/mpd.fifo
}

func ExampleSearchModFile() {
	file := SearchModFile()
	println(file)

	// Output:
	//
	//
}
