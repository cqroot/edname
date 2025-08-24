package app

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cqroot/edname/internal/ediff"
	"github.com/cqroot/edname/internal/generator"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/sergi/go-diff/diffmatchpatch"
)

type App struct{}

func Run(editor string, path string, dirOpt bool, dirOnlyOpt bool, allOpt bool) error {
	g := generator.New(path, dirOpt, dirOnlyOpt, allOpt)

	items, err := g.Generate()
	if err != nil {
		return err
	}

	ed := ediff.New(editor)
	ed.AppendItems(items)
	pairs, err := ed.Run()
	if err != nil {
		return err
	}

	if len(pairs) == 0 {
		return nil
	}

	PrintPairs(pairs)

	fmt.Print("Confirm to rename the above file [y/N] ")
	cfmReader := bufio.NewReader(os.Stdin)
	cfmText, err := cfmReader.ReadString('\n')
	if err != nil {
		return err
	}

	cfmText = strings.Map(func(r rune) rune {
		if r == '\n' || r == '\r' || r == ' ' {
			return -1
		}
		return r
	}, cfmText)
	if strings.ToLower(cfmText) != "y" {
		return nil
	}
	fmt.Println()

	for _, pair := range pairs {
		err := os.Rename(
			filepath.Join(path, pair.Prev),
			filepath.Join(path, pair.Curr),
		)
		if err != nil {
			return err
		}
	}

	fmt.Printf("%d items have been renamed.\n", len(pairs))
	return err
}

func GetColoredDiffs(diffs []diffmatchpatch.Diff) (string, string) {
	var sbOld, sbNew strings.Builder

	for _, d := range diffs {
		switch d.Type {
		case diffmatchpatch.DiffEqual:
			sbOld.WriteString(d.Text)
			sbNew.WriteString(d.Text)
		case diffmatchpatch.DiffDelete:
			sbOld.WriteString(color.RedString(d.Text))
		case diffmatchpatch.DiffInsert:
			sbNew.WriteString(color.GreenString(d.Text))
		}
	}

	return sbOld.String(), sbNew.String()
}

func PrintPairs(pairs []ediff.DiffPair) {
	dmp := diffmatchpatch.New()

	t := table.NewWriter()
	t.SetStyle(table.StyleRounded)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Old Name", "New Name"})
	for idx, pair := range pairs {
		diffs := dmp.DiffMain(pair.Prev, pair.Curr, false)
		sOld, sNew := GetColoredDiffs(diffs)
		t.AppendRow(table.Row{idx, sOld, sNew})
	}
	t.Render()
}
