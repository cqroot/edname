<div align="center">
  <h1>Edname</h1>
  <p><i>Use your favorite <b>editor</b> to batch <b>rename</b> files and directories.</i></p>

  <p>
    <a href="https://github.com/cqroot/edname/actions">
      <img src="https://github.com/cqroot/edname/workflows/test/badge.svg" alt="Action Status" />
    </a>
    <a href="https://codecov.io/gh/cqroot/edname">
      <img src="https://codecov.io/gh/cqroot/edname/branch/main/graph/badge.svg" alt="Codecov" />
    </a>
    <a href="https://goreportcard.com/report/github.com/cqroot/edname">
      <img src="https://goreportcard.com/badge/github.com/cqroot/edname" alt="Go Report Card" />
    </a>
    <a href="https://github.com/cqroot/edname/tags">
      <img src="https://img.shields.io/github/v/tag/cqroot/edname" alt="Git tag" />
    </a>
    <a href="https://github.com/cqroot/edname/blob/main/LICENSE">
      <img src="https://img.shields.io/github/license/cqroot/edname" />
    </a>
    <a href="https://github.com/cqroot/edname/issues">
      <img src="https://img.shields.io/github/issues/cqroot/edname" />
    </a>
  </p>
</div>

## Features

- Edit filenames in a familiar text editor environment.
- Preview changes with colored diffs before applying.
- Support for including or excluding directories and dotfiles.

## Installation

### From source

To install `edname` from source, ensure you have Go installed and run:

```bash
go install github.com/cqroot/edname@latest
```

## Usage

Execute the following command in the directory where you need to rename the files:

```bash
edname [directory]
```

This command will open an editor (`$EDITOR` or `vim`) with a buffer listing the files in the current directory.

Change the file name in the buffer, save, and exit.

Note that **do not change the number of lines and do not adjust the order**.

### Flags

- `-a, --all`: Include entries starting with '.' (dotfiles).
- `-d, --directory`: Include directories.
- `-D, --directory-only`: Rename directories only.
- `-e, --editor`: Specify the editor to use (overrides `$EDITOR`).

### Examples

1. Include directories and dotfiles:

   ```bash
   edname -a -d ~/target_directory
   ```

2. Use a specific editor:

   ```bash
   edname -e nano
   ```

3. Rename directories only:
   ```bash
   edname -D
   ```

## Development

### Running Tests

To run tests, use:

```bash
make test
```

## Contributing

Contributions are welcome! Feel free to open an issue to report a bug, suggest a new feature, or submit a pull request.

## License

This project is open source and available under the [MIT License](LICENSE).
