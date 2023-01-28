package app

import (
	"fmt"

	"github.com/cqroot/ediff"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/cqroot/edname/internal/file"
)

type App struct{}

func Run(editor string, path string, dirOpt bool, dirOnlyOpt bool, allOpt bool) error {
	backend := file.New(path, dirOpt, dirOnlyOpt, allOpt)
	items, err := backend.Generate()
	if err != nil {
		return err
	}

	ed := ediff.New(editor)
	ed.AppendItems(items)
	pairs, err := ed.Run()
	if err != nil {
		return err
	}

	for _, pair := range pairs {
		err := backend.Rename(pair.Prev, pair.Curr)
		if err != nil {
			return err
		}

		fmt.Printf(
			"%s %s %s\n",
			pair.Prev,
			text.FgGreen.Sprint("->"),
			pair.Curr,
		)
	}

	return err
}
