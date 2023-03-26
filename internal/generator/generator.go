package generator

import (
	"os"
)

type GenerateOpt struct {
	WorkDir               string
	ShouldContainDir      bool
	ShouldOnlyContainDir  bool
	ShouldContainDotFiles bool
}

type Generater struct {
	opt GenerateOpt
}

func New(opt GenerateOpt) *Generater {
	return &Generater{
		opt: opt,
	}
}

func (r Generater) Generate() ([]string, error) {
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

		if !r.opt.ShouldContainDotFiles && info.Name()[0] == '.' {
			continue
		}

		if r.opt.ShouldOnlyContainDir {
			if !info.IsDir() {
				continue
			}
		} else if !r.opt.ShouldContainDir {
			if info.IsDir() {
				continue
			}
		}

		result = append(result, info.Name())
	}

	return result, nil
}
