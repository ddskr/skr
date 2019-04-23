package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ddskr/skr/cli"
	"github.com/ddskr/skr/grammar"
	"github.com/ddskr/skr/grammar/ast"
	"github.com/ddskr/skr/grammar/selector"
	"github.com/kr/pretty"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func parser(cfg *cli.SkrConfig) {
	mod := cfg.Path[0] + "/src/" + cfg.Input
	files := make([]string, 0)

	roots := make([]*ast.Node, 0)
	filepath.Walk(mod, func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fmt.Println(p, path.Ext(p))
		if path.Ext(p) != ".skr" {
			return nil
		}
		buf, e := ioutil.ReadFile(p)
		if e != nil {
			panic(e)
		}
		files = append(files, info.Name())

		t, e := ast.Parse(p, string(buf), func(err grammar.SyntaxError) bool {
			pretty.Println("parser", err)
			return true
		})
		if e != nil {
			if e.Error() == "exactly one root node is expected" {
				println("source code should start with 'package <name>'")
				return nil
			}
			pretty.Println("ast.Parse", e)
			return nil
		}
		// pretty.Println(t)
		root := t.Root()
		tmp(0, root)

		roots = append(roots, root)
		return nil
	})
	return
}

func unquote(s string) string {
	a, e := strconv.Unquote(`"` + s + `"`)
	if e != nil {
		panic(e)
	}
	return a
}

/*
func expr2Value(p *generater.Pkg, e *ast.Node) value.Value {
	s := e.Text()
	s = unquote(s[1 : len(s)-1])
	return p.GlobalStr(s)
}
*/

func parserExpr(e *ast.Node) (value.Value, types.Type) {
	v, _ := strconv.Atoi(e.Text())
	t := types.I32
	return constant.NewInt(t, int64(v)), t
}

func tmp(dep int, n *ast.Node) {
	space(dep)
	fmt.Println(n.LineColumn())
	space(dep)
	fmt.Println(n.Type(), n.IsValid(),
		strings.ReplaceAll(n.Text(), "\n", "\\n"))
	for _, c := range n.Children(selector.Any) {
		tmp(dep+1, c)
	}
}

func space(dep int) {
	for dep > 0 {
		dep--
		fmt.Print("  ")
	}
}
