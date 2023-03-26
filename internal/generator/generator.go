package generator

import (
	"os"
)

type Generater struct {
	workDir               string
	shouldContainDir      bool
	shouldOnlyContainDir  bool
	shouldContainDotFiles bool
}

func New(workDir string, shouldContainDir, shouldOnlyContainDir, shouldContainDotFiles bool) *Generater {
	return &Generater{
		workDir:               workDir,
		shouldContainDir:      shouldContainDir,
		shouldOnlyContainDir:  shouldOnlyContainDir,
		shouldContainDotFiles: shouldContainDotFiles,
	}
}

func (r Generater) Generate() ([]string, error) {
	entries, err := os.ReadDir(r.workDir)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		if !r.shouldContainDotFiles && info.Name()[0] == '.' {
			continue
		}

		if r.shouldOnlyContainDir {
			if !info.IsDir() {
				continue
			}
		} else if !r.shouldContainDir {
			if info.IsDir() {
				continue
			}
		}

		result = append(result, info.Name())
	}

	return result, nil
}
