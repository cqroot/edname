package executor

import (
	"os"
	"path/filepath"
)

type Executor struct {
	workDir string
}

func New(workDir string) *Executor {
	return &Executor{
		workDir: workDir,
	}
}

func (e Executor) Rename(prev string, curr string) error {
	err := os.Rename(
		filepath.Join(e.workDir, prev),
		filepath.Join(e.workDir, curr),
	)
	if err != nil {
		return err
	}

	return nil
}
