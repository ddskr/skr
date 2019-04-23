package main

import (
	"github.com/ddskr/skr/cli"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "skrc",
		Short: "Skrlang Parser",
		Long:  `This tool is used to parse *.skr to *.ll.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cli.CFG.Input = args[0]
			parser(cli.CFG)
		},
	}
)

func main() {
	fs := rootCmd.PersistentFlags()
	fs.StringVarP(&cli.CFG.Output, "output", "o", "pkg.ll", "output file")
	rootCmd.Execute()
}
