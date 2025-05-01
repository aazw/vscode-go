package cerrors

import (
	"fmt"

	crErrors "github.com/cockroachdb/errors"
)

func init() {
	// センチネル値 (sentinel) の検証
	// ErrorKind を追加したら ErrorKindCount が一つ下に移動し、その値が必ず増える
	// len(constructors)の値と比較することで、両者の値が違えば自動的にずれ(=constructors定義追加漏れ)を検知できる
	// init() でプログラム起動時（またはテスト時）に即座にパニックを起こし検知できる
	// 単体テストでも検知可能
	if len(constructors) != int(ErrorKindCount) {
		panic(fmt.Sprintf(
			"cerrors: constructors マップの要素数(%d) が ErrorKindCount(%d) と一致しません",
			len(constructors), ErrorKindCount,
		))
	}
}

// customErrorConstructor は CustomError を生成するコンストラクタオブジェクト
// New メソッドでオプションを適用して CustomError を返す
type customErrorConstructor struct {
	errCode string
	detail  string
}

// ErrorKind は enum風に定義されたエラー種別を表す
type ErrorKind int

// New は 事前に定義されたErrorKindの情報をもとに CustomError を作成し、Functional Options Pattern でカスタマイズして返す
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
	// ErrorKindCount はセンチネル値(sentinel/配列やリストの終端を示す特別な値)であって、上記の要素数として使うためのもので、
	// 最下部にあることで iota によって、要素が追加されるたびにその値が増える.
	// この一番最下部にあるべきものである ErrorKindCount の値と len(constructors) の値をチェックすることで自動的にずれを検知できる
	// ずれ = constructorsへの定義追加漏れ
	ErrorKindCount
)

// constructors は各 ErrorKind に対するカスタムエラーコンストラクタをキー付きで保持する
var constructors = map[ErrorKind]customErrorConstructor{
	ErrUnknown:      {"UNKNOWN_ERROR", "an unknown error occurred"},
	ErrInvalidInput: {"INVALID_INPUT", "入力が無効です"},
}
