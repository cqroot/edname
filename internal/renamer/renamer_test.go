package renamer_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cqroot/edname/internal/renamer"
)

func testGenerate(t *testing.T, expect []string, opt renamer.RenameOpt) {
	r, err := renamer.New(opt)
	require.Nil(t, err)

	actual, err := r.Generate()
	require.Nil(t, err)

	require.Equal(t, expect, actual)
}

func TestGenerate(t *testing.T) {
	testGenerate(t, []string{
		"test_file_a",
		"test_file_b",
	}, renamer.RenameOpt{
		WorkDir:              "./testdata",
		ShouldRenameDir:      false,
		ShouldOnlyRenameDir:  false,
		ShouldRenameDotFiles: false,
	})

	testGenerate(t, []string{
		"test_dir_a",
		"test_dir_b",
		"test_file_a",
		"test_file_b",
	}, renamer.RenameOpt{
		WorkDir:              "./testdata",
		ShouldRenameDir:      true,
		ShouldOnlyRenameDir:  false,
		ShouldRenameDotFiles: false,
	})

	testGenerate(t, []string{
		"test_dir_a",
		"test_dir_b",
	}, renamer.RenameOpt{
		WorkDir:              "./testdata",
		ShouldRenameDir:      false,
		ShouldOnlyRenameDir:  true,
		ShouldRenameDotFiles: false,
	})

	testGenerate(t, []string{
		".test_file_a",
		".test_file_b",
		"test_file_a",
		"test_file_b",
	}, renamer.RenameOpt{
		WorkDir:              "./testdata",
		ShouldRenameDir:      false,
		ShouldOnlyRenameDir:  false,
		ShouldRenameDotFiles: true,
	})

	testGenerate(t, []string{
		".test_dir_a",
		".test_dir_b",
		"test_dir_a",
		"test_dir_b",
	}, renamer.RenameOpt{
		WorkDir:              "./testdata",
		ShouldRenameDir:      false,
		ShouldOnlyRenameDir:  true,
		ShouldRenameDotFiles: true,
	})

	testGenerate(t, []string{
		".test_dir_a",
		".test_dir_b",
		".test_file_a",
		".test_file_b",
		"test_dir_a",
		"test_dir_b",
		"test_file_a",
		"test_file_b",
	}, renamer.RenameOpt{
		WorkDir:              "./testdata",
		ShouldRenameDir:      true,
		ShouldOnlyRenameDir:  false,
		ShouldRenameDotFiles: true,
	})

	testGenerate(t, []string{
		"test_dir_a",
		"test_dir_b",
	}, renamer.RenameOpt{
		WorkDir:              "./testdata",
		ShouldRenameDir:      true,
		ShouldOnlyRenameDir:  true,
		ShouldRenameDotFiles: false,
	})

	testGenerate(t, []string{
		".test_dir_a",
		".test_dir_b",
		"test_dir_a",
		"test_dir_b",
	}, renamer.RenameOpt{
		WorkDir:              "./testdata",
		ShouldRenameDir:      true,
		ShouldOnlyRenameDir:  true,
		ShouldRenameDotFiles: true,
	})
}
