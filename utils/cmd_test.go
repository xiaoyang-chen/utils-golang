package utils

import (
	"reflect"
	"testing"
	"time"
)

func TestUnixCmdSyncBytes(t *testing.T) {
	type args struct {
		strCmd string
	}
	tests := []struct {
		name       string
		args       args
		wantCmdRes []byte
		wantErr    bool
	}{
		{
			name: "ls",
			args: args{
				strCmd: "ls",
			},
			wantCmdRes: []byte("LICENSE\nREADME.md\nbig.go\nbig_test.go\ncmd.go\ncmd_test.go\ngo.mod\nnumconvert.go\nnumconvert_test.go\nprocess_exit_handle.go\nprocess_exit_handle_test.go\nrun_with_recover.go\nrun_with_recover_test.go\nset.go\nsql_param.go\nsql_param_test.go\nstrbytesconv.go\nstrbytesconv_test.go\nstring_parse.go\nstring_parse_test.go\ntimeconvert.go\ntimeconvert_test.go\n"),
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCmdRes, err := UnixCmdSyncBytes(tt.args.strCmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnixCmdSyncBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCmdRes, tt.wantCmdRes) {
				t.Errorf("UnixCmdSyncBytes() = %v, %s, want %v, %s", gotCmdRes, Bytes2Str(gotCmdRes), tt.wantCmdRes, Bytes2Str(tt.wantCmdRes))
			}
		})
	}
}

func TestUnixCmdSyncString(t *testing.T) {
	type args struct {
		strCmd string
	}
	tests := []struct {
		name       string
		args       args
		wantCmdRes string
		wantErr    bool
	}{
		{
			name: "ls",
			args: args{
				strCmd: "ls",
			},
			wantCmdRes: "LICENSE\nREADME.md\nbig.go\nbig_test.go\ncmd.go\ncmd_test.go\ngo.mod\nnumconvert.go\nnumconvert_test.go\nprocess_exit_handle.go\nprocess_exit_handle_test.go\nrun_with_recover.go\nrun_with_recover_test.go\nset.go\nsql_param.go\nsql_param_test.go\nstrbytesconv.go\nstrbytesconv_test.go\nstring_parse.go\nstring_parse_test.go\ntimeconvert.go\ntimeconvert_test.go\n",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCmdRes, err := UnixCmdSyncString(tt.args.strCmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnixCmdSyncString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCmdRes != tt.wantCmdRes {
				t.Errorf("UnixCmdSyncString() = %v, want %v", gotCmdRes, tt.wantCmdRes)
			}
		})
	}
}

func TestUnixCmdResLinesWithTimeout(t *testing.T) {
	type args struct {
		strCmd  string
		timeout time.Duration
	}
	tests := []struct {
		name      string
		args      args
		wantLines []string
		wantErr   bool
	}{
		{
			name: "ls",
			args: args{
				strCmd:  "ls",
				timeout: 1 * time.Second,
			},
			wantLines: []string{
				"LICENSE",
				"README.md",
				"big.go",
				"big_test.go",
				"cmd.go",
				"cmd_test.go",
				"go.mod",
				"numconvert.go",
				"numconvert_test.go",
				"process_exit_handle.go",
				"process_exit_handle_test.go",
				"run_with_recover.go",
				"run_with_recover_test.go",
				"set.go",
				"sql_param.go",
				"sql_param_test.go",
				"strbytesconv.go",
				"strbytesconv_test.go",
				"string_parse.go",
				"string_parse_test.go",
				"timeconvert.go",
				"timeconvert_test.go",
			},
			wantErr: false,
		},
		{
			name: "echo $chenxy",
			args: args{
				strCmd:  "echo $chenxy",
				timeout: 1 * time.Second,
			},
			wantLines: []string{},
			wantErr:   false,
		},
		{
			name: "echo c",
			args: args{
				strCmd:  "echo c",
				timeout: 1 * time.Second,
			},
			wantLines: []string{"c"},
			wantErr:   false,
		},
		{
			name: "docker ps",
			args: args{
				strCmd:  "docker ps",
				timeout: 1 * time.Second,
			},
			wantLines: []string{},
			wantErr:   false,
		},
		{
			name: "chenxy",
			args: args{
				strCmd:  "chenxy",
				timeout: 1 * time.Second,
			},
			wantLines: []string{},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLines, err := UnixCmdResLinesWithTimeout(tt.args.strCmd, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnixCmdResLinesWithTimeout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLines, tt.wantLines) {
				t.Errorf("UnixCmdResLinesWithTimeout() = %v, want %v", gotLines, tt.wantLines)
			}
		})
	}
}
