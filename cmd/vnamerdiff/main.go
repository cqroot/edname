package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cqroot/vinamer/internal"
	"github.com/cqroot/vinamer/renamer"
)

func main() {
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
