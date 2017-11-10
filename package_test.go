package stamp

import (
	"os"
	"testing"
)

func Test_NewStamp(t *testing.T) {
	build, err := NewStamp()
	if err != nil {
		t.Fatal(err)
	}
	if build == nil {
		t.Errorf("NewBuild() should return a build")
	}
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir("/")
	_, err = NewStamp()
	if err == nil {
		t.Error("NewStamp() should fail when not in git repository")
	}

}

func xTest_GoTemplate(t *testing.T) {
	tpl := GoTemplate()
	if tpl == nil {
		t.Error("GoTemplate() should always return a template")
	}
}
