package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/timest/env"
)

// SkrConfig for tools
type SkrConfig struct {
	Bin   string   `default:"/root/skr/bin" env:"SKRBIN"`
	Cache string   `default:"/root/skr/cache" env:"SKRCACHE"`
	Path  []string `default:"/root/skr" env:"SKRPATH" slice_sep:":"`
	Mod   string   `default:"" env:"SKRMOD"`

	Output string
}

var (
	// CFG from env(s) & arg(s)
	CFG = new(SkrConfig)

	rootCmd = &cobra.Command{
		Use:   "skrc",
		Short: "Skrlang Parser",
		Long:  `This tool is used to parse *.skr to *.ll.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			parser(args[0])
		},
	}
)

func main() {
	//if e := godotenv.Load(); e != nil {
	//	log.Fatal("Error loading .env file")
	//}
	godotenv.Load()

	env.IgnorePrefix()
	err := env.Fill(CFG)
	if err != nil {
		panic(err)
	}
	// pretty.Println(CFG)

	fs := rootCmd.PersistentFlags()
	fs.StringVarP(&CFG.Output, "output", "o", "pkg.ll", "output file")
	rootCmd.Execute()
}
