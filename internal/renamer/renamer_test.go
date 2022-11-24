package renamer_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/cqroot/edname/internal/renamer"
)

func TestRenamer_GenerateRenameItems(t *testing.T) {
	testRenamerGenerateRenameItems(t, []string{
		"test_file_a",
		"test_file_b",
	}, false, false, false)

	testRenamerGenerateRenameItems(t, []string{
		"test_dir_a",
		"test_dir_b",
		"test_file_a",
		"test_file_b",
	}, true, false, false)

	testRenamerGenerateRenameItems(t, []string{
		"test_dir_a",
		"test_dir_b",
	}, false, true, false)

	testRenamerGenerateRenameItems(t, []string{
		".test_file_a",
		".test_file_b",
		"test_file_a",
		"test_file_b",
	}, false, false, true)

	testRenamerGenerateRenameItems(t, []string{
		".test_dir_a",
		".test_dir_b",
		"test_dir_a",
		"test_dir_b",
	}, false, true, true)

	testRenamerGenerateRenameItems(t, []string{
		".test_dir_a",
		".test_dir_b",
		".test_file_a",
		".test_file_b",
		"test_dir_a",
		"test_dir_b",
		"test_file_a",
		"test_file_b",
	}, true, false, true)

	testRenamerGenerateRenameItems(t, []string{
		"test_dir_a",
		"test_dir_b",
	}, true, true, false)

	testRenamerGenerateRenameItems(t, []string{
		".test_dir_a",
		".test_dir_b",
		"test_dir_a",
		"test_dir_b",
	}, true, true, true)
}

func testRenamerGenerateRenameItems(t *testing.T, expect []string, dirOpt bool, dirOnlyOpt bool, allOpt bool) {
	r := renamer.New("./testdata", dirOpt, dirOnlyOpt, allOpt)

	chItems := make(chan string)
	go r.GenerateRenameItems(chItems)

	var actual []string = make([]string, 0)
	for file := range chItems {
		actual = append(actual, file)
	}
	assert.Equal(t, expect, actual)
}
