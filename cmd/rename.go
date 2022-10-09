package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"

	"github.com/cqroot/clibox/internal"
)

func init() {
	rootCmd.AddCommand(renameCmd)
}

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "An efficient batch renaming tool for vimer",
	Long:  `An efficient batch renaming tool for vimer`,
	Run:   RunRenameCmd,
}

type RenamePair struct {
	OldName string
	NewName string
}

func RunRenameCmd(cmd *cobra.Command, args []string) {
	var opsId string = fmt.Sprintf("%d", time.Now().Unix())
	var oldFile string = fmt.Sprintf("/tmp/clibox-rename-old-%s", opsId)
	var newFile string = fmt.Sprintf("/tmp/clibox-rename-new-%s", opsId)

	currentPath, err := os.Getwd()
	internal.ExitIfError(err)

	createTmpFiles(currentPath, oldFile, newFile)
	defer removeTmpFiles(oldFile, newFile)

	runEditor(oldFile, newFile)

	renamePairs := generateRenamePair(oldFile, newFile)

	startRename(renamePairs, currentPath)
}

func createTmpFiles(currentPath string, oldFile string, newFile string) {
	files, err := ioutil.ReadDir(currentPath)
	internal.ExitIfError(err)

	fOld, err := os.OpenFile(oldFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0444)
	internal.ExitIfError(err)
	defer fOld.Close()

	fNew, err := os.Create(newFile)
	internal.ExitIfError(err)
	defer fNew.Close()

	for _, f := range files {
		fOld.WriteString(f.Name())
		fOld.WriteString("\n")

		fNew.WriteString(f.Name())
		fNew.WriteString("\n")
	}
	fOld.Sync()
	fNew.Sync()
}

func removeTmpFiles(oldFile string, newFile string) {
	os.Remove(oldFile)
	os.Remove(newFile)
}

func runEditor(oldFile string, newFile string) {
	var editor string = os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	var args []string = []string{
		// "-c", fmt.Sprintf("command RenameDiff :vertical diffsplit %s", oldFile),
		// "-c", "nmap <C-p> :RenameDiff<CR>",
		// "-c", "RenameDiff",
		"-d", oldFile, newFile,
		"-c", "wincmd l",
		"-c", "foldopen",
		"-c", "autocmd BufEnter * if winnr(\"$\") == 1 | execute \"normal! :q!\\<CR>\" | endif",
	}
	vicmd := exec.Command(editor, args...)
	vicmd.Stdin = os.Stdin
	vicmd.Stdout = os.Stdout
	err := vicmd.Run()
	internal.ExitIfError(err)
}

func generateRenamePair(oldFile string, newFile string) []RenamePair {
	fOld, err := os.Open(oldFile)
	internal.ExitIfError(err)
	oldScanner := bufio.NewScanner(fOld)
	oldScanner.Split(bufio.ScanLines)

	fNew, err := os.Open(newFile)
	internal.ExitIfError(err)
	newScanner := bufio.NewScanner(fNew)
	newScanner.Split(bufio.ScanLines)

	var renamePairs []RenamePair = make([]RenamePair, 0)
	for oldScanner.Scan() && newScanner.Scan() {
		oldName := oldScanner.Text()
		newName := newScanner.Text()
		if oldName == newName {
			continue
		}

		renamePairs = append(renamePairs, RenamePair{
			OldName: oldName,
			NewName: newName,
		})
	}

	return renamePairs
}

func startRename(renamePairs []RenamePair, currentPath string) {
	if len(renamePairs) == 0 {
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Old Name", "New Name"})
	for idx, pair := range renamePairs {
		t.AppendRow(table.Row{idx, pair.OldName, pair.NewName})
	}
	t.Render()

	fmt.Print("Confirm to rename the above file [y/N] ")
	cfmReader := bufio.NewReader(os.Stdin)
	cfmText, err := cfmReader.ReadString('\n')
	internal.ExitIfError(err)

	if cfmText != "y\n" && cfmText != "Y\n" {
		return
	}

	for idx, pair := range renamePairs {
		fmt.Printf(
			"%s %s %s %s\n",
			text.FgGreen.Sprintf("%d:", idx),
			pair.OldName,
			text.FgGreen.Sprint("->"),
			pair.NewName,
		)
		err := os.Rename(
			filepath.Join(currentPath, pair.OldName),
			filepath.Join(currentPath, pair.NewName),
		)
		internal.ExitIfError(err)
	}
}
