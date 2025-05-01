package cerrors

import (
	"fmt"
	"io"
	"log/slog"
	"runtime"

	crErrors "github.com/cockroachdb/errors"
)

// contextualMsg はコンテキストメッセージと、それが追加されたファイル:行番号を保持
type contextualMsg struct {
	Message string `json:"message,omitempty"` // コンテキストを表すメッセージ本体
	File    string `json:"file,omitempty"`    // このメッセージを追加した該当するコードのある絶対パス
	Line    int    `json:"line,omitempty"`    // 上記ファイル内の該当行番号
}

// CustomError はエラーコード、詳細メッセージ、実行時コンテキスト情報、原因エラー、およびスタックトレースを保持するカスタムエラー型
type CustomError struct {
	errCode string                         // エラーを一意二識別するコード
	detail  string                         // エラー型を説明する静的な情報
	ctxMsgs []contextualMsg                // 実行時のコンテキストを説明するメッセージ等情報 (contextual message)
	cause   error                          // このエラーの原因となった元のエラー
	stack   *crErrors.ReportableStackTrace // stacktraceを自分で保持 (errors.WithStackにすると、*withstack.withStack型になるので、slog.LogValuerのLogValue()が呼ばれない)
}

// Functional Options Pattern
type Option func(*CustomError)

// WithContextualMessage はフォーマットなしでコンテキストメッセージを追加する
// 呼び出し元のファイルと行番号をキャプチャする
func WithContextualMessage(msg string) Option {
	return func(e *CustomError) {
		// 呼び出し元の情報をキャプチャ
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "?"
			line = 0
		}

		e.ctxMsgs = append(e.ctxMsgs, contextualMsg{
			Message: msg,
			File:    file, // 絶対パス
			Line:    line,
		})
	}
}

// WithContextualMessagef はフォーマット付きでコンテキストメッセージを追加する
// 呼び出し元のファイルと行番号をキャプチャする
func WithContextualMessagef(format string, args ...any) Option {
	return func(e *CustomError) {
		// 呼び出し元の情報をキャプチャ
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "?"
			line = 0
		}

		e.ctxMsgs = append(e.ctxMsgs, contextualMsg{
			Message: fmt.Sprintf(format, args...),
			File:    file, // 絶対パス
			Line:    line,
		})
	}
}

// WithCause はエラーの原因(err)をCustomErrorに設定する
func WithCause(err error) Option {
	return func(e *CustomError) {
		e.cause = err
	}
}

// addContext はskipに基づいて呼び出し元情報を取得する内部ヘルパー
func (e *CustomError) addContextualMessage(msg string, skip int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "?"
		line = 0
	}
	e.ctxMsgs = append(e.ctxMsgs, contextualMsg{Message: msg, File: file, Line: line})
}

// AddContextualMessage は作成後に新しいコンテキストメッセージを追加し、直接の呼び出し元をキャプチャする
func (e *CustomError) AddContextualMessage(msg string) {
	e.addContextualMessage(msg, 1)
}

// AddContextualMessagef は作成後に新しいコンテキストメッセージを追加し、直接の呼び出し元をキャプチャする
func (e *CustomError) AddContextualMessagef(format string, args ...any) {
	e.addContextualMessage(fmt.Sprintf(format, args...), 1)
}

// Error は CustomError のコード、詳細メッセージ、追加されたコンテキスト情報、および原因エラーを組み合わせた文字列表現を返す
// 組み込みの error インターフェースを実装
func (e *CustomError) Error() string {

	base := fmt.Sprintf("[%s] %s", e.errCode, e.detail)

	for _, cxtMsg := range e.ctxMsgs {
		base += fmt.Sprintf(": %s (%s:%d)", cxtMsg.Message, cxtMsg.File, cxtMsg.Line)
	}

	if e.cause != nil {
		base += ": " + e.cause.Error()
	}

	return base
}

// Unwrap はerrors.Is/Asが原因エラーを処理できるようにする
func (e *CustomError) Unwrap() error {
	return e.cause
}

// Code はエラーコードを返す
func (e *CustomError) Code() string {
	return e.errCode
}

// Detail は静的な詳細メッセージを返します
func (e *CustomError) Detail() string {
	return e.detail
}

// CtxMsgs はコンテキストメッセージのスライスを返す
func (e *CustomError) CtxMsgs() []contextualMsg {
	return e.ctxMsgs
}

// LogValue は構造化ロギングのためのslog.LogValuerを実装したもの
func (e *CustomError) LogValue() slog.Value {

	attrs := []slog.Attr{
		slog.String("code", e.errCode),
		slog.String("detail", e.detail),
	}

	if len(e.ctxMsgs) > 0 {
		attrs = append(attrs, slog.Any("context", e.ctxMsgs))
	}

	if e.cause != nil {
		attrs = append(attrs, slog.Any("cause", e.cause))
	}

	if e.stack != nil {
		attrs = append(attrs, slog.Any("stacktrace", e.stack))
	}

	return slog.GroupValue(attrs...)
}

// Format はfmt.Formatterを実装し、%+v指定子でスタックトレースの出力をサポートする
func (e *CustomError) Format(f fmt.State, c rune) {
	if c == 'v' && f.Flag('+') {
		io.WriteString(f, e.Error())
		if e.stack != nil {
			for _, fr := range e.stack.Frames {
				if fr.InApp {
					fmt.Fprintf(f, "\n\t%s\n\t\t%s:%d", fr.Function, fr.Filename, fr.Lineno)
				}
			}
		}
		return
	}
	io.WriteString(f, e.Error())
}
