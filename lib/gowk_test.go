package lib

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	code := `//
package main

import "fmt"

func main() {
fmt.Println("Hello")
}
`
	content := []byte(code)
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Error(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		t.Error(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Error(err)
	}

	err := goRun([]string{tmpfile.Name()})
	if err != nil {
		t.Error(err)
		return
	}
}
