package utils

import (
	"bufio"
	"context"
	"errors"
	"io"
	"os/exec"
	"runtime"
	"time"
)

const (
	osType   = runtime.GOOS
	osLinux  = "linux"
	osDarwin = "darwin"
)

var errUnsupportedOS = errors.New("unsupported OS type " + osType)

// UnixCmdSyncBytes only get stdout
func UnixCmdSyncBytes(strCmd string) (cmdRes []byte, err error) {

	if err = unixCmdOsTypeCheck(); err != nil {
		return
	}

	var cmd = exec.Command("/bin/bash", "-c", strCmd)
	cmdRes, err = cmd.Output()
	return
}

// UnixCmdSyncString only get stdout
func UnixCmdSyncString(strCmd string) (cmdRes string, err error) {

	if err = unixCmdOsTypeCheck(); err != nil {
		return
	}

	var cmd = exec.Command("/bin/bash", "-c", strCmd)
	var bs []byte
	if bs, err = cmd.Output(); err != nil {
		return
	}
	cmdRes = Bytes2Str(bs)
	return
}

// UnixCmdResLinesWithTimeout only get stdout, timeout is valid when timeout > 0
func UnixCmdResLinesWithTimeout(strCmd string, timeout time.Duration) (lines []string, err error) {

	if err = unixCmdOsTypeCheck(); err != nil {
		return
	}

	var cmd *exec.Cmd
	if timeout > 0 {
		var ctxTimeOut, ctxCancelFunc = context.
			WithTimeout(context.Background(), timeout)
		defer ctxCancelFunc()
		cmd = exec.CommandContext(ctxTimeOut, "/bin/bash", "-c", strCmd)
	} else {
		cmd = exec.Command("/bin/bash", "-c", strCmd)
	}
	var stdoutPipe io.ReadCloser
	if stdoutPipe, err = cmd.StdoutPipe(); err != nil {
		return
	}
	if err = cmd.Start(); err != nil {
		return
	}

	var errChan = make(chan error, 1)
	go func() { errChan <- cmd.Wait() }()
	var reader = bufio.NewReader(stdoutPipe)
	lines = make([]string, 0, 8)
	var line string
	var lenLine int
	for {
		select {
		case err = <-errChan:
			if err != nil {
				return
			}
		default:
			switch line, err = reader.ReadString('\n'); err {
			case nil:
				if lenLine = len(line); lenLine > 1 {
					lines = append(lines, line[:len(line)-1])
				}
			case io.EOF:
				err = nil
				if line != "" {
					lines = append(lines, line)
				}
				return
			default:
				return
			}
		}
	}
}

func unixCmdOsTypeCheck() (err error) {

	if osType != osLinux && osType != osDarwin {
		err = errUnsupportedOS
	}
	return
}
