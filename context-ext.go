package utils

import (
	"context"
	"time"
)

// disableTimeoutCtx is a custom context wrapper, designed a context like timeout context, timeout with spec time, but if disableCh channel was closed or receives value before timeout, the timeout will be ignored when DeadlineExceeded happen
type disableTimeoutCtx struct {
	timeoutCtx       context.Context
	timeoutCtxCancel context.CancelFunc
	done             chan struct{}   // our own "done" channel
	cause            error           // last cause (DeadlineExceeded or explicit cancel)
	cancelCh         chan struct{}   // will be closed when cancel() invoked
	disableCh        <-chan struct{} // external signal to ignore timeout
}

var _ context.Context = (*disableTimeoutCtx)(nil)

// Deadline returns the timeout deadline (or parent's if earlier)
func (c *disableTimeoutCtx) Deadline() (deadline time.Time, ok bool) { return c.timeoutCtx.Deadline() }

// Done returns a channel closed when the context is cancelled
func (c *disableTimeoutCtx) Done() <-chan struct{} { return c.done }

// Err returns why the context was cancelled (or nil if active)
func (c *disableTimeoutCtx) Err() (err error) {

	select {
	case <-c.done:
		err = c.cause
	default:
	}
	return
}

// Value forwards to timeoutCtx
func (c *disableTimeoutCtx) Value(key any) any { return c.timeoutCtx.Value(key) }

// WithDisableTimeout creates a context with timeout that can be ignored via signal channel disableCh
//   - timeout: initial timeout duration (e.g. 5s)
//   - disableCh: when closed or receives value before timeout → the timeout will be ignored when DeadlineExceeded happen
//
// Returns:
//   - ctx: the new context
//   - cancel: call to explicitly cancel (like normal WithCancel)
//   - Note: disableCh should not be nil
func WithDisableTimeout(
	parent context.Context,
	timeout time.Duration,
	disableCh <-chan struct{},
) (context.Context, context.CancelFunc) {

	if disableCh == nil {
		panic("disableCh cannot be nil")
	}
	// make disableTimeoutCtx
	disCtx := &disableTimeoutCtx{
		done:      make(chan struct{}),
		cancelCh:  make(chan struct{}),
		disableCh: disableCh,
	}
	disCtx.timeoutCtx, disCtx.timeoutCtxCancel = context.WithTimeout(parent, timeout)
	// 1. timeout && disable, wait cancel() invoked and set disCtx.cause and close(disCtx.done)
	// 2. timeout && !disable, cancel do nothing with disCtx.done and disCtx.cause
	// 3. !timeout, cancel do nothing with disCtx.done and disCtx.cause
	go func() {
		<-disCtx.timeoutCtx.Done() // timeoutCtx was cancelled → propagate
		var timeoutCtxErr = disCtx.timeoutCtx.Err()
		if timeoutCtxErr == context.DeadlineExceeded {
			select {
			case <-disCtx.disableCh: // already disabled → ignore
				<-disCtx.cancelCh // wait cancel() invoked
				disCtx.cause = context.Canceled
				close(disCtx.done)
				return
			default:
			}
		}
		disCtx.cause = timeoutCtxErr
		close(disCtx.done)
	}()
	return disCtx, func() {
		select {
		case <-disCtx.cancelCh:
		default:
			close(disCtx.cancelCh)
		}
		disCtx.timeoutCtxCancel()
	}
}
