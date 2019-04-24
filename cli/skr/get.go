package main

import (
	"github.com/ddskr/skr/cli"
	"github.com/kr/pretty"
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
			// TODO:
			pretty.Println(cli.CFG)
		},
	}
)

func init() {
	rootCmd.AddCommand(getCmd)
}
