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
	_ = viper.BindPFlag("include-all", rootCmd.Flags().Lookup("all"))
	viper.SetDefault("include-all", false)

	rootCmd.Flags().StringVarP(&flagConfig, "config", "c", "", "config file. default $HOME/.config/edname/config.toml")

	rootCmd.Flags().BoolVarP(&flagDirectory, "directory", "d", false, "include directory")
	_ = viper.BindPFlag("include-directory", rootCmd.Flags().Lookup("directory"))
	viper.SetDefault("include-directory", false)

	rootCmd.Flags().BoolVarP(&flagDirectoryOnly, "directory-only", "D", false, "rename directory only")
	_ = viper.BindPFlag("include-directory-only", rootCmd.Flags().Lookup("directory-only"))
	viper.SetDefault("include-directory-only", false)

	rootCmd.Flags().StringP("editor", "e", "", "")
	_ = viper.BindPFlag("editor", rootCmd.Flags().Lookup("editor"))
	viper.SetDefault("editor", "vim")

	rootCmd.Flags().StringVarP(&flagWorkingDirectory, "working-directory", "w", "", "")

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
		viper.GetString("editor"),
		flagWorkingDirectory,
		viper.GetBool("include-directory"),
		viper.GetBool("include-directory-only"),
		viper.GetBool("include-all"),
	)

	r.Execute()
}
