package customerrors

import (
	"fmt"
	"log/slog"

	"github.com/cockroachdb/errors"
)

type CustomError struct {
	errCode string                       // エラーを一意二識別するコード
	metaMsg string                       // エラー型を説明する静的なメッセージ
	ctxMsg  string                       // 実行時のコンテキストを説明するメッセージ (contextual message)
	cause   error                        // このエラーの原因となった元のエラー
	stack   *errors.ReportableStackTrace // stacktraceを自分で保持 (errors.WithStackにすると、*withstack.withStack型になるので、slog.LogValuerのLogValue()が呼ばれない)
}

// Functional Options Pattern
type Option func(*CustomError)

func WithContextualMessage(msg string) Option {
	return func(e *CustomError) {
		e.ctxMsg = msg
	}
}

func WithContextualMessagef(format string, args ...any) Option {
	return func(e *CustomError) {
		e.ctxMsg = fmt.Sprintf(format, args...)
	}
}

func WithCause(err error) Option {
	return func(e *CustomError) {
		e.cause = err
	}
}

func newCustomError(errCode string, metaMsg string) func(...Option) error {

	return func(options ...Option) error {
		customError := &CustomError{
			errCode: errCode,
			metaMsg: metaMsg,
		}

		// Functional Options Pattern
		for _, option := range options {
			option(customError)
		}

		// １. withstack でラップして生のスタックを取る
		wrapped := errors.WithStackDepth(customError, 1)

		// ２. 取れた ReportableStackTrace を自前フィールドに格納
		customError.stack = errors.GetReportableStackTrace(wrapped)

		// ３. CustomError 本体を返す（ここだけで stack 情報持ち回り）
		return customError
	}
}

func (e *CustomError) Error() string {

	base := fmt.Sprintf("[%s] %s", e.errCode, e.metaMsg)

	if e.ctxMsg != "" {
		base += ": " + e.ctxMsg
	}

	if e.cause != nil {
		base += ": " + e.cause.Error()
	}

	return base
}

// Unwrap allows errors.Is/As to work with the underlying cause
func (e *CustomError) Unwrap() error {
	return e.cause
}

// Code returns the error code
func (e *CustomError) Code() string {
	return e.errCode
}

// MetaMsg returns the static meta message
func (e *CustomError) MetaMsg() string {
	return e.metaMsg
}

// ContextMsg returns the contextual message
func (e *CustomError) ContextMsg() string {
	return e.ctxMsg
}

// LogValue implements slog.LogValuer for structured logging
func (e *CustomError) LogValue() slog.Value {

	attrs := []slog.Attr{
		slog.String("code", e.errCode),
		slog.String("description", e.metaMsg),
	}

	if e.ctxMsg != "" {
		attrs = append(attrs, slog.String("message", e.ctxMsg))
	}

	if e.cause != nil {
		attrs = append(attrs, slog.Any("cause", e.cause))
	}

	if e.stack != nil {
		attrs = append(attrs, slog.Any("stacktrace", e.stack))
	}

	return slog.GroupValue(attrs...)
}
