package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		c := b[i]
		switch {
			case c >= 'a' && c <= 'z':
				b[i] = 'a' + ((c - 'a' + 13) % 26)
			case c >= 'A' && c <= 'Z':
				b[i] = 'A' + ((c - 'A' + 13) % 26)
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
