package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func (i *spaceEraser) Read(b []byte) (n int, err error) {
	n, err = i.r.Read(b)
	var toReturn = []byte{}
	for _, char := range b {
		if char != 32 {
			toReturn = append(toReturn, char)
		}
	}
	copy(b, toReturn)
	return n, err
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}

// Implement a type that satisfies the io.Reader interface and reads from another io.Reader,
// modifying the stream by removing the spaces.
// Expected output: "Helloworld!"
