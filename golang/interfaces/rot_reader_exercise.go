package interfaces1

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot13.r.Read(p)
	for i := 0; i < n; i++ {
		if p[i] >= 'A' && p[i] <= 'Z' {
			p[i] = 'A' + (p[i]-'A'+13)%26
		} else if p[i] >= 'a' && p[i] <= 'z' {
			p[i] = 'a' + (p[i]-'a'+13)%26
		}
	}
	return
}

func RotReaderInterface() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
