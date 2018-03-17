package project

import (
	"os"

	"github.com/IdanLoo/gopack/util/config"
)

// Operation is some shell command in gopack.yaml
type Operation struct {
	Build []string
	Run   []string
}

// NewOperation is constructor of Operation
func newOperation(path string) (*Operation, error) {
	if !config.IsExist(path) {
		return nil, os.ErrNotExist
	}

	operation := &Operation{}

	if err := config.Parse(path, operation); err != nil {
		return nil, err
	}

	return operation, nil
}
