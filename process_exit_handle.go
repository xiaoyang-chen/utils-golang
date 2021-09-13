package utils

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	exitMgrOnce      sync.Once
	exitMgrSingleton *exitMgr
)

type exitMgr struct {
	sync.Mutex
	exitFuncList []func()
	exitSign     chan os.Signal
}

func initexitMgr() {

	exitMgrSingleton = &exitMgr{
		exitFuncList: make([]func(), 0, 8),
		exitSign:     make(chan os.Signal, 1),
	}

	// syscall.SIGHUP 终端控制进程结束(终端连接断开)
	// syscall.SIGINT 用户发送INTR字符(Ctrl+C)触发
	// syscall.SIGTERM 结束程序(可以被捕获、阻塞或忽略)
	// syscall.SIGQUIT 用户发送QUIT字符(Ctrl+/)触发
	signal.Notify(
		exitMgrSingleton.exitSign,
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT,
	)
	go func() {
		<-exitMgrSingleton.exitSign
		for _, f := range exitMgrSingleton.exitFuncList {
			f()
		}
		os.Exit(0)
	}()
}

func RegisterExitFunc(exitFuncList ...func()) {

	if exitMgrSingleton == nil {
		exitMgrOnce.Do(initexitMgr)
	}

	exitMgrSingleton.Lock()
	defer exitMgrSingleton.Unlock()
	exitMgrSingleton.exitFuncList = append(exitMgrSingleton.exitFuncList, exitFuncList...)
}

func ManualExit() {

	if exitMgrSingleton == nil {
		os.Exit(0)
		return
	}
	// send many times, for blocking process
	for {
		exitMgrSingleton.exitSign <- syscall.SIGQUIT
	}
}
