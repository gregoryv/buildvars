package main

import (
	"testing"
	"os/exec"
	"io/ioutil"
)

func Test_x(t *testing.T) {
	out, err := exec.Command("./go-buildvars").Output()
	if err != nil {
		t.Errorf("%s: %s", out, err)
	}
	ioutil.WriteFile("x.go", out, 0755)

	// Check that x.go is compilable
	out, err = exec.Command("go", "build", ".").Output()
	if err != nil {
		t.Errorf("%s: %s", out, err)
	}
	
}
