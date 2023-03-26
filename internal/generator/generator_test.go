package generator_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cqroot/edname/internal/generator"
)

func testGenerate(t *testing.T, expect []string, opt generator.GenerateOpt) {
	r := generator.New(opt)

	actual, err := r.Generate()
	require.Nil(t, err)

	require.Equal(t, expect, actual)
}

func TestGenerate(t *testing.T) {
	testGenerate(t, []string{
		"test_file_a",
		"test_file_b",
	}, generator.GenerateOpt{
		WorkDir:               "./testdata",
		ShouldContainDir:      false,
		ShouldOnlyContainDir:  false,
		ShouldContainDotFiles: false,
	})

	testGenerate(t, []string{
		"test_dir_a",
		"test_dir_b",
		"test_file_a",
		"test_file_b",
	}, generator.GenerateOpt{
		WorkDir:               "./testdata",
		ShouldContainDir:      true,
		ShouldOnlyContainDir:  false,
		ShouldContainDotFiles: false,
	})

	testGenerate(t, []string{
		"test_dir_a",
		"test_dir_b",
	}, generator.GenerateOpt{
		WorkDir:               "./testdata",
		ShouldContainDir:      false,
		ShouldOnlyContainDir:  true,
		ShouldContainDotFiles: false,
	})

	testGenerate(t, []string{
		".test_file_a",
		".test_file_b",
		"test_file_a",
		"test_file_b",
	}, generator.GenerateOpt{
		WorkDir:               "./testdata",
		ShouldContainDir:      false,
		ShouldOnlyContainDir:  false,
		ShouldContainDotFiles: true,
	})

	testGenerate(t, []string{
		".test_dir_a",
		".test_dir_b",
		"test_dir_a",
		"test_dir_b",
	}, generator.GenerateOpt{
		WorkDir:               "./testdata",
		ShouldContainDir:      false,
		ShouldOnlyContainDir:  true,
		ShouldContainDotFiles: true,
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
	}, generator.GenerateOpt{
		WorkDir:               "./testdata",
		ShouldContainDir:      true,
		ShouldOnlyContainDir:  false,
		ShouldContainDotFiles: true,
	})

	testGenerate(t, []string{
		"test_dir_a",
		"test_dir_b",
	}, generator.GenerateOpt{
		WorkDir:               "./testdata",
		ShouldContainDir:      true,
		ShouldOnlyContainDir:  true,
		ShouldContainDotFiles: false,
	})

	testGenerate(t, []string{
		".test_dir_a",
		".test_dir_b",
		"test_dir_a",
		"test_dir_b",
	}, generator.GenerateOpt{
		WorkDir:               "./testdata",
		ShouldContainDir:      true,
		ShouldOnlyContainDir:  true,
		ShouldContainDotFiles: true,
	})
}
