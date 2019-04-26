package utils

import (
	"fmt"
	"testing"
)

func cmp(a, b []string) bool {
	l := len(a)
	if l != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func checkParsePackageRepo(t *testing.T, i string, v []string, official bool) {
	repoPath, remote, author, repo, dirPath, auto := ParsePackageRepo(i)
	if repoPath == v[0] &&
		remote == v[1] &&
		author == v[2] &&
		repo == v[3] &&
		cmp(dirPath, v[4:]) && auto == official {
		t.Log(i, "pass")
		return
	}
	t.Error(i, "error ->", repoPath, remote, author, repo, dirPath)
}

func TestParsePackageRepo(t *testing.T) {
	checkParsePackageRepo(t, "os", []string{
		"https://github.com/ddskr/skrlang.git",
		"github.com", "ddskr", "skrlang",
		"github.com", "ddskr", "skrlang", "os",
	}, true)

	checkParsePackageRepo(t, "path", []string{
		"https://github.com/ddskr/skrlang.git",
		"github.com", "ddskr", "skrlang",
		"github.com", "ddskr", "skrlang", "path",
	}, true)

	checkParsePackageRepo(t, "path/filepath", []string{
		"https://github.com/ddskr/skrlang.git",
		"github.com", "ddskr", "skrlang",
		"github.com", "ddskr", "skrlang", "path", "filepath",
	}, true)

	checkParsePackageRepo(t, "github.com/ddosakura/skrlangx/net", []string{
		"https://github.com/ddosakura/skrlangx.git",
		"github.com", "ddosakura", "skrlangx",
		"github.com", "ddosakura", "skrlangx", "net",
	}, false)

	checkParsePackageRepo(t, "gitee.com/xxx/yyy", []string{
		"https://gitee.com/xxx/yyy.git",
		"gitee.com", "xxx", "yyy",
		"gitee.com", "xxx", "yyy",
	}, false)
}

func ExampleLocalPackageSearch() {
	dirs := LocalPackageSearch([]string{"/root/skr/src"}, "my_local_pacakge/xxx")
	// pretty.Println(dirs)
	fmt.Println(dirs[0])

	// Output:
	//
	// /root/skr/src/my_local_pacakge/xxx
}
