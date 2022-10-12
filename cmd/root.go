package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/cqroot/vinamer/internal"
	"github.com/cqroot/vinamer/renamer"
)

var rootCmd = &cobra.Command{
	Use:   "viname",
	Short: "An efficient batch renaming tool for vimer",
	Long:  `An efficient batch renaming tool for vimer`,
	Run:   RunRootCmd,
}

func Execute() {
	err := rootCmd.Execute()
	internal.ExitIfError(err)
}

func RunRootCmd(cmd *cobra.Command, args []string) {
	var opsId string = fmt.Sprintf("%d", time.Now().Unix())
	var oldFile string = fmt.Sprintf("/tmp/vinamer-old-%s", opsId)
	var newFile string = fmt.Sprintf("/tmp/vinamer-new-%s", opsId)

	currentPath, err := os.Getwd()
	internal.ExitIfError(err)

	renamer.CreateTmpFiles(currentPath, oldFile, newFile)
	defer renamer.RemoveTmpFiles(oldFile, newFile)

	renamer.RunEditor(oldFile, newFile)

	renamePairs := renamer.GenerateRenamePair(oldFile, newFile)

	renamer.StartRename(renamePairs, currentPath)
}

