package utils

import (
	"errors"
	"testing"
)

var errStrForTest = errors.New("for test")
var errStrErrNotNil = errors.New("err not nil")

func TestRunWithRecover(t *testing.T) {
	type args struct {
		runFunc func() (err error)
	}
	tests := []struct {
		name               string
		args               args
		wantRunWithRecover func() (err error)
		isPanic            bool
	}{
		{
			name: "err == nil and not panic",
			args: args{
				runFunc: func() error { return nil },
			},
			// wantRunWithRecover: func() error { panic("not implemented") },
			isPanic: false,
		},
		{
			name: "err == nil but panic(errStrForTest)",
			args: args{
				runFunc: func() (err error) { err = nil; panic(errStrForTest) },
			},
			// wantRunWithRecover: func() error { panic("not implemented") },
			isPanic: true,
		},
		{
			name: "err = errStrErrNotNil and not panic",
			args: args{
				runFunc: func() (err error) { err = errStrErrNotNil; return },
			},
			// wantRunWithRecover: func() error { panic("not implemented") },
			isPanic: false,
		},
		{
			name: "err = errStrErrNotNil but panic(errStrForTest)",
			args: args{
				runFunc: func() (err error) { err = errStrErrNotNil; panic(errStrForTest) },
			},
			// wantRunWithRecover: func() error { panic("not implemented") },
			isPanic: true,
		},
	}
	var runFuncErr, recoverErr error
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if gotRunWithRecover := RunWithRecover(tt.args.runFunc); !reflect.DeepEqual(gotRunWithRecover, tt.wantRunWithRecover) {
			// 	t.Errorf("RunWithRecover() = %v, want %v", gotRunWithRecover, tt.wantRunWithRecover)
			// }
			recoverErr = RunWithRecover(tt.args.runFunc)()
			if recoverErr != nil {
				t.Logf("recoverErr: %s", recoverErr.Error())
			} else {
				t.Logf("recoverErr == nil")
			}
			if tt.isPanic {
				runFuncErr = errStrForTest
			} else {
				runFuncErr = tt.args.runFunc()
			}
			if runFuncErr != nil {
				t.Logf("runFuncErr: %s", runFuncErr.Error())
			} else {
				t.Logf("runFuncErr == nil")
			}
		})
	}
}
