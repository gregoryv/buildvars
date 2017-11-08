package stamp

import (
	"testing"
	"os"
)

func Test_NewStamp(t *testing.T) {
	build, err := NewStamp()
	if err != nil {
		t.Fatal(err)
	}
	if build == nil {
		t.Errorf("NewBuild() should return a build")
	}
	os.Chdir("/")
	_, err = NewStamp()
	if err == nil {
		t.Error("NewStamp() should fail when not in git repository")
	}
}

func Test_GoTemplate(t *testing.T) {
	tpl := GoTemplate()
	if tpl == nil {
		t.Error("GoTemplate() should always return a template")
	}
}
