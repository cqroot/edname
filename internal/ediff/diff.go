package ediff

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func (e Ediff) runEditor(file string) error {
	args := append(e.editorArgs, file)

	edcmd := exec.Command(e.editor, args...)
	edcmd.Stdin = os.Stdin
	edcmd.Stdout = os.Stdout

	var errb bytes.Buffer
	edcmd.Stderr = &errb

	err := edcmd.Run()

	if e.ignoreEditorError {
		return nil
	}

	if err != nil {
		return fmt.Errorf(
			"run editor %s: %w\n    stderr = %s",
			e.editor, err, errb.String(),
		)
	}

	return nil
}

func (e Ediff) createTemp() (string, error) {
	// Create temp file
	tmp, err := os.CreateTemp("", "ediff-")
	if err != nil {
		return "", fmt.Errorf("create temp file: %w", err)
	}
	defer tmp.Close()

	// Write items to temp file
	if _, err = tmp.WriteString(strings.Join(e.items, "\n")); err != nil {
		return "", fmt.Errorf("create temp file: %w", err)
	}
	if err := tmp.Sync(); err != nil {
		return "", fmt.Errorf("create temp file: %w", err)
	}

	return tmp.Name(), nil
}

func (e Ediff) Run() ([]DiffPair, error) {
	tmpName, err := e.createTemp()
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpName)

	// Run editor
	if err := e.runEditor(tmpName); err != nil {
		return nil, err
	}

	// Read new items from temp file
	tmp, err := os.Open(tmpName)
	if err != nil {
		return nil, err
	}
	defer tmp.Close()

	scanner := bufio.NewScanner(tmp)
	scanner.Split(bufio.ScanLines)

	idx := 0
	pairs := make([]DiffPair, 0)
	for scanner.Scan() {
		newItem := scanner.Text()
		if newItem == e.items[idx] {
			idx += 1
			continue
		}

		pairs = append(pairs, DiffPair{
			Prev: e.items[idx],
			Curr: newItem,
		})
		idx += 1
	}

	if idx != len(e.items) {
		return nil, ErrDifferentItemCount
	}

	return pairs, nil
}
