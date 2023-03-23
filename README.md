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

# Features

- Batch rename files and directories with your favorite editor.

# Installation

## From source

```bash
go install github.com/cqroot/edname@latest
```

# Usage

Execute the following command in the directory where you need to rename the files:

```bash
edname
```

This command will open an editor (`vim` by default) with a buffer listing the files in the current directory.

Change the file name in the buffer, and then exit.

Note that do not change the number of lines and do not adjust the order.
