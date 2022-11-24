package cmd

import (
	"os"
	"path"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cqroot/edname/internal/renamer"
)

var (
	flagAll              bool
	flagConfig           string
	flagDiff             bool
	flagDirectory        bool
	flagDirectoryOnly    bool
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
	rootCmd.Flags().BoolVar(&flagDiff, "diff", false, "diff mode (only works when the editor is vim-like editor)")
	rootCmd.Flags().BoolVarP(&flagDirectory, "directory", "d", false, "include directory")
	rootCmd.Flags().BoolVarP(&flagDirectoryOnly, "directory-only", "D", false, "rename directory only")
	rootCmd.Flags().StringP("editor", "e", "", "")
	rootCmd.Flags().StringVarP(&flagWorkingDirectory, "working-directory", "w", "", "")

	_ = viper.BindPFlag("editor", rootCmd.Flags().Lookup("editor"))
	viper.SetDefault("editor", "vim")
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func InitConfig() {
	if flagConfig != "" {
		// Use config file from the flag.
		viper.SetConfigFile(flagConfig)
	} else {
		configFile, err := xdg.ConfigFile("edname/config.toml")
		cobra.CheckErr(err)
		viper.SetConfigFile(configFile)
	}

	_ = viper.ReadInConfig()
}

func RunRootCmd(cmd *cobra.Command, args []string) {
	InitConfig()

	if !path.IsAbs(flagWorkingDirectory) {
		cwd, err := os.Getwd()
		cobra.CheckErr(err)

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

	if flagDiff {
		r.RunEditorDiff(viper.GetString("editor"))
	} else {
		r.RunEditor(viper.GetString("editor"))
	}

	renamePairs := r.GenerateRenamePairs()

	r.StartRename(renamePairs)
}
