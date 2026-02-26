package utils

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWithDisableTimeout(t *testing.T) {

	var ctxParent = context.WithValue(context.Background(), "test1", "test1")
	ctxParent = context.WithValue(ctxParent, "test2", "test2")
	var wg sync.WaitGroup
	wg.Add(2)
	// test 1
	go func() {
		defer wg.Done()
		disable := make(chan struct{})
		ctx, cancel := WithDisableTimeout(ctxParent, 3*time.Second, disable)
		defer cancel()
		go func() { // Simulate work
			time.Sleep(2 * time.Second)
			close(disable)
			fmt.Println("test 1, Work 1st stage almost done... disabling timeout")
			time.Sleep(2 * time.Second)
			fmt.Println("test 1, Work all done")
			fmt.Println(ctx.Deadline())
			fmt.Println(ctx.Value("test1"))
		}()
		// wait
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("test 1, Finished normally (timeout was disabled)")
		case <-ctx.Done():
			fmt.Println("test 1, Context ended:", ctx.Err())
		}
	}()
	// test 2
	go func() {
		defer wg.Done()
		disable := make(chan struct{})
		ctx, cancel := WithDisableTimeout(ctxParent, 3*time.Second, disable)
		defer cancel()
		go func() { // Simulate work
			time.Sleep(4 * time.Second)
			close(disable)
			fmt.Println("test 2, Work 1st stage almost done... disabling timeout")
			time.Sleep(2 * time.Second)
			fmt.Println("test 2, Work all done")
			fmt.Println(ctx.Deadline())
			fmt.Println(ctx.Value("test2"))
		}()
		// wait
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("test 2, Finished normally (timeout was disabled)")
		case <-ctx.Done():
			fmt.Println("test 2, Context ended:", ctx.Err())
		}
	}()
	wg.Wait()

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		timeout   time.Duration
		disableCh <-chan struct{}
		want      context.Context
		want2     context.CancelFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := WithDisableTimeout(context.Background(), tt.timeout, tt.disableCh)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("WithDisableTimeout() = %v, want %v", got, tt.want)
			}
			if true {
				t.Errorf("WithDisableTimeout() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
