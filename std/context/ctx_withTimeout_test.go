package contextlearning_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// Context 计时退出
func TestWithTimeout(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("子协程退出")
				return
			default:
				fmt.Printf("子协程正在执行\n")
				time.Sleep(time.Second * 1)
			}
		}
	}(ctx)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("子协程退出")
				return
			default:
				fmt.Printf("子协程正在执行\n")
				time.Sleep(time.Second * 1)
			}
		}
	}(ctx)

	time.Sleep(time.Second * 5)
	fmt.Println("调用cancel")
	cancelFunc()
	time.Sleep(time.Second * 3)
}
