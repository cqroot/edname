package file_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cqroot/edname/internal/file"
)

func TestGenerate(t *testing.T) {
	tfunc := func(t *testing.T, expect []string, dirOpt bool, dirOnlyOpt bool, allOpt bool) {
		b := file.New(
			"./testdata",
			dirOpt,
			dirOnlyOpt,
			allOpt,
		)
		actual, err := b.Generate()
		if err != nil {
			t.Fatal(err)
		}
		require.Equal(t, expect, actual)
	}

	tfunc(t, []string{
		"test_file_a",
		"test_file_b",
	}, false, false, false)

	tfunc(t, []string{
		"test_dir_a",
		"test_dir_b",
		"test_file_a",
		"test_file_b",
	}, true, false, false)

	tfunc(t, []string{
		"test_dir_a",
		"test_dir_b",
	}, false, true, false)

	tfunc(t, []string{
		".test_file_a",
		".test_file_b",
		"test_file_a",
		"test_file_b",
	}, false, false, true)

	tfunc(t, []string{
		".test_dir_a",
		".test_dir_b",
		"test_dir_a",
		"test_dir_b",
	}, false, true, true)

	tfunc(t, []string{
		".test_dir_a",
		".test_dir_b",
		".test_file_a",
		".test_file_b",
		"test_dir_a",
		"test_dir_b",
		"test_file_a",
		"test_file_b",
	}, true, false, true)

	tfunc(t, []string{
		"test_dir_a",
		"test_dir_b",
	}, true, true, false)

	tfunc(t, []string{
		".test_dir_a",
		".test_dir_b",
		"test_dir_a",
		"test_dir_b",
	}, true, true, true)
}
