package stamp

import (
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
}

func Test_GoTemplate(t *testing.T) {
	tpl := GoTemplate()
	if tpl == nil {
		t.Error("GoTemplate() should always return a template")
	}
}
