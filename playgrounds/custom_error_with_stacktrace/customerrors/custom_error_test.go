package customerrors

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	stdErrors "errors"
	"log/slog"

	crErrors "github.com/cockroachdb/errors"
)

func TestCustomError_Basic(t *testing.T) {
	// cause に標準 errors を使ってみる
	cause := stdErrors.New("root cause")
	err := ErrUnknown(
		WithContextualMessage("ctx msg"),
		WithCause(cause),
	)

	// 1) Error() の文字列フォーマット
	s := err.Error()
	wantPrefix := "[UNKNOWN_ERROR] an unknown error occurred"
	if !strings.HasPrefix(s, wantPrefix) {
		t.Errorf("Error() = %q; want prefix %q", s, wantPrefix)
	}
	if !strings.Contains(s, "ctx msg") || !strings.HasSuffix(s, ": root cause") {
		t.Errorf("Error() = %q; want include context and cause", s)
	}

	// 2) Unwrap / errors.Is / errors.As の確認
	if !stdErrors.Is(err, cause) {
		t.Error("errors.Is(err, cause) = false; want true")
	}
	var ce *CustomError
	if !stdErrors.As(err, &ce) {
		t.Error("errors.As(err, *CustomError) = false; want true")
	}

	// 3) アクセサメソッド
	if got := ce.Code(); got != "UNKNOWN_ERROR" {
		t.Errorf("Code() = %q; want %q", got, "UNKNOWN_ERROR")
	}
	if got := ce.MetaMsg(); got != "an unknown error occurred" {
		t.Errorf("MetaMsg() = %q; want %q", got, "an unknown error occurred")
	}
	if got := ce.ContextMsg(); got != "ctx msg" {
		t.Errorf("ContextMsg() = %q; want %q", got, "ctx msg")
	}

	// 4) stack フィールド が埋まっていること
	if ce.stack == nil {
		t.Error("stack is nil; want non-nil ReportableStackTrace")
	}
}

func TestCustomError_JSONLogging(t *testing.T) {
	var buf bytes.Buffer

	// JSONHandler で LogValue() が呼ばれることを確認
	handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{
		AddSource: false, // テストではソース不要
	})
	logger := slog.New(handler)

	err := ErrUnknown(
		WithContextualMessage("ctx msg"),
		WithCause(crErrors.New("root cause")),
	)
	logger.Error("operation failed", "err", err)

	// 出力された JSON をパース
	var rec map[string]any
	if err2 := json.Unmarshal(buf.Bytes(), &rec); err2 != nil {
		t.Fatalf("json unmarshal error: %v\nraw: %s", err2, buf.String())
	}

	// "err" フィールドがオブジェクトになっていること
	obj, ok := rec["err"].(map[string]any)
	if !ok {
		t.Fatalf("rec[\"err\"] has wrong type: %T", rec["err"])
	}

	// 各キーの検証
	tests := []struct {
		key, want string
	}{
		{"errcode", "UNKNOWN_ERROR"},
		{"metamsg", "an unknown error occurred"},
		{"ctxmsg", "ctx msg"},
	}
	for _, tt := range tests {
		if got, _ := obj[tt.key].(string); got != tt.want {
			t.Errorf("json err.%s = %q; want %q", tt.key, got, tt.want)
		}
	}

	// stacktrace がオブジェクトになっており、
	// その中の frames が非空スライスであることをチェック
	stRaw, exists := obj["stacktrace"]
	if !exists {
		t.Error("json err.stacktrace missing")
	}
	stObj, ok := stRaw.(map[string]any)
	if !ok {
		t.Fatalf("json err.stacktrace type = %T; want object containing frames", stRaw)
	}
	framesRaw, exists := stObj["frames"]
	if !exists {
		t.Error("json err.stacktrace.frames missing")
	}
	framesArr, ok := framesRaw.([]any)
	if !ok || len(framesArr) == 0 {
		t.Errorf("json err.stacktrace.frames = %v; want non-empty array", framesRaw)
	}
}
