package cmd

import (
	"os"
	"path"

	"github.com/cqroot/goutil/errutil"
	"github.com/spf13/cobra"

	"github.com/cqroot/vina/internal/renamer"
)

var (
	flagAll              bool
	flagDiff             bool
	flagDirectory        bool
	flagDirectoryOnly    bool
	flagEditor           string
	flagWorkingDirectory string

	rootCmd = &cobra.Command{
		Use:   "vina",
		Short: "Use your favorite editor to batch rename files and directories.",
		Long: `Use your favorite editor to batch rename files and directories.

Originally designed for vim, but not just vim.

Notice:
1. Do not add or subtract lines.
2. Unchanged lines are ignored.`,
		Run: RunRootCmd,
	}
)

func init() {
	rootCmd.Flags().BoolVarP(&flagAll, "all", "a", false, "do not ignore entries starting with .")
	rootCmd.Flags().BoolVar(&flagDiff, "diff", false, "diff mode (only works when the editor is vim-like editor)")
	rootCmd.Flags().BoolVarP(&flagDirectory, "directory", "d", false, "include directory")
	rootCmd.Flags().BoolVarP(&flagDirectoryOnly, "directory-only", "D", false, "rename directory only")
	rootCmd.Flags().StringVarP(&flagEditor, "editor", "e", "$EDITOR", "")
	rootCmd.Flags().StringVarP(&flagWorkingDirectory, "working-directory", "w", "", "")
}

func Execute() {
	err := rootCmd.Execute()
	errutil.ExitIfError(err)
}

func RunRootCmd(cmd *cobra.Command, args []string) {
	if !path.IsAbs(flagWorkingDirectory) {
		cwd, err := os.Getwd()
		errutil.ExitIfError(err)

		if flagWorkingDirectory == "" {
			flagWorkingDirectory = cwd
		} else {
			flagWorkingDirectory = path.Join(cwd, flagWorkingDirectory)
		}
	}

	r := renamer.New(
		flagWorkingDirectory,
		flagDirectory,
		flagDirectoryOnly,
		flagAll,
	)

	r.CreateTmpFiles()
	defer r.RemoveTmpFiles()

	if flagEditor == "$EDITOR" {
		flagEditor = os.Getenv("EDITOR")
	}

	if flagDiff {
		r.RunEditorDiff(flagEditor)
	} else {
		r.RunEditor(flagEditor)
	}

	renamePairs := r.GenerateRenamePairs()

	r.StartRename(renamePairs)
}
