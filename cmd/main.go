package main

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/cqroot/vina/internal/errutil"
	"github.com/cqroot/vina/internal/renamer"
)

func main() {
	app := &cli.App{
		Name: "vina",
		Usage: `Use your favorite editor to batch rename files and directories.

Originally designed for vim, but not just vim.

Notice:
1. Do not add or subtract lines.
2. Unchanged lines are ignored.`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Value:   false,
				Usage:   "do not ignore entries starting with .",
			},
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
			&cli.BoolFlag{
				Name:  "directory-only",
				Value: false,
				Usage: "rename directory only",
			},
			&cli.StringFlag{
				Name:    "editor",
				Aliases: []string{"e"},
				Value:   "$EDITOR",
			},
		},
		Action: runCmd,
	}

	errutil.ExitIfError(app.Run(os.Args))
}

func runCmd(cCtx *cli.Context) error {
	currentPath, err := os.Getwd()
	errutil.ExitIfError(err)

	r := renamer.New(
		currentPath,
		cCtx.Bool("directory"),
		cCtx.Bool("directory-only"),
		cCtx.Bool("all"),
	)

	r.CreateTmpFiles()
	defer r.RemoveTmpFiles()

	var editor string = cCtx.String("editor")
	if editor == "$EDITOR" {
		editor = os.Getenv("EDITOR")
	}

	if cCtx.Bool("diff") {
		r.RunEditorDiff(editor)
	} else {
		r.RunEditor(editor)
	}

	renamePairs := r.GenerateRenamePairs()

	r.StartRename(renamePairs)

	return nil
}
