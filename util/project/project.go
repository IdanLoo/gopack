package project

import (
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/IdanLoo/gopack/util/config"
)

// Project is a map to repository
type Project struct {
	Proj   string
	Src    string
	Bin    string
	Config string
}

type execType int

const (
	_ execType = iota
	execBuild
	execRun
)

func newProject(name string) *Project {
	proj := workspace + "/" + name

	return &Project{
		proj,
		proj + "/src",
		proj + "/bin",
		proj + "/src/gopack.yaml",
	}
}

// Of name return a Project Object.
// If project not exist, then create.
func Of(name string) (*Project, error) {
	var (
		proj = newProject(name)
		err  error
	)

	if !config.IsExist(proj.Proj) {
		err = proj.createDir()
	}

	if err != nil {
		return nil, err
	}

	return proj, nil
}

func (p *Project) createDir() error {
	for _, path := range []string{p.Proj, p.Src, p.Bin} {
		if err := createDir(path); err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) cleanSrc() error {
	return os.RemoveAll(p.Src)
}

// Build to run build commands
func (p *Project) Build() error {
	return p.exec(execBuild)
}

// Run to run run commands
func (p *Project) Run() error {
	return p.exec(execRun)
}

func (p *Project) exec(t execType) error {
	operation, err := newOperation(p.Config)

	if err != nil {
		return err
	}

	var cmdStr string

	switch t {
	case execBuild:
		cmdStr = strings.Join(operation.Build, " && ")
	case execRun:
		cmdStr = strings.Join(operation.Run, " && ")
	default:
		return errors.New("can not exec this function")
	}

	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	return cmd.Run()
}
