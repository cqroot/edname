package main

import (
	"bufio"
	"fmt"
	"os"

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

func PrintHelpMessage() {
	fmt.Println(`[ ViNa ]
Modify the buffer on the right to rename.
 
Notice:
1. Do not add or subtract lines.
2. Do not modify the buffer on the left.
3. Unchanged lines are ignored.`)

	cfmReader := bufio.NewReader(os.Stdin)
	_, err := cfmReader.ReadByte()
	errutil.ExitIfError(err)

	fmt.Println("Opening editor...")
}

func runCmd(cCtx *cli.Context) error {
	currentPath, err := os.Getwd()
	errutil.ExitIfError(err)

	PrintHelpMessage()

	r := renamer.New(currentPath, cCtx.Bool("directory"), cCtx.Bool("all"))

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
