package stamp

import (
	"io/ioutil"
	"testing"
)

func TestChangelog_Version(t *testing.T) {
	content, err := ioutil.ReadFile("changelog.md")
	if err != nil {
		panic(err)
	}
	changelog := NewChangelog(content)
	version, err := changelog.Version()
	if err != nil {
		t.Error(err)
	}
	if version == "" {
		t.Fail()
	}
}
