package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/shynome/err4/pkg/transpile"
	"golang.org/x/sync/errgroup"
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
	flag.BoolVar(&args.err4, "err4", false, "transform weather the content include ierr build tag")
}

func main() {
	flag.Parse()

	if args.input == "" && len(os.Args) > 1 {
		args.input = os.Args[1]
	}

	if finfo, err := os.Stat(args.input); err != nil {
		log.Fatalln(err)
	} else if !finfo.IsDir() {
		if err := gen(args.input, args.out); err != nil {
			log.Fatal(err)
		}
		return
	}

	entries, err := os.ReadDir(args.input)
	if err != nil {
		log.Fatalln(err)
	}
	eg := new(errgroup.Group)
	for _, _f := range entries {
		f := _f
		eg.Go(func() error {
			input := f.Name()
			if f.IsDir() ||
				strings.HasSuffix(input, "_ierr.go") ||
				strings.HasSuffix(input, "_ierr_test.go") {
				return nil
			}
			input = filepath.Join(args.input, input)
			err := gen(input, "")
			if errors.Is(err, ErrNotErr4GoFile) {
				return nil
			}
			return err
		})
	}
	if err := eg.Wait(); err != nil {
		log.Fatalln(err)
	}

}

func gen(input, out string) error {
	if input == "" {
		return fmt.Errorf("input file is required")
	}

	if out == "" {
		if input == "-" {
			out = "-"
		} else {
			out = err4Path(input)
		}
	}

	var src io.Reader = nil
	if input == "-" {
		src = os.Stdin
	}

	output, err4file, err := transpile.Transform(input, src)
	if err != nil {
		return err
	}

	if !args.err4 {
		args.err4 = err4file
	}
	if !args.err4 {
		return ErrNotErr4GoFile
	}

	if out == "-" {
		io.Copy(os.Stdout, &output)
		return nil
	}

	if err := os.WriteFile(out, output.Bytes(), os.ModePerm); err != nil {
		return err
	}
	return nil
}

var ErrNotErr4GoFile = fmt.Errorf("file content don't include ierr build tag")

func err4Path(f string) string {
	if strings.HasSuffix(f, "_test.go") {
		return strings.TrimSuffix(f, "_test.go") + "_ierr_test.go"
	}
	return strings.TrimSuffix(f, ".go") + "_ierr.go"
}
