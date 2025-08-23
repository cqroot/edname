package generator_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cqroot/edname/internal/generator"
)

func testGenerate(
	t *testing.T,
	expect []string,
	shouldContainDir, shouldOnlyContainDir, shouldContainDotFiles bool,
) {
	r := generator.New("./testdata", shouldContainDir, shouldOnlyContainDir, shouldContainDotFiles)

	actual, err := r.Generate()
	require.Nil(t, err)

	require.Equal(t, expect, actual)
}

func TestGenerate(t *testing.T) {
	testGenerate(t, []string{
		"test_file_a",
		"test_file_b",
	},
		false, false, false,
	)

	testGenerate(t, []string{
		"test_dir_a",
		"test_dir_b",
		"test_file_a",
		"test_file_b",
	},
		true, false, false,
	)

	testGenerate(t, []string{
		"test_dir_a",
		"test_dir_b",
	},
		false, true, false,
	)

	testGenerate(t, []string{
		".test_file_a",
		".test_file_b",
		"test_file_a",
		"test_file_b",
	},
		false, false, true,
	)

	testGenerate(t, []string{
		".test_dir_a",
		".test_dir_b",
		"test_dir_a",
		"test_dir_b",
	},
		false, true, true,
	)

	testGenerate(t, []string{
		".test_dir_a",
		".test_dir_b",
		".test_file_a",
		".test_file_b",
		"test_dir_a",
		"test_dir_b",
		"test_file_a",
		"test_file_b",
	},
		true, false, true,
	)

	testGenerate(t, []string{
		"test_dir_a",
		"test_dir_b",
	},
		true, true, false,
	)

	testGenerate(t, []string{
		".test_dir_a",
		".test_dir_b",
		"test_dir_a",
		"test_dir_b",
	},
		true, true, true,
	)
}
