package renamer

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/cqroot/vina/internal/errutil"
)

type RenamePair struct {
	OldName string
	NewName string
}

type Renamer struct {
	OldFile string
	NewFile string

	Path string

	DirOpt     bool
	DirOnlyOpt bool
	AllOpt     bool
}

func New(path string, dirOpt bool, dirOnlyOpt bool, allOpt bool) *Renamer {
	var opsId string = fmt.Sprintf("%d", time.Now().Unix())
	var oldFile string = fmt.Sprintf("/tmp/vina-old-%s", opsId)
	var newFile string = fmt.Sprintf("/tmp/vina-new-%s", opsId)

	return &Renamer{
		NewFile:    newFile,
		OldFile:    oldFile,
		Path:       path,
		DirOpt:     dirOpt,
		DirOnlyOpt: dirOnlyOpt,
		AllOpt:     allOpt,
	}
}

func (r Renamer) GenerateRenameItems(ch chan<- string) {
	entries, err := os.ReadDir(r.Path)
	errutil.ExitIfError(err)

	for _, entry := range entries {
		info, err := entry.Info()
		errutil.ExitIfError(err)

		if !r.AllOpt && info.Name()[0] == '.' {
			continue
		}

		if r.DirOnlyOpt {
			if !info.IsDir() {
				continue
			}
		} else if !r.DirOpt {
			if info.IsDir() {
				continue
			}
		}

		ch <- info.Name()
	}

	close(ch)
}

func (r Renamer) CreateTmpFiles() {
	fOld, err := os.OpenFile(r.OldFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0444)
	errutil.ExitIfError(err)
	defer fOld.Close()

	fNew, err := os.Create(r.NewFile)
	errutil.ExitIfError(err)
	defer fNew.Close()

	chItems := make(chan string)
	go r.GenerateRenameItems(chItems)

	for file := range chItems {
		_, err = fOld.WriteString(file)
		errutil.PrintIfError(err)
		_, err = fOld.WriteString("\n")
		errutil.PrintIfError(err)

		_, err = fNew.WriteString(file)
		errutil.PrintIfError(err)
		_, err = fNew.WriteString("\n")
		errutil.PrintIfError(err)
	}

	errutil.PrintIfError(fOld.Sync())
	errutil.PrintIfError(fNew.Sync())
}

func (r Renamer) RemoveTmpFiles() {
	errutil.PrintIfError(os.Remove(r.OldFile))
	errutil.PrintIfError(os.Remove(r.NewFile))
}

func (r Renamer) RunEditor(editor string) {
	var args []string = []string{
		r.NewFile,
	}

	edcmd := exec.Command(editor, args...)
	edcmd.Stdin = os.Stdin
	edcmd.Stdout = os.Stdout
	err := edcmd.Run()
	errutil.ExitIfError(err)
}

func (r Renamer) RunEditorDiff(editor string) {
	var args []string = []string{
		"-d", r.OldFile, r.NewFile,
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

func (r Renamer) GenerateRenamePairs() []RenamePair {
	fOld, err := os.Open(r.OldFile)
	errutil.ExitIfError(err)
	oldScanner := bufio.NewScanner(fOld)
	oldScanner.Split(bufio.ScanLines)

	fNew, err := os.Open(r.NewFile)
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

func (r Renamer) StartRename(renamePairs []RenamePair) {
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
			filepath.Join(r.Path, pair.OldName),
			filepath.Join(r.Path, pair.NewName),
		)
		errutil.ExitIfError(err)
	}
}
