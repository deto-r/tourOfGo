package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (reader MyReader) Read(rb []byte) (n int, e error) {
	for n, e = 0, nil; n < len(rb); n++ {
		rb[n] = 'A'
	}
	return
}

func main() {
reader.Validate(MyReader{})
}

