package contextlearning_test

import (
	"fmt"
	"context"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type User struct {
	Name string
}

func ForV2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("结束")
			return
		default:
			time.Sleep(time.Second)
			u := ctx.Value("key")
			user, ok := u.(*User)
			if !ok {
				panic("转换错误")
			}
			intn := rand.Intn(100)
			it := strconv.Itoa(intn)
			user.Name = "李四" + it
			fmt.Println(user.Name)
		}
	}
}

// 传递参数的context
func TestWithValue(t *testing.T) {

	u := &User{
		Name: "张三",
	}
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "key", u)
	go ForV2(ctx)

	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second)
	fmt.Println(u)
	time.Sleep(time.Second * 3)

}
