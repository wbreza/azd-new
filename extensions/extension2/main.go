package main

import (
	"fmt"
	"os"

	"github.com/wbreza/azd-new/extensions/extension2/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}