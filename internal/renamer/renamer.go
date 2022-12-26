package renamer

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cqroot/edname/internal/file"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type RenamePair struct {
	OldName string
	NewName string
}

type Renamer struct {
	Editor     string
	Path       string
	DirOpt     bool
	DirOnlyOpt bool
	AllOpt     bool
}

func New(editor string, path string, dirOpt bool, dirOnlyOpt bool, allOpt bool) *Renamer {
	return &Renamer{
		Editor:     editor,
		Path:       path,
		DirOpt:     dirOpt,
		DirOnlyOpt: dirOnlyOpt,
		AllOpt:     allOpt,
	}
}

func (r Renamer) GenerateRenameItems() ([]string, error) {
	entries, err := os.ReadDir(r.Path)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

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

		result = append(result, info.Name())
	}

	return result, nil
}

func (r Renamer) Execute() error {
	backend := file.Backend{
		Path:       r.Path,
		DirOpt:     r.DirOpt,
		DirOnlyOpt: r.DirOnlyOpt,
		AllOpt:     r.AllOpt,
	}
	items, err := backend.Generate()
	if err != nil {
		return err
	}

	tmp, err := os.CreateTemp("", "edname-")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())

	if _, err = tmp.WriteString(strings.Join(items, "\n")); err != nil {
		return err
	}

	if err := tmp.Sync(); err != nil {
		return err
	}

	r.RunEditor(tmp.Name())

	tmp.Seek(0, 0)
	scanner := bufio.NewScanner(tmp)
	scanner.Split(bufio.ScanLines)

	idx := 0
	result := make([][]string, 0)
	for scanner.Scan() {
		newItem := scanner.Text()
		if newItem == items[idx] {
			idx += 1
			continue
		}

		pair := []string{items[idx], newItem}
		result = append(result, pair)
		idx += 1
	}

	fmt.Printf("%+v\n", result)
	if r.confirm(result) {
		backend.Rename(result, func(old string, new string) {
			fmt.Printf(
				"%s %s %s\n",
				old,
				text.FgGreen.Sprint("->"),
				new,
			)
		})
	}
	return nil
}

func (r Renamer) RunEditor(file string) error {
	var args []string = []string{
		file,
	}

	edcmd := exec.Command(r.Editor, args...)
	edcmd.Stdin = os.Stdin
	edcmd.Stdout = os.Stdout
	err := edcmd.Run()
	return err
}

func (r Renamer) confirm(renamePairs [][]string) bool {
	if len(renamePairs) == 0 {
		return false
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Old Name", "New Name"})
	for idx, pair := range renamePairs {
		t.AppendRow(table.Row{idx, pair[0], pair[1]})
	}
	t.Render()

	fmt.Print("Confirm to rename the above file [y/N] ")
	cfmReader := bufio.NewReader(os.Stdin)
	cfmText, err := cfmReader.ReadString('\n')
	if err != nil {
		return false
	}

	if cfmText != "y\n" && cfmText != "Y\n" {
		return false
	}

	return true
}
