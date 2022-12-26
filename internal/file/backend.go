package file

import (
	"os"
	"path/filepath"
)

type Backend struct {
	Path       string
	DirOpt     bool
	DirOnlyOpt bool
	AllOpt     bool
}

func (b Backend) Generate() ([]string, error) {
	entries, err := os.ReadDir(b.Path)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		if !b.AllOpt && info.Name()[0] == '.' {
			continue
		}

		if b.DirOnlyOpt {
			if !info.IsDir() {
				continue
			}
		} else if !b.DirOpt {
			if info.IsDir() {
				continue
			}
		}

		result = append(result, info.Name())
	}

	return result, nil
}

func (b Backend) Rename(namePairs [][]string, finishFunc func(string, string)) error {
	for _, names := range namePairs {
		err := os.Rename(
			filepath.Join(b.Path, names[0]),
			filepath.Join(b.Path, names[1]),
		)
		if err != nil {
			return err
		}
		finishFunc(names[0], names[1])
	}
	return nil
}
