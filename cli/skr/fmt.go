package main

import (
	"github.com/ddskr/skr/cli"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"
)

var (
	fmtCmd = &cobra.Command{
		Use:   "fmt",
		Short: "fmt source code",
		Long:  `Format *.skr (like skrfmt).`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cli.CFG.Input = args[0]
			// TODO:
			pretty.Println(cli.CFG)
		},
	}
)

func init() {
	rootCmd.AddCommand(fmtCmd)
}
