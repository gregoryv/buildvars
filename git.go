package stamp

type Revisioner interface {
	Revision() (string, error)
}

type Git struct {
	wd string
}

func NewGit(wd string) *Git {
	return &Git{
		wd: wd,
	}
}
