package main

import (
	"os"
	"path"

	"github.com/cqroot/goutil/errutil"
	"github.com/urfave/cli/v2"

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
				Name:  "diff",
				Value: false,
				Usage: "diff mode (only works when the editor is vim-like editor)",
			},
			&cli.BoolFlag{
				Name:    "directory",
				Aliases: []string{"d"},
				Value:   false,
				Usage:   "include directory",
			},
			&cli.BoolFlag{
				Name:    "directory-only",
				Aliases: []string{"D"},
				Value:   false,
				Usage:   "rename directory only",
			},
			&cli.StringFlag{
				Name:    "editor",
				Aliases: []string{"e"},
				Value:   "$EDITOR",
			},
			&cli.PathFlag{
				Name:    "working-directory",
				Aliases: []string{"w"},
				Value:   "",
			},
		},
		Action: runCmd,
	}

	errutil.ExitIfError(app.Run(os.Args))
}

func runCmd(cCtx *cli.Context) error {
	var workPath string = cCtx.Path("working-directory")
	if !path.IsAbs(workPath) {
		cwd, err := os.Getwd()
		errutil.ExitIfError(err)

		if workPath == "" {
			workPath = cwd
		} else {
			workPath = path.Join(cwd, workPath)
		}
	}

	r := renamer.New(
		workPath,
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
