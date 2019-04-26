package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/ddskr/skr/cli"
)

func init() {
	// `utils` use the cfg in `cli`, so can't move this to `cli`
	cli.CFG.Mod = SearchModFile()
}

// SearchModFile search dir of "skr.mod" or "go.mod"
func SearchModFile() string {
	reg := regexp.MustCompile(`^((skr)|(go))\.mod$`)

	// this is dir of program
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	// this is work dir
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return upSearchFile(reg, dir)
}

func upSearchFile(reg *regexp.Regexp, dir string) string {
	// println("finding in", dir)
	if dir == "" || dir == "/" || dir == filepath.Dir(dir) {
		return ""
	}
	list, err := ioutil.ReadDir(dir)
	if err != nil {
		return ""
	}
	for _, v := range list {
		if !v.IsDir() && reg.MatchString(v.Name()) {
			return filepath.Join(dir, v.Name())
		}
	}
	return upSearchFile(reg, filepath.Dir(dir))
}
