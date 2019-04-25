package utils

import (
	"path"
	"strings"

	"github.com/ddskr/skr/cli"
)

// ParsePackageRepo to get info
// TODO: multi dir(repo)
// need to change:
//     1. parse "xxx/yyy"
// 	       -> https://github.com/ddskr/skrlang/xxx/yyy
//         -> file://path/to/src/xxx/yyy
//     2. cli.CFG.DefaultSrcPath() -> $SKRPATH (PATH[0]|PATH[1]|PATH[2]|...)
func ParsePackageRepo(pkg string) (repoPath, rootPath, repo, dirPath string) {
	var (
		part   = strings.Split(pkg, "/")
		remote = "github.com"
		user   = "ddskr"
	)
	repo = "skrlang"
	switch len(part) {
	case 1, 2:
		part = append([]string{remote, user, repo}, part...)
	default:
		remote = part[0]
		user = part[1]
		repo = part[2]
	}
	repoPath = "https://" +
		strings.Join([]string{remote, user, repo}, "/") + ".git"
	rootPath = path.Join(cli.CFG.DefaultSrcPath(), remote, user)
	dirPath = path.Join(append([]string{cli.CFG.DefaultSrcPath()}, part[:len(part)]...)...)
	return
}
