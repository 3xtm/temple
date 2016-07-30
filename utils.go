package main

import (
	"io"
	"log"
	"os"
)

func safeOpen(path string) io.Reader {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Unable to open %s: %s", path, err)
	}
	return file
}
