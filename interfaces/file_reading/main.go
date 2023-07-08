package main

import (
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("invalid filename passed to program")
	}

	var fileName string = os.Args[1]
	file, e := os.Open(fileName)
	if e != nil {
		panic("could not open file " + fileName)
	}

	_, e = io.Copy(os.Stdout, file)
	if e != nil {
		panic(e)
	}
}
