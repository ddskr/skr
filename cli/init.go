package cli

import (
	"github.com/joho/godotenv"
	"github.com/timest/env"
)

// SkrConfig for tools
type SkrConfig struct {
	Bin   string   `default:"/root/skr/bin" env:"SKRBIN"`
	Cache string   `default:"/root/skr/cache" env:"SKRCACHE"`
	Path  []string `default:"/root/skr" env:"SKRPATH" slice_sep:":"`
	Mod   string   `default:"" env:"SKRMOD"`

	Input  string
	Output string
}

var (
	// CFG from env(s) & arg(s)
	CFG = new(SkrConfig)
)

func init() {
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
}
