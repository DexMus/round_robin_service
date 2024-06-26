package framework

import (
	"context"
	"fmt"
	"runtime"

	"code.byted.org/gopkg/logs/v2/log"
)

// @Function: Recover for goroutine
// @Param: context
// @Return: error
func RoutineRecoveryWithError(ctx context.Context) error {
	return printRecover(ctx, recover())
}

// @Function: Recover for goroutine
// @Param: context
// @Return:
func RoutineRecovery(ctx context.Context) {
	_ = printRecover(ctx, recover())
}

func printRecover(ctx context.Context, e interface{}) error {
	if e == nil {
		return nil
	}
	const size = 64 << 10
	buf := make([]byte, size)
	buf = buf[:runtime.Stack(buf, false)]
	err := fmt.Errorf("goroutine panic %s: %s", e, buf)
	log.V1.CtxError(ctx, "%v", err)
	return err
}
