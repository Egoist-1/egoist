package contextlearning_test

import (
	"fmt"
	"testing"
	"time"
)


func Select_chan(stop chan int,i int)  {
	for {
		select{
		case <-stop:
			fmt.Printf("协程%d结束了\n",i)
			return
		default:
			fmt.Printf("%d\n",i)
			time.Sleep(time.Second  * 2)
		}
	}	
}

//子协程主动通知关闭
func TestSelect_chan(t *testing.T) {
	stop := make(chan int)	
	for i := 0; i < 10; i++ {
		go Select_chan(stop,i)
	}
	//陆续通知子协程关闭
	time.Sleep(time.Second * 3)
	for i := 0; i < 10; i++ {
		stop <- 1
	}
}