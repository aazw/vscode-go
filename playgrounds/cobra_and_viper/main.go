package main

import (
	"fmt"
	"os"

	"github.com/aazw/vscode-go/playgrounds/cobra_cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		// Cobra はエラーを返すだけなので明示的に終了させる
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
