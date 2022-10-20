package errutil

import (
	"fmt"
	"os"
)

func ExitIfError(err error) {
	if err != nil {
		ExitWithError(err)
	}
}

func ExitWithError(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	os.Exit(1)
}

func PrintIfError(err error) {
	if err != nil {
		PrintWithError(err)
	}
}

func PrintWithError(err error) {
	fmt.Printf("ERROR:%s\n", err)
}
