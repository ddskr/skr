package utils

import (
	"path"
	"strings"
)

// ParsePackageRepo to get info
func ParsePackageRepo(pkg string) (repoPath, remote, user, repo string, dirPath []string, auto bool) {
	var (
		part = strings.Split(pkg, "/")
	)
	remote = "github.com"
	user = "ddskr"
	repo = "skrlang"
	switch len(part) {
	case 1, 2:
		part = append([]string{remote, user, repo}, part...)
		auto = true
	default:
		remote = part[0]
		user = part[1]
		repo = part[2]
	}
	repoPath = "https://" +
		strings.Join([]string{remote, user, repo}, "/") + ".git"
	// rootPath = path.Join(cli.CFG.DefaultSrcPath(), remote, user)
	// dirPath = path.Join(append([]string{cli.CFG.DefaultSrcPath()}, part[:len(part)]...)...)
	dirPath = part[:len(part)]
	return
}

// LocalPackageSearch search package in local src dir
func LocalPackageSearch(dirs []string, pkg string) []string {
	_, _, _, _, dirPath, auto := ParsePackageRepo(pkg)
	if auto {
		dirPath = dirPath[2:]
	} else {
		dirPath = append([]string{""}, dirPath...)
	}
	ret := make([]string, 0, len(dirs))
	for _, v := range dirs {
		dirPath[0] = v
		ret = append(ret, path.Join(dirPath...))
	}
	return ret
}
