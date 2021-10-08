package utils

import (
	"fmt"
	"runtime/debug"
)

const strRunWithRecoverPanicFmt = "panic err: %s\ndebug stack: %s"

// RunWithRecover 返回
func RunWithRecover(runFunc func() (err error)) (runWithRecover func() (err error)) {

	runWithRecover = func() (err error) {
		defer func() {
			if panicErr, ok := recover().(error); ok {
				err = fmt.Errorf(strRunWithRecoverPanicFmt,
					panicErr.Error(), Bytes2Str(debug.Stack()),
				)
			}
		}()
		err = runFunc()
		return
	}
	return
}
