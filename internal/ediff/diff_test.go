package ediff_test

import (
	"testing"

	"github.com/cqroot/edname/internal/ediff"
	"github.com/stretchr/testify/require"
)

func TestDiff(t *testing.T) {
	ed := ediff.New("sed")
	ed.SetEditorArgs([]string{
		"-i", "s/item/obj/g",
	})

	ed.AppendItems([]string{
		"item 1",
		"item 2",
		"item 3",
	})

	pairs, err := ed.Run()
	require.Nil(t, err)

	require.Equal(t, []ediff.DiffPair{
		{Prev: "item 1", Curr: "obj 1"},
		{Prev: "item 2", Curr: "obj 2"},
		{Prev: "item 3", Curr: "obj 3"},
	}, pairs)
}
