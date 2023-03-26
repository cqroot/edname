package cmd

import (
	"os"
	"path/filepath"

	"github.com/cqroot/edname/internal/app"
	"github.com/spf13/cobra"
)

var (
	flagAll              bool
	flagConfig           string
	flagDirectory        bool
	flagDirectoryOnly    bool
	flagEditor           string
	flagWorkingDirectory string

	rootCmd = &cobra.Command{
		Use:   "edname",
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
	rootCmd.Flags().StringVarP(&flagConfig, "config", "c", "", "config file. default $HOME/.config/edname/config.toml")
	rootCmd.Flags().BoolVarP(&flagDirectory, "directory", "d", false, "include directory")
	rootCmd.Flags().BoolVarP(&flagDirectoryOnly, "directory-only", "D", false, "rename directory only")
	rootCmd.Flags().StringVarP(&flagEditor, "editor", "e", "", "")
	rootCmd.Flags().StringVarP(&flagWorkingDirectory, "working-directory", "w", "", "")
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func RunRootCmd(cmd *cobra.Command, args []string) {
	var err error

	if !filepath.IsAbs(flagWorkingDirectory) {
		flagWorkingDirectory, err = filepath.Abs(flagWorkingDirectory)
		cobra.CheckErr(err)
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
		flagWorkingDirectory,
		flagDirectory,
		flagDirectoryOnly,
		flagAll,
	)
	cobra.CheckErr(err)
}
