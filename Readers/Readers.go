package main

//import "fmt"
import "golang.org/x/tour/reader"

// MyReader is a struct
type MyReader struct{}

// TODO: Add a Read([]byte) (n int, err error) method to MyReader.
func (reader MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		//fmt.Println(b[i])
		b[i] = 65
		//fmt.Println(b[i])
	}
	n = len(b)
	err = nil
	return 
}

func main() {
	reader.Validate(MyReader{})
}
