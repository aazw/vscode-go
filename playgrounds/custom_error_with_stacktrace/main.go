package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/cockroachdb/errors"

	"github.com/aazw/vscode-go/playgrounds/custom_error_with_stacktrace/customerrors"
)

func main() {
	var logFormat string
	flag.StringVar(&logFormat, "log-format", "text", "format: (json|pretty-text|text)")
	flag.Parse()

	var handler slog.Handler
	switch logFormat {
	case "json", "JSON", "Json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		})
	case "pretty-text", "prettyText", "ptext":
		handler = customerrors.NewPrettyTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		})
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		})
	}

	logger := slog.New(handler)

	logger.Error("error message", "err", func001())
}

func func001() error {
	return func002()
}

func func002() error {
	return customerrors.AppendContextualMessage(func003(), "huge")
}

func func003() error {
	return customerrors.ErrUnknown.New(
		customerrors.WithContextualMessage("hello world"),
		customerrors.WithCause(errors.New("hoge error")),
	)
}
