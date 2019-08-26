package main

import (
	"fmt"
	"strings"
)


type MyReader struct{
	source string
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read([]byte) (int, error) {
	r:= strings.NewReader(mr.source)
	byte := make([]byte, 8)
	n, err := r.Read(byte)
	fmt.Print(byte)
	return n, err
}

func main() {
	a := MyReader{"abcsakjhsdkajshdkjhsad"}
	var b []byte
	n, err := a.Read(b)
	fmt.Println(n, "---", err)
}