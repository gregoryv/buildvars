package main

import (
	"io/ioutil"
	"os/exec"
	"testing"
	"bytes"
)

func Test_Generate(t *testing.T) {
	var buf []byte
	out := bytes.NewBuffer(buf)
	err := Generate(out)
	if err != nil {
		t.Errorf("%s: %s", out.String(), err)
	}
	ioutil.WriteFile("x.go", out.Bytes(), 0755)

	// Check that x.go is compilable
	buf, err = exec.Command("go", "build", ".").Output()
	if err != nil {
		t.Errorf("%s: %s", buf, err)
	}

}
