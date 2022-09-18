package cmd

import (
	"github.com/cqroot/clibox/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clibox",
	Short: "Clibox is a terminal toolbox",
	Long:  `Clibox is a terminal toolbox`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	err := rootCmd.Execute()
	internal.ExitIfError(err)
}
