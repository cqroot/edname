package renamer

import (
	"os"
	"path/filepath"
)

type RenameOpt struct {
	WorkDir              string
	ShouldRenameDir      bool
	ShouldOnlyRenameDir  bool
	ShouldRenameDotFiles bool
}

type Renamer struct {
	opt RenameOpt
}

func New(opt RenameOpt) (*Renamer, error) {
	if opt.WorkDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		opt.WorkDir = cwd
	}
	if !filepath.IsAbs(opt.WorkDir) {
		var err error
		opt.WorkDir, err = filepath.Abs(opt.WorkDir)
		if err != nil {
			return nil, err
		}
	}

	return &Renamer{
		opt: opt,
	}, nil
}

func (r Renamer) Generate() ([]string, error) {
	entries, err := os.ReadDir(r.opt.WorkDir)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		if !r.opt.ShouldRenameDotFiles && info.Name()[0] == '.' {
			continue
		}

		if r.opt.ShouldOnlyRenameDir {
			if !info.IsDir() {
				continue
			}
		} else if !r.opt.ShouldRenameDir {
			if info.IsDir() {
				continue
			}
		}

		result = append(result, info.Name())
	}

	return result, nil
}

func (r Renamer) Rename(prev string, curr string) error {
	err := os.Rename(
		filepath.Join(r.opt.WorkDir, prev),
		filepath.Join(r.opt.WorkDir, curr),
	)
	if err != nil {
		return err
	}

	return nil
}
