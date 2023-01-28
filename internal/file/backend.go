package file

import (
	"os"
	"path/filepath"
)

type Backend struct {
	path       string
	dirOpt     bool
	dirOnlyOpt bool
	allOpt     bool
}

func New(path string, dirOpt bool, dirOnlyOpt bool, allOpt bool) *Backend {
	return &Backend{
		path:       path,
		dirOpt:     dirOpt,
		dirOnlyOpt: dirOnlyOpt,
		allOpt:     allOpt,
	}
}

func (b Backend) Generate() ([]string, error) {
	entries, err := os.ReadDir(b.path)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		if !b.allOpt && info.Name()[0] == '.' {
			continue
		}

		if b.dirOnlyOpt {
			if !info.IsDir() {
				continue
			}
		} else if !b.dirOpt {
			if info.IsDir() {
				continue
			}
		}

		result = append(result, info.Name())
	}

	return result, nil
}

func (b Backend) Rename(prev string, curr string) error {
	err := os.Rename(
		filepath.Join(b.path, prev),
		filepath.Join(b.path, curr),
	)
	if err != nil {
		return err
	}

	return nil
}
