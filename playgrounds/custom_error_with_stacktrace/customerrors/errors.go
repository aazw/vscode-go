package customerrors

import (
	"fmt"

	crErrors "github.com/cockroachdb/errors"
)

func init() {
	// センチネル (sentinel) の検証
	// ErrorKind を追加したら必ず ErrorKindCount が一つ下に移動するため値が増える
	// len(constructors)の値と比較することで、これが違えば自動的にずれを検知できる
	// init() でプログラム起動時（またはテスト時）に即座にパニックを起こし検知できる
	// 単体テストでも検知可能
	if len(constructors) != int(ErrorKindCount) {
		panic(fmt.Sprintf(
			"customerrors: constructors マップの要素数(%d) が ErrorKindCount(%d) と一致しません",
			len(constructors), ErrorKindCount,
		))
	}
}

// ErrorConstructor は CustomError を生成するコンストラクタオブジェクト
// New メソッドでオプションを適用して CustomError を返す
type customErrorConstructor struct {
	errCode string
	detail  string
}

// ErrorKind は enum風に定義されたエラー種別を表す
type ErrorKind int

// New は ErrorConstructor の設定（コード／詳細）をもとに CustomError を作成し、Functional Options Pattern でカスタマイズして返す
func (ek ErrorKind) New(options ...Option) error {

	ctor, ok := constructors[ek]
	if !ok {
		return fmt.Errorf("invalid ErrorKind: %d", ek)
	}

	customError := &CustomError{
		errCode: ctor.errCode,
		detail:  ctor.detail,
	}

	// Functional Options Pattern でのオプションの処理
	for _, option := range options {
		option(customError)
	}

	// １. withstack でラップして生のスタックを取る
	wrapped := crErrors.WithStackDepth(customError, 1)

	// ２. 取れた ReportableStackTrace を自前フィールドに格納
	customError.stack = crErrors.GetReportableStackTrace(wrapped)

	// ３. CustomError 本体を返す（ここだけで stack 情報持ち回り）
	return customError
}

// ErrXxxの定義はすべてここで行う
const (
	// ErrUnknown は定義されていないエラー全般を表す
	ErrUnknown ErrorKind = iota

	// ErrInvalidInput は無効な入力を表す
	ErrInvalidInput

	// --- 新しい ErrorKind は常にこの上↑に追加する ---
	//
	// ErrorKindCount はセンチネル(配列やリストの終端を示す特別な値)であって、上記の要素数として使うためのもので、
	// 一番最下部にあるべきものである ErrorKindCountの値とlen(constructors)の値をチェックすることで自動的にずれを検知できる
	ErrorKindCount
)

// constructors は各 ErrorKind のカスタムエラーコンストラクタをキー付きで保持する
var constructors = map[ErrorKind]customErrorConstructor{
	ErrUnknown:      {"UNKNOWN_ERROR", "an unknown error occurred"},
	ErrInvalidInput: {"INVALID_INPUT", "入力が無効です"},
}
