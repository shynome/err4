package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dietsche/rfsnotify"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
	"github.com/shynome/err4/pkg/transpile"
	"gopkg.in/fsnotify.v1"
)

func watch(root string) (err error) {
	defer err2.Handle(&err)

	w := try.To1(rfsnotify.NewWatcher())
	defer w.Close()

	root = try.To1(filepath.Abs(root))
	try.To(w.AddRecursive(root))

	for {
		select {
		case ev, ok := <-w.Events:
			if !ok {
				return
			}
			go func(ev fsnotify.Event) {
				if strings.HasSuffix(ev.Name, "_err4.go") {
					return
				}
				if !strings.HasSuffix(ev.Name, ".go") {
					return
				}

				defer err2.Catch(func(err error) {
					log.Println("handle change failed. err", err)
				})
				switch ev.Op {
				case fsnotify.Rename:
					fallthrough
				case fsnotify.Write:
					// log.Println("file", ev.Name)
					b, err4file, err := transpile.Transform(ev.Name, nil)
					if err != nil {
						// hidden err
						return
					}
					if err4file {
						p := err4path(ev.Name)
						try.To(os.WriteFile(p, b.Bytes(), os.ModePerm))
					}
				case fsnotify.Remove:
					p := err4path(ev.Name)
					try.To(os.Remove(p))
				}
			}(ev)
		case err, ok := <-w.Errors:
			if !ok {
				return nil
			}
			log.Println("error:", err)
		}
	}

}

func err4path(f string) string {
	return strings.TrimSuffix(f, ".go") + "_err4.go"
}
