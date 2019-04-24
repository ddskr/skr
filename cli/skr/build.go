package main

import (
	"github.com/ddskr/skr/cli"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"
)

var (
	buildCmd = &cobra.Command{
		Use:   "build",
		Short: "parse & compile",
		Long:  `Parse *.skr to *.ll and auto-compile (*.ll by llc; *.asm by nasm).`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cli.CFG.Input = args[0]
			// TODO:
			pretty.Println(cli.CFG)
		},
	}
)

func init() {
	rootCmd.AddCommand(buildCmd)
}
