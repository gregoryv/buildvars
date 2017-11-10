package stamp

import (
    "testing"
)

func TestNewGit(t *testing.T) {
	data := []struct {
		wd string
		exp bool
	}{
		{"", true},
	}
	for _, d := range data {
		g := NewGit(d.wd)
		if d.exp && g == nil {
			t.Error("NewGit(%q) returned nil")
		}
	}
}
