package stamp

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"testing"
)

func TestNewStamp(t *testing.T) {
	s := NewStamp()
	exp := "unknown"
	if s.Revision != exp {
		t.Errorf("Default revision should be %q", exp)
	}
	if s.ChangelogVersion != exp {
		t.Errorf("Default ChangelogVersion should be %q", exp)
	}
}

func TestParseChangelog(t *testing.T) {
	s := NewStamp()
	var err error
	file := "CHANGELOG.md"
	fsig := fmt.Sprintf("ParseChangelog(%q)", file)
	if err = s.ParseChangelog(file); err != nil {
		t.Errorf("%s should be ok but failed %s", file, err)
	}
	if s.ChangelogVersion == "unknown" {
		t.Errorf("%s should set ChangelogVersion", fsig)
	}

	shouldFail := []struct {
		file string
	}{
		{"nosuchfile.md"},
		{"README.md"}, // Bad format
	}
	for _, d := range shouldFail {
		if err = s.ParseChangelog(d.file); err == nil {
			t.Errorf("ParseChangelog(%q) should fail", d.file)
		}
	}
}

func Test_Revision(t *testing.T) {
	rev, err := Revision(".")
	if err != nil {
		t.Fatal(err)
	}
	if rev == "unknown" {
		t.Errorf("rev should returned %q", rev)
	}
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir("/")
	_, err = Revision(".")
	if err == nil {
		t.Error("Revision() should fail when not in git repository")
	}

}

func Test_GoTemplate(t *testing.T) {
	tpl := NewGoTemplate()
	if tpl == nil {
		t.Error("GoTemplate() should always return a template")
	}
}

func Test_compile_template(t *testing.T) {
	bin := path.Join(os.TempDir(), "stamp")
	out, err := exec.Command("go", "build", "-o", bin,
		"github.com/gregoryv/stamp/cmd/stamp").CombinedOutput()
	if err != nil {
		t.Errorf("%s: %s", out, err)
	}
}
