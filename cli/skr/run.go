package main

import (
	"github.com/ddskr/skr/cli"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"
)

var (
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "parse & run",
		Long:  `Parse *.skr to *.ll and run *.ll by lli.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cli.CFG.Input = args[0]
			// TODO:
			pretty.Println(cli.CFG)
		},
	}
)

func init() {
	rootCmd.AddCommand(runCmd)
}
