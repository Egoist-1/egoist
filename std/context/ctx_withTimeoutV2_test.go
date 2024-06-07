package contextlearning

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*3)
	ctx, _ = context.WithTimeout(ctx, time.Second*6)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				println("退出1")
				return
			default:
				time.Sleep(time.Second)
				println("子协程真在运行")
			}
		}
	}(ctx)

	time.Sleep(time.Second * 10)
	fmt.Println("结束")

}
