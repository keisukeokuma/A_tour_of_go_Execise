package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
// (構造体)　read (引数)(返り値)
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
