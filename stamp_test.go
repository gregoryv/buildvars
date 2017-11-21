package stamp

import (
	"os"
	"os/exec"
	"path"
	"testing"
)

func Test_Parse(t *testing.T) {
	build, err := Parse(".")
	if err != nil {
		t.Fatal(err)
	}
	if build == nil {
		t.Errorf("NewBuild() should return a build")
	}
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir("/")
	_, err = Parse(".")
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

func Test_compile_template(t *testing.T) {
	out, err := exec.Command("go", "build", "-o", path.Join(os.TempDir(), "stamp"),
		"github.com/gregoryv/stamp/cmd/stamp").CombinedOutput()
	if err != nil {
		t.Errorf("%s: %s", out, err)
	}
}
