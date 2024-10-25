package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
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
