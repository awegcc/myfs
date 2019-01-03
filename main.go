package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func usage() {
	progName := filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", progName)
	fmt.Fprintf(os.Stderr, "  %s file.zip mountpoint\n", progName)
	flag.PrintDefaults()
}

func main() {
	progName := filepath.Base(os.Args[0])
	log.SetFlags(0)
	log.SetPrefix(progName + ": ")

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 2 {
		usage()
		os.Exit(2)
	}
	usage()
}
