package main

import (
	"io"
	"os"
	"strings"
)

type Rot13 struct {
	r io.Reader
}

// dummy stream what comes in goes out
func (rot13 Rot13) _Read(b []byte) (int, error) {
	l, e := rot13.r.Read(b)
	return l, e
}

func (rot13 Rot13) Read(b []byte) (int, error) {
	l, e := rot13.r.Read(b)
	if e == nil {
		for i, j := 0, byte(0); i < len(b); i += 1 {
			j = b[i]
			if j == 0 {
				break
			}
			if j >= 65 && j <= 90 {
				if (j + 13) > 90 {
					j = 64 + j + 13 - 90
				} else {
					j = j + 13
				}
				b[i] = j
			} else if j >= 97 && j <= 122 {
				if (j + 13) > 122 {
					j = 96 + j + 13 - 122
				} else {
					j = j + 13
				}
				b[i] = j
			}
		}
	}
	return l, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	stringReader := Rot13{
		r: s,
	}

	io.Copy(os.Stdout, &stringReader)
}
