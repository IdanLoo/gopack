package project

import (
	"github.com/IdanLoo/gopack/util/git"
)

// Clone a project
func Clone(name, branch, url string) error {
	proj, err := Of(name)

	if err != nil {
		return err
	}

	if err = proj.cleanSrc(); err != nil {
		return err
	}

	return git.Clone(url, branch, proj.Src)
}
