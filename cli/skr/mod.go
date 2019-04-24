package main

import (
	"github.com/ddskr/skr/cli"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"
)

var (
	modCmd = &cobra.Command{
		Use:   "mod",
		Short: "module tool",
		Long:  `Manage packages.`,
	}

	modInitCmd = &cobra.Command{
		Use:   "init",
		Short: "init module",
		Long:  `Generate skr.mod in git repo`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO:
			pretty.Println(cli.CFG)
		},
	}
)

func init() {
	modCmd.AddCommand(modInitCmd)

	rootCmd.AddCommand(modCmd)
	rootCmd.AddCommand(moduleCmd)
}

var (
	moduleCmd = &cobra.Command{
		Use:   "module",
		Short: "about module",
		Long: `A module is a collection of packages.

The package in skr -> The module in llvm
The module in skr -> *.a (collection of *.o)
`}
)
