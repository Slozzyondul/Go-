package interfaces1

import "golang.org/x/tour/reader"

type MyReader struct{}

// Read method to satisfy the io.Reader interface
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func ReaderInterfaceInfinityLoop() {
	reader.Validate(MyReader{})
}
