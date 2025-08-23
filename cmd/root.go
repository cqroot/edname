package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cqroot/edname/internal/app"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	flagAll           bool
	flagDirectory     bool
	flagDirectoryOnly bool
	flagEditor        string

	rootCmd = &cobra.Command{
		Use:   "edname",
		Args:  cobra.MatchAll(cobra.RangeArgs(0, 1), cobra.OnlyValidArgs),
		Short: "Use your favorite editor to batch rename files and directories.",
		Long: `Use your favorite editor to batch rename files and directories.

Notice:
1. Do not add or subtract lines.
2. Unchanged lines are ignored.`,
		Run: RunRootCmd,
	}
)

func init() {
	rootCmd.Flags().BoolVarP(&flagAll, "all", "a", false, "do not ignore entries starting with .")
	rootCmd.Flags().BoolVarP(&flagDirectory, "directory", "d", false, "include directory")
	rootCmd.Flags().
		BoolVarP(&flagDirectoryOnly, "directory-only", "D", false, "rename directory only")
	rootCmd.Flags().StringVarP(&flagEditor, "editor", "e", "", "")
}

func CheckErr(err error) {
	if err != nil {
		fmt.Fprintln(
			os.Stderr,
			color.New(color.BgHiRed, color.FgHiBlack).Sprintf("Error:"),
			err,
		)
		os.Exit(1)
	}
}

func Execute() {
	err := rootCmd.Execute()
	CheckErr(err)
}

func RunRootCmd(cmd *cobra.Command, args []string) {
	var err error

	wd := ""
	if len(args) == 1 {
		wd = args[0]
	}

	if !filepath.IsAbs(wd) {
		wd, err = filepath.Abs(wd)
		CheckErr(err)
	}

	if flagEditor == "" {
		envEditor := os.Getenv("EDITOR")
		if envEditor != "" {
			flagEditor = envEditor
		} else {
			flagEditor = "vim"
		}
	}

	err = app.Run(
		flagEditor,
		wd,
		flagDirectory,
		flagDirectoryOnly,
		flagAll,
	)
	CheckErr(err)
}
