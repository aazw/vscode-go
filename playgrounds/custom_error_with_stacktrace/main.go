package main

import (
	"log/slog"
	"os"

	"github.com/cockroachdb/errors"

	"github.com/aazw/vscode-go/playgrounds/custom_error_with_stacktrace/customerrors"
)

func main() {
	// handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	// 	AddSource: true,
	// })

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	})
	logger := slog.New(handler)

	logger.Error("error message", "err", func001())
}

func func001() error {
	return func002()
}

func func002() error {
	return func003()
}

func func003() error {
	return customerrors.ErrUnknown(
		customerrors.WithContextualMessage("hello world"),
		customerrors.WithCause(errors.New("hoge error")),
	)
}
