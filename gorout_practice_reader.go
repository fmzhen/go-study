package main

import "golang.org/x/tour/reader"

//error
type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (int, error) {
	//b = append(b, 'A')   it is nessary to add "b = "
	for i := 1; i < len(b); i++ {
		b[i-1] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
