package main

import (
	// "fmt"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
Exercise: rot13Reader
A common pattern is an io.Reader that wraps another io.Reader, modifying the
stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of
compressed data) and returns a *gzip.Reader that also implements io.Reader (a
stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader,
modifying the stream by applying the rot13 substitution cipher to all
alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing
its Read method.
*/

type rot13Reader struct {
	R io.Reader
}

func (r *rot13Reader) Read(buffer []byte) (n int, err error) {
	n, err = r.R.Read(buffer)
	for i := 0; i < n; i++ {
		buffer[i] = rot13(buffer[i])
	}
	return
}

func rot13(char byte) byte {
	switch {
	case char >= 'A' && char <= 'Z': // 65 - 90
		return 'A' + (char-'A'+13)%26
	case char >= 'a' && char <= 'z': // 97 - 122
		return 'a' + (char-'a'+13)%26
	default:
		return char
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r) // You cracked the code!
	fmt.Println()
}
