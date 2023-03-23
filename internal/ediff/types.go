package ediff

type DiffPair struct {
	Prev string
	Curr string
}

type Ediff struct {
	editor            string
	editorArgs        []string
	ignoreEditorError bool
	items             []string
}

func New(editor string) *Ediff {
	return &Ediff{
		editor:            editor,
		ignoreEditorError: false,
		items:             make([]string, 0),
	}
}

func (e *Ediff) AppendItem(item string) {
	e.items = append(e.items, item)
}

func (e *Ediff) AppendItems(items []string) {
	e.items = append(e.items, items...)
}

func (e *Ediff) SetEditorArgs(editorArgs []string) {
	e.editorArgs = editorArgs
}

func (e *Ediff) SetIgnoreEditorError(ignoreEditorError bool) {
	e.ignoreEditorError = ignoreEditorError
}
