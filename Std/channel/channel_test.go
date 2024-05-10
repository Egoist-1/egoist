package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 超时控制
func TestChannel_timeout(t *testing.T)  {
	done := do()
	select{
	case <- done:
		fmt.Println("没有超时")
	case <- time.After(10 * time.Second):
		fmt.Println("超时了")
	}
}

func do() <-chan struct{}{
	do := make(chan struct{})
	go func ()  {
		time.Sleep(5 * time.Second)
		do <- struct{}{}
	}()
	return do
}

var wg sync.WaitGroup

//显示最大并发数
func TestChannelV2(t *testing.T){
	//最大并发数2
	limits := make(chan int, 2)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func (s int)  {
			defer wg.Done()
			limits <- s	
			fmt.Println(<-limits)		
		}(i)
	}
	wg.Wait()
}


