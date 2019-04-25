package utils

import (
	"testing"
)

func checkParsePackageRepo(t *testing.T, i string, v []string) {
	repoPath, rootPath, repo, dirPath := ParsePackageRepo(i)
	if repoPath == v[0] &&
		rootPath == v[1] &&
		repo == v[2] &&
		dirPath == v[3] {
		t.Log(i, "pass")
		return
	}
	t.Error(i, "error ->", repoPath, rootPath, repo, dirPath)
}

func TestParsePackageRepo(t *testing.T) {
	checkParsePackageRepo(t, "os", []string{
		"https://github.com/ddskr/skrlang.git",
		"/root/skr/src/github.com/ddskr",
		"skrlang",
		"/root/skr/src/github.com/ddskr/skrlang/os",
	})

	checkParsePackageRepo(t, "path", []string{
		"https://github.com/ddskr/skrlang.git",
		"/root/skr/src/github.com/ddskr",
		"skrlang",
		"/root/skr/src/github.com/ddskr/skrlang/path",
	})

	checkParsePackageRepo(t, "path/filepath", []string{
		"https://github.com/ddskr/skrlang.git",
		"/root/skr/src/github.com/ddskr",
		"skrlang",
		"/root/skr/src/github.com/ddskr/skrlang/path/filepath",
	})

	checkParsePackageRepo(t, "github.com/ddosakura/skrlangx/net", []string{
		"https://github.com/ddosakura/skrlangx.git",
		"/root/skr/src/github.com/ddosakura",
		"skrlangx",
		"/root/skr/src/github.com/ddosakura/skrlangx/net",
	})

	checkParsePackageRepo(t, "gitee.com/xxx/yyy", []string{
		"https://gitee.com/xxx/yyy.git",
		"/root/skr/src/gitee.com/xxx",
		"yyy",
		"/root/skr/src/gitee.com/xxx/yyy",
	})
}
