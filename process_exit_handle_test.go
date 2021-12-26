package utils

import (
	"fmt"
	"testing"
)

func TestRegisterExitFunc(t *testing.T) {
	type args struct {
		exitFuncList []func()
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				exitFuncList: []func(){
					func() {
						fmt.Println("1.0")
					},
					func() {
						fmt.Println("1.1")
					},
				},
			},
		},
		{
			name: "2",
			args: args{
				exitFuncList: []func(){
					func() {
						fmt.Println("2.0")
					},
					func() {
						fmt.Println("2.1")
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterExitFunc(tt.args.exitFuncList...)
		})
	}
	ManualExit()
}
