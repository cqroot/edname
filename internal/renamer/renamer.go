package renamer

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/cqroot/vina/internal/errutil"
)

type RenamePair struct {
	OldName string
	NewName string
}

func PrintHelpMessage() {
	fmt.Println("[ ViNa ]")
	fmt.Println("Modify the buffer on the right to rename.")
	fmt.Println(" ")
	fmt.Println("Notice:")
	fmt.Println("1. Do not add or subtract lines.")
	fmt.Println("2. Do not modify the buffer on the left.")
	fmt.Println("3. Unchanged lines are ignored.")

	cfmReader := bufio.NewReader(os.Stdin)
	_, err := cfmReader.ReadByte()
	errutil.ExitIfError(err)

	fmt.Println("Opening editor...")
}

func CreateTmpFiles(currentPath string, oldFile string, newFile string, dirMode bool, all bool) {
	entries, err := os.ReadDir(currentPath)
	errutil.ExitIfError(err)

	fOld, err := os.OpenFile(oldFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0444)
	errutil.ExitIfError(err)
	defer fOld.Close()

	fNew, err := os.Create(newFile)
	errutil.ExitIfError(err)
	defer fNew.Close()

	for _, entry := range entries {
		info, err := entry.Info()
		errutil.ExitIfError(err)

		if !all && info.Name()[0] == '.' {
			continue
		}

		if !dirMode {
			if info.IsDir() {
				continue
			}
		}

		fOld.WriteString(info.Name())
		fOld.WriteString("\n")

		fNew.WriteString(info.Name())
		fNew.WriteString("\n")
	}
	fOld.Sync()
	fNew.Sync()
}

func RemoveTmpFiles(oldFile string, newFile string) {
	os.Remove(oldFile)
	os.Remove(newFile)
}

func RunEditor(oldFile string, newFile string, editor string) {
	var args []string = []string{
		newFile,
	}

	edcmd := exec.Command(editor, args...)
	edcmd.Stdin = os.Stdin
	edcmd.Stdout = os.Stdout
	err := edcmd.Run()
	errutil.ExitIfError(err)
}

func RunEditorDiff(oldFile string, newFile string, editor string) {
	var args []string = []string{
		"-d", oldFile, newFile,
		"-c", "wincmd l",
		"-c", "foldopen",
		"-c", "autocmd BufEnter * if winnr(\"$\") == 1 | execute \"normal! :q!\\<CR>\" | endif",
	}

	edcmd := exec.Command(editor, args...)
	edcmd.Stdin = os.Stdin
	edcmd.Stdout = os.Stdout
	err := edcmd.Run()
	errutil.ExitIfError(err)
}

func GenerateRenamePair(oldFile string, newFile string) []RenamePair {
	fOld, err := os.Open(oldFile)
	errutil.ExitIfError(err)
	oldScanner := bufio.NewScanner(fOld)
	oldScanner.Split(bufio.ScanLines)

	fNew, err := os.Open(newFile)
	errutil.ExitIfError(err)
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

func StartRename(renamePairs []RenamePair, currentPath string) {
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
	errutil.ExitIfError(err)

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
		errutil.ExitIfError(err)
	}
}
