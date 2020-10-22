package stamp

import (
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/gregoryv/asserter"
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
	ok := asserter.Wrap(t).Ok
	ok(s.ParseChangelog("changelog.md"))

	bad := asserter.Wrap(t).Bad
	bad(s.ParseChangelog("nosuchfile.md"))
	bad(s.ParseChangelog("README.md"))
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
