package main

import (
	"fmt"
	"io"
	"bytes"
	"os"
)

func countLines(reader io.Reader) (int, error) {
	buffer := make([]byte, 1024*32)
	count := 0
	separator := []byte{'\n'}

	for {
		char, err := reader.Read(buffer)
		count += bytes.Count(buffer[:char], separator)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}

	}

}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Expected 1 argument, but got: %d", len(args))
		return
	}

	path := args[0]
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open %s", path)
		return
	}

	count, err := countLines(f)
	fmt.Fprintf(os.Stdout, "%d %s\n", count, path)
	f.Close()
}
