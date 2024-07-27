package interfaces1

//The Go standard library contains many implementations of this interface,
//including files, network connections, compressors, ciphers, and others.
//It returns an io.EOF error when the stream ends.

import (
	"fmt"
	"io"
	"strings"
)

func ReaderInterface() {
	r := strings.NewReader("wow, Reader!")

	b := make([]byte, 12)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
