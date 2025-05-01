package customerrors

import (
	"errors"
	"fmt"
)

// AppendContextualMessage は err が *CustomError の場合、AppendContextualMessage の呼び出し元をキャプチャしてコンテキストを追加する
func AppendContextualMessage(err error, msg string) error {
	var ce *CustomError
	if errors.As(err, &ce) {
		ce.addContextualMessage(msg, 2)
		return ce
	}
	return err
}

// AppendContextualMessagef は err が *CustomError の場合、AppendContextualMessagef の呼び出し元をキャプチャしてコンテキストを追加する
func AppendContextualMessagef(err error, format string, args ...any) error {
	var ce *CustomError
	if errors.As(err, &ce) {
		ce.addContextualMessage(fmt.Sprintf(format, args...), 2)
		return ce
	}
	return err
}
