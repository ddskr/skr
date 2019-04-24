package main

import (
	"fmt"
	"strings"

	"github.com/ddskr/skr/cli"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "skr",
		Short: "skr tools",
		Long:  `Tools for managing Skr source code.`,
	}

	envCmd = &cobra.Command{
		Use:   "env",
		Short: "env & runtime config",
		Long:  `Env(s) & Runtime Config.`,
		Run: func(cmd *cobra.Command, args []string) {
			// "SKRBIN", "SKRCACHE", "SKRPATH", "SKRMOD"
			printENV("SKRBIN", cli.CFG.Bin)
			printENV("SKRCACHE", cli.CFG.Cache)
			printENV("SKRPATH", cli.CFG.Path)
			printENV("SKRMOD", cli.CFG.Mod)
		},
	}
)

func main() {
	rootCmd.AddCommand(envCmd)

	rootCmd.Execute()
}

func printENV(k string, v interface{}) {
	k += `="`
	if V, ok := v.(string); ok {
		k += V
	}
	if V, ok := v.([]string); ok {
		k += strings.Join(V, ":")
	}
	fmt.Println(k + `"`)
}
