package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"github.com/shynome/err4/pkg/transpile"
)

var args struct {
	input string
	out   string
	err4  bool

	watch bool
	root  string
}

func init() {
	flag.StringVar(&args.input, "f", "", "the input file path")
	flag.StringVar(&args.out, "o", "", "the out file path")
	flag.BoolVar(&args.err4, "err4", false, "transform weather the content include err4 build tag")
}

func main() {
	flag.Parse()

	if args.input == "" {
		log.Fatal("input file is required")
	}

	if args.out == "" {
		if args.input == "-" {
			args.out = "-"
		} else {
			args.out = err4path(args.input)
		}
	}

	var src io.Reader = nil
	if args.input == "-" {
		src = os.Stdin
	}

	output, err4file, err := transpile.Transform(args.input, src)
	if err != nil {
		log.Fatal(err)
	}

	if !args.err4 {
		args.err4 = err4file
	}
	if !args.err4 {
		log.Println("file content don't include err4 build tag")
		return
	}

	if args.out == "-" {
		io.Copy(os.Stdout, &output)
		return
	}

	if err := os.WriteFile(args.out, output.Bytes(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func err4path(f string) string {
	return strings.TrimSuffix(f, ".go") + "_err4.go"
}
