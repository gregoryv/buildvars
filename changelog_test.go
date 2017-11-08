package stamp

import (
	"io/ioutil"
	"testing"
)

func TestChangelog_Version(t *testing.T) {
	content, err := ioutil.ReadFile("CHANGELOG.md")
	if err != nil {
		panic(err)
	}
	data := []struct {
		buf []byte
		exp string
	}{
		{content, "Unreleased"},
		{[]byte(""), ""},
	}
	for _, d := range data {
		changelog := NewChangelog(d.buf[:])
		version, err := changelog.Version()
		if version != d.exp {
			t.Error(err)
		}
	}
}
