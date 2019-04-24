package main

import (
	"github.com/ddskr/skr/cli"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "skrfmt",
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

func main() {
	fs := rootCmd.PersistentFlags()
	fs.StringVarP(&cli.CFG.Output, "output", "o", "pkg.ll", "output file")
	rootCmd.Execute()
}
