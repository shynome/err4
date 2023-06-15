package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"log"
	"os"
	"strings"

	"github.com/shynome/err4/pkg/transpile"
	"golang.org/x/tools/go/ast/astutil"
)

var args struct {
	input string
	out   string
}

func init() {
	flag.StringVar(&args.input, "f", "", "the input file path")
	flag.StringVar(&args.out, "o", "", "the out file path")
}

func main() {
	flag.Parse()

	if args.input == "" {
		log.Fatal("input file is required")
		return
	}

	if args.out == "" {
		if args.input == "-" {
			args.out = "-"
		} else {
			args.out = strings.TrimSuffix(args.input, ".go") + "_err4.go"
		}
	}

	var src io.Reader = nil
	if args.input == "-" {
		src = os.Stdin
	}

	var output io.Writer
	if args.out == "-" {
		output = os.Stdout
	} else {
		if f, err := os.OpenFile(args.out, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm); err != nil {
			log.Fatal(err)
			return
		} else {
			output = f
		}
	}

	if err := gen(args.input, src, output); err != nil {
		log.Fatal(err)
	}

}

func gen(filepath string, input any, output io.Writer) (err error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, input, parser.ParseComments)
	if err != nil {
		return
	}
	astutil.Apply(f, nil, func(c *astutil.Cursor) bool {
		n := c.Node()
		switch x := n.(type) {
		case *ast.AssignStmt:
			transpile.ChangeErr4AssignStmt(x, c)
		case *ast.File:
			comments := make([]*ast.CommentGroup, len(x.Comments)+1)
			comments[0] = &ast.CommentGroup{
				List: []*ast.Comment{
					{Text: `// Code generated github.com/shynome/err4 DO NOT EDIT`},
				},
			}
			copy(comments[1:], x.Comments)
			x.Comments = comments
			transpile.RepalceErr4BuildTag(x)
		}
		return true
	})
	astutil.AddImport(fset, f, "github.com/shynome/err4")

	if err = printer.Fprint(output, fset, f); err != nil {
		return
	}
	return
}
