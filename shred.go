package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"os"
)

func shredFile(filename string, passes int) error {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	fileSize := info.Size()
	data := make([]byte, fileSize)

	for i := 0; i < passes; i++ {
		if _, err := rand.Read(data); err != nil {
			return fmt.Errorf("failed to generate random data: %w", err)
		}
		if _, err := file.WriteAt(data, 0); err != nil {
			return fmt.Errorf("failed to write random data: %w", err)
		}
	}

	return nil
}

func removeFile(filename string) error {
	if err := os.Remove(filename); err != nil {
		return fmt.Errorf("failed to remove file: %w", err)
	}
	return nil
}

func main() {
	var passes int
	flag.IntVar(&passes, "passes", 3, "number of times to overwrite the file")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] filename\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	filename := flag.Arg(0)

	if err := shredFile(filename, passes); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	if err := removeFile(filename); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("File has been securely deleted.")
}
