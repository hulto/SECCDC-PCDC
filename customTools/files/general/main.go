package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type config struct {
	ext     string
	size    int64
	list    bool
	hours   int64
	curTime time.Time
}

func main() {
	root := flag.String("d", ".", "Root directory to start search from.")
	list := flag.Bool("l", false, "List found files only. (No other actions performed)")
	ext := flag.String("e", "", "Only locate files with this extension.")
	size := flag.Int64("s", -1, "Only locate files >= to this size (in bytes).")
	file := flag.String("f", "", "Write output to this file.")
	hours := flag.Int64("t", -1, "Number of hours before the current time that the file was created/modified after.")
	flag.Parse()

	var out io.Writer
	var err error

	if *file == "" {
		out = os.Stdout
	} else {
		out, err = os.Create(*file)
		if err != nil {
			log.Fatal("Could not create file.", err)
		}
		defer out.(*os.File).Close()
	}

	c := config{
		ext:     *ext,
		size:    *size,
		list:    *list,
		hours:   *hours,
		curTime: time.Now(),
	}

	if err := run(*root, out, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, cfg config) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filterOut(path, cfg.ext, cfg.size, cfg.curTime, cfg.hours, info) {
			return nil
		}

		if cfg.list {
			return listFile(path, out, info)
		}

		return listFile(path, out, info)
	})
}

func filterOut(path, ext string, minSize int64, curTime time.Time, hours int64, info os.FileInfo) bool {
	if info.IsDir() {
		return true
	}

	if (ext != "" && filepath.Ext(path) != ext) || (info.Size() < minSize) || (info.ModTime().Before(curTime.Add(time.Duration(-hours)*time.Hour)) && hours > 0) {
		return true
	}

	return false
}

func listFile(path string, out io.Writer, info os.FileInfo) error {
	_, err := fmt.Fprintf(out, "%d %s\n\t%s\n", info.Size(), info.ModTime(), path)
	return err
}
