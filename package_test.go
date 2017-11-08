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
