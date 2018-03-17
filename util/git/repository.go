package git

import (
	"os"
	"os/exec"
)

// Clone a repository
func Clone(url, branch, dest string) error {
	cmd := exec.Command(git, "clone", "-b", branch, url, dest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
