package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/cqroot/vina/internal/errutil"
	"github.com/cqroot/vina/internal/renamer"
)

func main() {
	app := &cli.App{
		Name:  "vina",
		Usage: "An efficient batch renaming tool for vimer.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "diff",
				Aliases: []string{"d"},
				Value:   false,
				Usage:   "diff mode (only works when the editor is vim or neovim)",
			},
			&cli.BoolFlag{
				Name:    "directory",
				Aliases: []string{"D"},
				Value:   false,
				Usage:   "include directory",
			},
		},
		Action: runCmd,
	}

	errutil.ExitIfError(app.Run(os.Args))
}

func runCmd(cCtx *cli.Context) error {
	var opsId string = fmt.Sprintf("%d", time.Now().Unix())
	var oldFile string = fmt.Sprintf("/tmp/vina-old-%s", opsId)
	var newFile string = fmt.Sprintf("/tmp/vina-new-%s", opsId)

	currentPath, err := os.Getwd()
	errutil.ExitIfError(err)

	renamer.PrintHelpMessage()

	renamer.CreateTmpFiles(currentPath, oldFile, newFile, cCtx.Bool("directory"))
	defer renamer.RemoveTmpFiles(oldFile, newFile)

	if cCtx.Bool("diff") {
		renamer.RunEditorDiff(oldFile, newFile)
	} else {
		renamer.RunEditor(oldFile, newFile)
	}

	renamePairs := renamer.GenerateRenamePair(oldFile, newFile)

	renamer.StartRename(renamePairs, currentPath)

	return nil
}
