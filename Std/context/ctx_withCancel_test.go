package contextlearning_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// context.withcancel 创建可取消的context
func TestWith_Cancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				println("协程结束")
				return
			default:
				i := 1
				i++
				go func(ctx context.Context, i int) {
					for {
						select {
						case <-ctx.Done():
							println("子子协程结束")
							return
						default:
							time.Sleep(time.Second * 1)
							println(i)
						}
					}
				}(ctx, i)
				time.Sleep(time.Second * 1)
				println("还没有取消")
			}

		}
	}(ctx)
	time.Sleep(time.Second * 4)
	cancel()
	time.Sleep(time.Second)
	fmt.Print("解释\n")
	time.Sleep(time.Second * 7)
	fmt.Println("完全结束")
}
