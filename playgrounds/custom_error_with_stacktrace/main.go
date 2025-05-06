package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/cockroachdb/errors"

	"github.com/aazw/vscode-go/playgrounds/custom_error_with_stacktrace/cerrors"
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
		handler = cerrors.NewPrettyTextHandler(os.Stdout, &slog.HandlerOptions{
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
	err := func002()
	if err != nil {
		return cerrors.AppendCheckpoint(err)
	}
	return nil
}

func func002() error {
	//nolint:staticcheck // SA4023: デモ用にあえて残す
	err := func003()
	// nolint:staticcheck // SA4023: デモ用にあえて残す
	if err != nil {
		return cerrors.AppendMessage(err, "huge")
	}
	return nil
}

//nolint:staticcheck // SA4023: デモ用にあえて残す
func func003() error {
	return cerrors.ErrUnknown.New(
		cerrors.WithMessage("hello world"),
		cerrors.WithCause(errors.New("hoge error")),
	)
}
