package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"runtime"
)

func getHash(content interface{}) ([]byte, error) {
	buf := bytes.Buffer{}

	be := bendcoder{&buf}

	err := be.encode(content)

	if err != nil {
		return nil, err
	}

	h := sha1.New()

	io.Copy(h, &buf)
	sum := h.Sum(nil)
	fmt.Println(sum)
	return sum, nil
}

func PrintCurrentLine() {
	_, file, line, ok := runtime.Caller(0)
	if ok {
		fmt.Printf("Current file: %s, line: %d\n", file, line)
	} else {
		fmt.Println("Unable to retrieve line information.")
	}
}
