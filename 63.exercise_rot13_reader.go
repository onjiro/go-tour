package main

import (
	"io"
	"os"
	"strings"
)

var LookUpTable = map[byte]byte{
	'A': 'N',
	'B': 'O',
	'C': 'P',
	'D': 'Q',
	'E': 'R',
	'F': 'S',
	'G': 'T',
	'H': 'U',
	'I': 'V',
	'J': 'W',
	'K': 'X',
	'L': 'Y',
	'M': 'Z',
	'N': 'A',
	'O': 'B',
	'P': 'C',
	'Q': 'D',
	'R': 'E',
	'S': 'F',
	'T': 'G',
	'U': 'H',
	'V': 'I',
	'W': 'J',
	'X': 'K',
	'Y': 'L',
	'Z': 'M',
	'a': 'n',
	'b': 'o',
	'c': 'p',
	'd': 'q',
	'e': 'r',
	'f': 's',
	'g': 't',
	'h': 'u',
	'i': 'v',
	'j': 'w',
	'k': 'x',
	'l': 'y',
	'm': 'z',
	'n': 'a',
	'o': 'b',
	'p': 'c',
	'q': 'd',
	'r': 'e',
	's': 'f',
	'/': '/',
	't': 'g',
	'v': 'i',
	'w': 'j',
	'x': 'k',
	'y': 'l',
	'z': 'm',
	'u': 'h',
}

// rot13Reader is one implements of io.Reader, wraps io.Reader,
// modifying the stream by applying the ROT13 substitution cipher.
type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {
		value, exists := LookUpTable[p[i]]
		if exists {
			p[i] = value
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
