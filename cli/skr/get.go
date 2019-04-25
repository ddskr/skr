package main

import (
	"os"
	"os/exec"

	"github.com/ddskr/skr/cli"
	"github.com/ddskr/skr/utils"
	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "get source code",
		Long:  `Download source code by git.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cli.CFG.Input = args[0]

			repoPath, rootPath /*, repo, dirPath*/, _, _ := utils.ParsePackageRepo(cli.CFG.Input)
			if err := os.MkdirAll(rootPath, 0755); err != nil {
				panic(err)
			}

			c := exec.Command("git", "clone", repoPath)
			c.Dir = rootPath
			utils.Exec(c, true)
		},
	}
)

func init() {
	rootCmd.AddCommand(getCmd)
}
