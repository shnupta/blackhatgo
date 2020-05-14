package main

import (
	"fmt"
	"log"
	"os"
)

// FooReader defines an io.Reader to read from stdin
type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in> ")
	return os.Stdin.Read(b)
}

// FooWriter defines and io.Writer to write to stdout
type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	// Create buffer to hold input/output
	input := make([]byte, 4096)

	// Use reader to read input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	// Use writer to write ouptu
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)

	// Useful to note:
	// We can replace this sequential read then pass to write by using
	// io.Copy(&writer, &reader) which will do it for us
}
